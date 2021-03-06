#!/bin/bash

CURRENT_PATH=$(pwd)/$(dirname "${BASH_SOURCE[0]}")
source $CURRENT_PATH/config.sh


if [[ -n $(docker ps | grep $CONTAINER_NAME | awk '{ print $1 }') ]]; then
    docker exec -it $CONTAINER_NAME $@
else
    docker run \
           --rm \
           -it \
           --name $CONTAINER_NAME \
           -v $CURRENT_PATH/../../../:/go/src/github.com/slok/daton \
           -p 3000:3000 \
           -p 6060:6060 \
           $IMAGE_NAME $@
fi