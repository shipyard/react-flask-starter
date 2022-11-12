#!/bin/bash

CELERY_PID=$(cat /var/run/celery/worker.pid)
kill -HUP $CELERY_PID
