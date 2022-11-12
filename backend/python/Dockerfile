FROM python:3.8.8-alpine3.13

WORKDIR /srv

# Install system dependencies for: uWSGI, poetry, watchman
RUN apk add --update --no-cache \
      gcc \
      libc-dev \
      libffi-dev \
      openssl-dev \
      bash \
      git \
      libtool \
      m4 \
      g++ \
      autoconf \
      automake \
      build-base \
      postgresql-dev

# cryptography module incompatibility with PEP517
# https://github.com/pyca/cryptography/issues/5771
ENV CRYPTOGRAPHY_DONT_BUILD_RUST=1

# Install poetry
RUN pip install --upgrade pip
RUN pip install poetry

# Create an app user, prepare permissions, and run as the user
RUN adduser -S app
RUN mkdir -p /var/run/celery && \
    chown -R app /var/run/celery /srv
USER app

# Install Python dependencies
ADD pyproject.toml poetry.lock ./
RUN poetry install

# Add the project
# NOTE Run the install again to install the project
ADD src ./src
ADD migrations ./migrations
RUN poetry install

# Add system files
ADD filesystem /

# Set the default command
ENV FLASK_APP src/entry:flask_app
CMD /entrypoints/web.sh
