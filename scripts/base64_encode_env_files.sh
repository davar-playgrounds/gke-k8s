#!/usr/bin/env bash

xargs -I'{}' -n 1 sh -c "awk '{ split(\$0,a,\"=\"); gsub(/\"/, \"\", a[2]); \"base64 <<< \\\"\"a[2]\"\\\"\" | getline x; print a[1]\"_B64=\"x\"\n\"a[1]\"=\\\"\"a[2]\"\\\"\" }' <<< '{}'" < <(cat "${@}")
