# Notes

## Running

If you want to run your own version of 'zanerock/react-flask-starter':
1. Get a Giphy API token.
2. Get your Shipayrd API token.

When running locally:
- Create a `.env.development.local` with `GIPHY_API_KEY=< your Giphy API key>`. This way, you can call `/api/v1/gif-search` without needing to set the `giphy_api_key` URL parameter.
- You can call all the API and S3 bucket endpoints from `http://localhost:8080` thanks to the reverse proxy routing.

To support the GitHub Action, add the following secrets to your GitHub repo:
- `GIPHY_API_KEY`
- `SHIPYARD_API_TOKEN`

## Cool stuff

- Added [Go based server to handle new API endpoints](https://github.com/zanerock/react-flask-starter/tree/main/backend-go) with [new service in the `docker-compose.yml` file](https://github.com/zanerock/react-flask-starter/blob/main/docker-compose.yml#L32).
- Added nginx reverse proxy to the [`docker-compose.yml`](https://github.com/zanerock/react-flask-starter/blob/main/docker-compose.yml#L85) with a [simple configuration](https://github.com/zanerock/react-flask-starter/blob/main/proxy/reverse_proxy.conf).
- Made it so you can set the Giphy API key in the go server's host environment to simplify calls to the `/api/v1/gif-search` endpoint.
- The `/api/v1/gif-search` endpoint downloads the GIF images and [uploads them to the S3 bucket in parallel using Goroutines](https://github.com/zanerock/react-flask-starter/blob/main/backend-go/route-gif-search-create.go#L162).
- [The call to Giphy is automatically retried with exponential backoff.](https://github.com/zanerock/react-flask-starter/blob/main/backend-go/route-gif-search-create.go#L103)
- There's pretty decent error reporting and correct status stetting (400 vs 500) throughout.

## Issue and solutions

### Go based API related issues

I wanted to do the new API in golang since that's the new language of choice. Overall, this was pretty smooth, but there were a few issues directly related to go.

- Tranlating JSON to and from go structs is notoriously tedious. Luckily, knowing this I searched for and found [this handy JSON to go struct tool](https://mholt.github.io/json-to-go/) which I used to generate the [necessary code](https://github.com/zanerock/react-flask-starter/blob/main/backend-go/types.go) (which I then broke up into parts). TODO: link to code.
- I ran into an issue when calling `/api/v1/files/upload` from the go server. I would get "HTTP/1.x transport connection broken" errors and it wasn't clear why.
  - Eventually, I looked at the HTML spec for ['Form content types'](https://www.w3.org/TR/html401/interact/forms.html#h-17.13.4) and compared the example for 'multipart/form-data' request bodies with embedded file data to what the go server was sending to the python endpoint and I eventually noticed that the closing boundary string was missing from the request the go server was sending.
  - It turns out that the position of the `multipartWriter.Close()` is critical and must come before the upload request object is created. While this does make sense, it's a bit idiomatic as typically closing a resource is done with a `defer` statement and has no effect on the actual contents. To be fair, this [is noted in the go mime/multipart docs](https://pkg.go.dev/mime/multipart#Writer.Close), but not the other parts of the `Part` and `multipart.Writer` docs that I did read. The tutorials/examples I'd referenced ([here](https://hayageek.com/golang-send-file-upload-with-post-request/) and [here](https://stackoverflow.com/questions/51234464/upload-a-file-with-post-request-golang)) did show the `Close()` in the proper place an I'm not sure how it got messed up in the first place (since I did copy the example code as a starting point). I think I had some other error, and had maybe moved things around not realizing the significance of the timing of the `Close()`. Mostly just bad luck on this one.
- On a related note, the standard [`CreateFormFile()`](https://pkg.go.dev/mime/multipart#Writer.CreateFormFile) function, which creates the entry in the multipart/form-data `Part` for the file, does not have any way to express the actual type or encoding of the file.
  - While investigating the above problem, I though this might be the issue. Though it turned out not to be the problem[^1], I would have ended up wanting to address it anyway since it would be nice to have the content type saved in the S3 object entry so it would show up correctly later when retrieved.[^2]
  - I ended up researching this issue and applying [suggested workarounds](https://github.com/golang/go/issues/46771) to create my own data `Part` encodings that would reflect the correct type and encoding of the GIF image.

[^1]: The lack of proper content type and encoding headers in the file part might actually have caused other issues, but since I fixed it first, I'm not sure and didn't want to undo it just to find out.
[^2]: There might actually be some automatic content type info appended by the S3 bucket based on the file extension or something. I think this is true in some cases with the real S3 service, but whether that's actually true, and whether it's implemented in localstack, I'm not sure and, again, not worth undoing useful fixes just to find out.

### Python based API issues

- It's a bit odd that the `/api/v1/files/` endpoint needs to be called to create the S3 bucket. It's an odd side effect that I luckily sussed out pretty quick, but it would be good to note this or, better yet, create the bucket as part of the Python server initialization.
- The existing `/api/v1/files/uploads` does not use the specified filename of the uploaded file and instead creates a generic time based name.
  - This seems to contradict the directions in the provided spec. While you could take the direction in the spec to refer to only the artifacts saved with the GitHub action, it's not clear.
  - It also means that multiple files uploaded within the same second will overwrite each other.
  - See [suggested fix](#files-upload-fix).

### Gliphy issues

- Getting an account and setting up the API token and all that were fairly simple. The API/dev docs weren't anywhere on the main page that I saw, but googling fixed that easy enough.
- There was some confusion about the results from the Gliphy search API.
  - Looking through the results, I figured the `images.original.url` would be the link to the actual GIF.
  - However, when I tested this URL in the browser, I got a web page that showed the GIF, but embedded in a page (and not just the GIF itself).
  - Since I needed the GIF itself, I spent some time trying to figure out how to get the URL without scraping it from the web page.
  - Luckily, I trued `curl`-ing the URL and found that in that case, I actually did get the GIF itself. I gather that the Giphy server is idiomatically is looking at the `User-Agent` header or something to decide what to serve.

### Postgres/data storage issues

- I wanted to have an SQL file to create the `gif_search` table so I could store the results for the `/gif-search` endpoint (and retrieve them for `/gif-search/{id}`), but I ran into issues.
  - The Postgres image page talked about putting scripts in `/docker-entrypoint-init.d`, so I tried mapping a local directory to that directory on the host, but it wasn't creating the table and the logs were silent. I'm pretty sure my script was good, but I couldn't test it directly (see below).
  - Looking at the [`Dockerfile`](https://github.com/docker-library/postgres/blob/ce54cce510ed5da4ed9e1e66ddeb6e3300786813/13/bookworm/Dockerfile), I see where it [creates the directory in the Docker virtual host](https://github.com/docker-library/postgres/blob/ce54cce510ed5da4ed9e1e66ddeb6e3300786813/13/bookworm/Dockerfile#L75), but there didn't see to a `COPY` or anything.
  - I even tried copying the `Dockerfile` and init scripts to a `database/` dir, adding the `COPY` commands to the new `database/Dockerfile` and building from that setup, but it still didn't seem to work. I read through the [`docker-entrypoint.sh`](https://github.com/docker-library/postgres/blob/ce54cce510ed5da4ed9e1e66ddeb6e3300786813/13/bookworm/docker-entrypoint.sh) setup script, and I could see where it was supposed to [process the SQL files](https://github.com/docker-library/postgres/blob/ce54cce510ed5da4ed9e1e66ddeb6e3300786813/13/bookworm/docker-entrypoint.sh#L195), but it didn't seem to be.
  - The black box nature of everything was further deepened by the fact that no matter what I did, I was unable to connect to the Postgres DB directly. `psql -h localhost` failed with: "connection to server at "localhost" (127.0.0.1), port 5432 failed: Connection refused" (repeated err messages for IPV6 localhost).
  - I got the same result when attempting to connect from my local machine, and when logged into the Postgres server.
  - Even more mysterious, when I would run `ps aux` as root, it would only show the shell and `ps` process itself (with the original postgres 9.6 image).
  - I even tried looking directly in `/proc`, but it only showed the shell and `ls` processes.
  - Rather than try to debug the setup scripts with all this weirdness, I decided to look around in the existing source to see how it sets up the tables used in the ping API and found what I needed in the `migrations` folder.
  - I dug up the sqlalchemy 1.3 docs and figured out pretty quickly how to setup the new table I needed.
- There were also issues connecting Go to the DB.
  - My go server was initially unable to connect to the Postgres database, returning a similar error to the one I got when attempting to manually connect. (In retrospect, this may have been solvable with additional configuration, but seeing as things were acting really weird, I wanted to try and upgrade Postgres anyway.)
  - I tried upgrading to the latest Postgres, but that broke the Python server (possibly incompatible driver; I didn't investigate much).
  - I tried version 13, and that worked for everyone.

### Unifying the API

- Wanted to be able to hit the Python and Go server through the same host+port.
- Looked for a Go package to (reverse) proxy the calls to Python.
- Looked for a Python package to (reverse) proxy the calls to Go.
- Decided cleanest/simplest would be leverage Docker Compose and add an nginx reverse proxy (configuration ended up being 3 entries).
- I realized later that it might be possible to do something like this using Shipyard configurations, but I still think it's nice to have something that works the same for local and shipyard builds.

### Adding Python packages

- I ended up not needing to change the packages, but once I ran into the issue of flask breaking, I felt compelled to figure out a solution.
  - Had to mess around with local install of python in order to be able to install poetry so I can run `poetry lock`.
  - Regenerating the lock file caused breakage because apparently all Python projects are shit at specifying requirements; would find the breakage was due to incompatible changes included in a dependency whose version was non specified.
  - At first tried upgrading to supported packages, but that just caused a cascade of failures, again, due to multiple Python packages failing to actually specify dependency versions.
  - Decided to use the original, working `poetry.lock` to extract the specific installed working versions of libs; iteratively explicitly specify lib in `pyproject.toml`, run `poetry lock`, and attempt restart, hit another dependency failure, repeat.
  - Adding the following seemed to work:
  ```
  jinja2 = "2.11.3"
  markupsafe = "1.1.1"
  itsdangerous = "1.1.0"
  werkzeug = "1.0.1"
  sqlalchemy = "1.3.23"
  ```

## Recommendations

- Update the `pyproject.toml` requirements to pin the indirect library requirements since flask doesn't bother to specify it's actual requirements and will break when the requirements are updated. (Already done in [`shipyard/flask-react-starter` PR#19](https://github.com/shipyard/react-flask-starter/pull/19) submitted by myself).
- Provide instruction on where to find the GIF URL in the Giphy results and note that programmatic retrieval will retrieve the actual GIF even though viewing the URL through a browser does not.
- <span id="files-upload-fix"></span>Update the `/api/vi/files/upload/` endpoint to:
  1. [Fix return status when no 'file' provided to `/api/vi/files/upload/`.](https://github.com/zanerock/react-flask-starter/blob/main/backend/src/routes/localstack.py#L32)
  2. [Use `file.filename` for the file (if provided) and generate a filename (== object Key) only if necessary and fix the generated filename logic which results in the same filename being created for multiple uploads within the same second.](https://github.com/zanerock/react-flask-starter/blob/main/backend/src/routes/localstack.py#L35)
  3. [Take cognizance of the file 'Content-Type'.](https://github.com/zanerock/react-flask-starter/blob/main/backend/src/routes/localstack.py#L45)
  You can see my changes [here].

## TODOs

If this were a real project I would:
- Include proper method documentation.
- Do unit testing.
- Provide more extensive integration testing in the GitHub workflows.
