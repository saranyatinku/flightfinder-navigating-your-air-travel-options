#!/usr/bin/env bash

IMAGE_NAME="myredis"
REDIS_PORT=6379

function stage() {
    BLUE="\e[36m"
    RESET="\e[0m"
    msg="$1"
    
    echo
    echo -e "$BLUE$msg$RESET"
}

function checkPrerequsites() {
    stage "Checking prerequisites"

    command docker version > /dev/null 2>&1
    [[ $? != 0 ]] && echo "You need to install docker to run Redis" && exit 1
    
    echo "Done"
}

function runRedis() {
    stage "Running dockerized Redis server"

    SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
    docker run \
        --rm \
        --name $IMAGE_NAME \
        -p $REDIS_PORT:$REDIS_PORT \
        -v $SCRIPT_DIR/conf:/conf \
        redis:latest redis-server /conf/redis.conf # run redis-server with config file (to set a password)

    echo "Done"
}


checkPrerequsites
runRedis