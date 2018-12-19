#!/usr/bin/env bash

gcloud container clusters create "${CLUSTER_NAME}" \
 --cluster-version="1.11.3-gke.18" \
 --num-nodes="2" \
 --machine-type="n1-standard-4" \
 --addons="KubernetesDashboard" \
 --enable-network-policy \
 --no-enable-legacy-authorization \
 --scopes storage-rw