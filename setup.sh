#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export $(./scripts/base64_encode_env_files.sh $(find . -name "*.env"))

${DIR}/scripts/create.sh
${DIR}/scripts/setup.sh
${DIR}/scripts/create_userspace.sh