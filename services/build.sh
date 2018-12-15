#!/usr/bin/env bash

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

${DIR}/airports/build.sh
${DIR}/countries/build.sh
${DIR}/runways/build.sh
${DIR}/frontend/build.sh
${DIR}/runways-country/build.sh