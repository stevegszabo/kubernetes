# docker images

```
DOCKER_REGISTRY=notebook.local:5000
DOCKER_IMAGE=webapp
DOCKER_TAG=latest
DOCKER_FILE=Dockerfile.webapp
DOCKER_URL=http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG
DOCKER_LAYER=$(curl -s $DOCKER_URL | jq -r .fsLayers[0].blobSum)

# build and push
docker build -t $DOCKER_REGISTRY/$DOCKER_IMAGE:$DOCKER_TAG -f $DOCKER_FILE .
docker push $DOCKER_REGISTRY/$DOCKER_IMAGE:$DOCKER_TAG

# registry catalog
curl -s http://$DOCKER_REGISTRY/v2/_catalog | jq -r .
curl -s http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/tags/list | jq -r .
curl -s http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/tags/list | jq -r .tags[] | sort

# image manifest
curl -s http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | jq -r .
curl -s http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | jq -r .fsLayers[].blobSum
curl -s http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | jq -r .signatures[].signature

# image header
curl -s -I http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG
curl -s -I http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | grep Content-Type
curl -s -I http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | grep Content-Length
curl -s -I http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG | grep Docker-Content-Digest

# download image layer
curl -O http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/blobs/$DOCKER_LAYER

# pull image by digest
DOCKER_PULL_IMAGE=$DOCKER_REGISTRY/$DOCKER_IMAGE:$DOCKER_TAG
DOCKER_DIGEST=$(docker image inspect $DOCKER_PULL_IMAGE | jq -r .[0].RepoDigests[0] | awk -F@ '{print $2}')

docker image inspect $DOCKER_PULL_IMAGE | jq -r .

docker pull $DOCKER_REGISTRY/$DOCKER_IMAGE@$DOCKER_DIGEST
```
