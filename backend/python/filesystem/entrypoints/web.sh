#!/bin/bash

if [ ${DEV:-false} = 'true' ]; then
  poetry run flask run --host 0.0.0.0 --port 8080
else
  poetry run uwsgi -s /tmp/uwsgi.sock \
                   --manage-script-name \
                   --mount /=src/entry.py \
                   --callable flask_app \
                   --http 0.0.0.0:8080
fi
