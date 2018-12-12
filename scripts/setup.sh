#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user "$(gcloud config get-value account)"

kubectl apply -f "${DIR}/../configs/general/dashboard-view-role.yaml"

while read -r NAMESPACE; do
    while read -r FILE; do
        kubectl apply -f "${FILE}"
    done < <(find "${DIR}/../configs/${NAMESPACE}" -type "f" -name "*.yaml")
done < <(xargs -n 1 <<< "kube-system jenkins")