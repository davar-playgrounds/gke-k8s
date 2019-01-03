#!/usr/bin/env bash

env $(sed -e 's/#.*$//' -e '/^$/d' <(cat "${@:2}")) bash -c "./scripts/envsubst_ex.sh '$(echo $1 | tr ';' '\n')'"