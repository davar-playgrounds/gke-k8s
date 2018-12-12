#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

function openWebPage {
    local DOMAIN="http://localhost:${1:-8081}/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/#!/overview?namespace=$(basename "${DIR}")"

    until $(curl --output /dev/null --silent --head --fail "${DOMAIN}"); do
      sleep 2
    done

    echo "${DOMAIN}"
    open "${DOMAIN}"
}

openWebPage "${1:-8081}" &

docker run -it -v "$(pwd):/secrets" --entrypoint "/secrets/install_script.sh" --workdir "/secrets" -p "${1:-8081}:8081" lachlanevenson/k8s-kubectl