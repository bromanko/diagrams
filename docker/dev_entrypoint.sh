#!/usr/bin/env bash

set -e
cd /usr/src/app


function _yarn_install {
  if [ package.json -nt node_modules/.up-to-date ]; then
    echo "Installing node modules..."
    yarn
  fi
  touch node_modules/.up-to-date
}

_yarn_install

echo $@
eval $@
