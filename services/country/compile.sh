#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

go get github.com/BurntSushi/toml gopkg.in/mgo.v2 github.com/gorilla/mux github.com/tkanos/gonfig

mkdir -p "${DIR}/bin/"

go build -o "${DIR}/bin/countries" "${DIR}/src/main.go"