#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go get \
 github.com/BurntSushi/toml \
 gopkg.in/mgo.v2 \
 github.com/gorilla/mux \
 github.com/tkanos/gonfig \
 github.com/mhaddon/gke-k8s/services/common/src/vault \
 github.com/mhaddon/gke-k8s/services/common/src/config \
 github.com/mhaddon/gke-k8s/services/common/src/helper \
 github.com/mhaddon/gke-k8s/services/common/src/models \
 github.com/mhaddon/gke-k8s/services/common/src/persistence \
 github.com/patrickmn/go-cache

mkdir -p "${DIR}/bin/"

go build -o "${DIR}/bin/runways-country" "${DIR}/src/main.go"