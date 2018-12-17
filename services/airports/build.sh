#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
docker build -t registry.kube-system.svc.cluster.local/airports ${DIR}
