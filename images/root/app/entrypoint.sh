#!/bin/bash

WEBAPP_BIND="0.0.0.0:8080"
WEBAPP_APP="app.webapp:GAPPLICATION"
WEBAPP_LEVEL="debug"

gunicorn3 --bind $WEBAPP_BIND $WEBAPP_APP --log-level $WEBAPP_LEVEL --access-logfile - --error-logfile -

exit $?
