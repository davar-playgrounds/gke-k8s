#!/usr/bin/env bash

wget https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-228.0.0-linux-x86_64.tar.gz -O "${HOME}/google-cloud-sdk.tar.gz"
tar xzf ${HOME}/google-cloud-sdk.tar.gz -C ${HOME}
rm ${HOME}/google-cloud-sdk.tar.gz

yes | ${HOME}/google-cloud-sdk/install.sh

echo "export PATH=\${PATH}:${HOME}/google-cloud-sdk/bin" >> "${HOME}/.profile"
source "${HOME}/.profile"

gcloud auth activate-service-account --key-file "${HOME}/.creds/gcp_sa.key.json"

yes | gcloud auth configure-docker

ln -s /root/google-cloud-sdk/bin/gcloud /bin/gcloud