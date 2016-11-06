#!/usr/bin/env bash

set -e

NAME=`cat package.json | python -c "import json,sys;obj=json.load(sys.stdin);print obj['name']"`

DC_RUN_PARAMS="--service-ports"
function dc_add_env {
  DC_RUN_PARAMS="$DC_RUN_PARAMS -e $1=$2"
}

# Executables
DC=docker-compose
D=docker

## Docker compose names
DEV=dev

## Exported docker name
D_NS=bromanko
D_IMAGE_NAME=${NAME}
D_SERVICE_IMAGE=${D_NS}/${D_IMAGE_NAME}

BUILD_TAG=${BUILD_TAG:-localdev}
PACKAGE_VERSION=`cat package.json | python -c "import json,sys;obj=json.load(sys.stdin);print obj['version']"`
VERSION="${PACKAGE_VERSION}.${BUILD_TAG}"

dc_add_env VERSION ${VERSION}

function helptext {
    echo "Usage: ./go <command>"
    echo ""
    echo "Available commands are:"
    echo "    start                   Starts the development server"
}

function start {
  ${DC} run --rm ${DC_RUN_PARAMS} ${DEV} yarn run start
}

function test {
  info "Running ${1-all} tests"

  case "${1-all}" in
    with-coverage) CMD="test-cov"
    ;;
    watch) CMD="watch"
    ;;
    all) CMD="test"
    ;;
    *) echo -e $"Unknown command - please see usage below\n"
      helptext
      exit 1
  esac

  ${DC} run --rm ${DC_RUN_PARAMS} --entrypoint npm ${APEX_IMAGE} run ${CMD}
}

echo "Service wrapper for ${NAME}. current version: ${VERSION}"

case "$1" in
    help) helptext
    ;;
    start) shift; start
    ;;
    *)
      helptext
      exit 1
esac
