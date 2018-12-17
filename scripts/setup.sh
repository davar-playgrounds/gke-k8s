#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user "$(gcloud config get-value account)"

kubectl apply -f "${DIR}/../configs/general/dashboard-view-role.yaml"

export GCLOUD_ACCESS_TOKEN="$(gcloud auth print-access-token | base64)"
export GCLOUD_USER_NAME="$(base64 <<< "oauth2accesstoken")"
export GCLOUD_HOSTNAME="$(base64 <<< "eu.gcr.io")"

while read -r NAMESPACE; do
    while read -r FILE; do
        envsubst < "${FILE}" | kubectl apply -f -
    done < <(find "${DIR}/../configs/${NAMESPACE}" -type "f" -name "*.yaml")
done < <(xargs -n 1 <<< "kube-system jenkins")