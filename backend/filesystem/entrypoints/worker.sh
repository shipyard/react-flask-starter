#!/bin/bash

if [ ${DEV:-False} = 'true' ]; then
  # Run celery in the background
  # NOTE: -B is for the heartbeat task scheduler
  poetry run celery -B \
    -A src.entry.celery worker \
    --loglevel=info \
    --pidfile=/var/run/celery/worker.pid
else
  # Run celery in the foreground
  poetry run celery -B \
    -A src.entry.celery worker \
    --loglevel=info
fi
