#!/bin/bash +x

PROVIDER="docker.io"
USERNAME="yuril"
REPOSITORY_NAME="aerospike-ha"
VERSION_TAG="0.0.2"

image="$PROVIDER/$USERNAME/$REPOSITORY_NAME"

export DOCKER_BUILDKIT=1

make oapi-generate
docker buildx build --no-cache --push --tag $image:$VERSION_TAG --tag $image:latest --platform=linux/amd64 .
docker pull $image --platform linux/x86_64
# docker buildx build --push --tag $with_tag --tag $latest --platform=linux/arm64,linux/amd64 .
# docker run --rm -it <your-image> /bin/sh
