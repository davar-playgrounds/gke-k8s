#!/usr/bin/env bash

gcloud container clusters create "${CLUSTER_NAME}" \
 --cluster-version="${CLUSTER_VERSION}" \
 --num-nodes="${CLUSTER_NODES}" \
 --machine-type="${CLUSTER_MACHINE_TYPE}" \
 --addons="KubernetesDashboard" \
 --enable-network-policy \
 --no-enable-legacy-authorization \
 --scopes storage-rw