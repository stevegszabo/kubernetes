#!/bin/bash

set -o errexit

DOCKER_VERSION=${1-1.0.0}
DOCKER_FILE=Dockerfile.webapp
DOCKER_IMAGE=notebook.local:5000/webapp:$DOCKER_VERSION

sed -i "s/WEBAPP_VERSION=.*/WEBAPP_VERSION=$DOCKER_VERSION/g" $DOCKER_FILE

docker build -t $DOCKER_IMAGE -f $DOCKER_FILE .
docker push $DOCKER_IMAGE

exit 0
