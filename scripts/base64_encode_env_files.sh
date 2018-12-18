#!/usr/bin/env bash

xargs -I'{}' -n 1 bash -c "awk '{ split(\$0,a,\"=\"); gsub(/\"/, \"\", a[2]); gsub(\" \", \"\\\ \", \$0); \"echo -n \\\"\"a[2]\"\\\" | base64\" | getline x; print a[1]\"_B64=\"x\"\n\"\$0 }' <<< '{}'" < <(cat "${@}")
