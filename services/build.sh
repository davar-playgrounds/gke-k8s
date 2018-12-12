#!/usr/bin/env bash

set -x

./airports/build.sh
./countries/build.sh
./runways/build.sh
./frontend/build.sh