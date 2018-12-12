#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

mkdir -p "${DIR}/json"

while read -r FILE_PATH; do
    FILE_NAME="$(basename "${FILE_PATH}" ".csv")"

    echo "converting ${FILE_NAME}..."

    csvjson "${FILE_PATH}" > "${DIR}/json/${FILE_NAME}.json"
done < <(find "${DIR}/csv" -type "f" -name "*.csv")