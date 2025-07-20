#!/usr/bin/env bash

# This script is intended to be executed from within "openapitools/openapi-generator-cli" Docker container
set -e


SPEC_PATH=api/openapi3.yaml
APISERVER_PATH=cmd/finder_web/app/apiserver
APISERVER_PORT=8080

# Generate golang server 
docker-entrypoint.sh \
    generate \
    -i ${SPEC_PATH} \
    -g go-gin-server \
    --additional-properties=enumClassPrefix=true,serverPort=${APISERVER_PORT} \
    --package-name=apiserver \
    -o ${APISERVER_PATH}

# Remove redundant files
rm "${APISERVER_PATH}/main.go" 
rm "${APISERVER_PATH}/go/api_default.go" 
rm "${APISERVER_PATH}/Dockerfile" 
rm "${APISERVER_PATH}/.openapi-generator-ignore" 
rm -rf "${APISERVER_PATH}/api" 
rm -rf "${APISERVER_PATH}/.openapi-generator" 

# Remove redundant code
sed "/Index is the index handler/,+3  d" -i "${APISERVER_PATH}/go/routers.go" 