#!/bin/bash

CURRENT_PATH=$(dirname "${BASH_SOURCE[0]}")
source $CURRENT_PATH/config.sh

docker build -t $IMAGE_NAME -f environment/dev/Dockerfile $CURRENT_PATH/../../../