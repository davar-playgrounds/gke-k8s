#!/usr/bin/env bash

envsubst "$(env | awk '{ split($0,a,"="); print "$"a[1] }' | xargs | sed "s/ /:/g")" <<< "${1}"