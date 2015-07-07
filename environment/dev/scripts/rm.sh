#!/bin/bash

CURRENT_PATH=$(dirname "${BASH_SOURCE[0]}")
source $CURRENT_PATH/config.sh

docker rm -fv $CONTAINER_NAME