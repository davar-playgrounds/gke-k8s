#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

(cd ${DIR}/secrets/${1} && ./deploy.sh ${2})