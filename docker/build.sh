#!/bin/bash

set -o errexit

DOCKER_VERSION=${1-1.0.0}
DOCKER_FILE=Dockerfile

DOCKER_REGISTRY=docker.io
DOCKER_USERNAME=steveszabo
DOCKER_IMAGE=$DOCKER_REGISTRY/$DOCKER_USERNAME/webapp:$DOCKER_VERSION

sed -i "s/WEBAPP_VERSION=.*/WEBAPP_VERSION=$DOCKER_VERSION/g" $DOCKER_FILE

##docker login -u $DOCKER_USERNAME $DOCKER_REGISTRY
printf "DOCKER_REGISTRY: [%s]\n" $DOCKER_REGISTRY
printf "DOCKER_USERNAME: [%s]\n" $DOCKER_USERNAME
printf "DOCKER_IMAGE: [%s]\n" $DOCKER_IMAGE
##docker build -t $DOCKER_IMAGE -f $DOCKER_FILE .
##docker push $DOCKER_IMAGE

exit 0
