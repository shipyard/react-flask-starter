# Notes

- Need to add `requests` Python module to backend server.
  - Had to mess around with local install of python in order to be able to install poetry so I can run `poetry lock`.
  - Regenerating the lock file caused breakage because apparently all Python projects are shit at specifying requirements; would find the breakage was due to incompatible changes included in a dependency whose version was non specified.
  - At first tried upgradin to supported packages, but that just caused a cascade of failures, again, due to multiple Python packages failing to actually specify dependency versions.
  - Ended up forcing install of Python 3.8 (rather than default 3.12) to eliminate that as an issue.
  - Eventually used the original, working `poetry.lock` to extract the specific installed working versions of lib; iteratively explicitly specify lib in `pyproject.toml`, run `poetry lock`, and attempt restart, hit another dependency failure, repeat.
  - The following seemed to work:
  ```
  jinja2 = "2.11.3"
  markupsafe = "1.1.1"
  itsdangerous = "1.1.0"
  werkzeug = "1.0.1"
  sqlalchemy = "1.3.23"
  ```
