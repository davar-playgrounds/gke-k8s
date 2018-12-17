#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export $(./scripts/base64_encode_env_files.sh $(find . -name "*.env"))

${DIR}/scripts/create.sh
sleep 30
${DIR}/scripts/setup.sh
sleep 15
${DIR}/scripts/create_userspace.sh