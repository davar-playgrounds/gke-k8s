#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
docker build -t eu.gcr.io/michaelhaddon-223413/jenkins:${1:-latest} ${DIR}
