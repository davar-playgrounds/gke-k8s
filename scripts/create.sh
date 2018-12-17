#!/usr/bin/env bash

gcloud container clusters create mhaddon-k8-test \
 --cluster-version="1.11.3-gke.18" \
 --num-nodes="3" \
 --addons="KubernetesDashboard" \
 --enable-network-policy \
 --no-enable-legacy-authorization