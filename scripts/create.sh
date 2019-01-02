#!/usr/bin/env bash

gcloud container clusters create "${CLUSTER_NAME}" \
 --cluster-version="${CLUSTER_VERSION}" \
 --num-nodes="${CLUSTER_NODES_PER_ZONE}" \
 --machine-type="${CLUSTER_MACHINE_TYPE}" \
 --addons="KubernetesDashboard" \
 --region="${CLUSTER_REGION}" \
 --node-locations="${CLUSTER_NODE_ZONES}" \
 --enable-network-policy \
 --no-enable-legacy-authorization \
 --scopes storage-rw