#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

export $(./scripts/base64_encode_env_files.sh $(find . -name "*.env"))
export GCP_SA_KEY_B64="$(base64 ${DIR}/gcp_sa.key.json)"

${DIR}/scripts/create.sh
sleep 30
${DIR}/scripts/setup.sh
sleep 15
${DIR}/scripts/create_userspace.sh