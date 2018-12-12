#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

${DIR}/scripts/create.sh
${DIR}/scripts/setup.sh
${DIR}/scripts/create_userspace.sh