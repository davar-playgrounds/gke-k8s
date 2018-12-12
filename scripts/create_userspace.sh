#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

TYPE="${1:-apply}"

while read -r ENTRY; do

  USERNAME="$(jq -r ".name" <<< "${ENTRY}")"
  ADMIN="$(jq -r ".admin" <<< "${ENTRY}")"
  ENDPOINT="$(kubectl config view -o jsonpath="{.clusters[?(@.name == \"$(kubectl config current-context)\")].cluster.server}" | xargs)"
  SECRETS_DIR="${DIR}/../secrets/${USERNAME}"

  while read -r FILE; do
    NAMESPACE="${USERNAME}" envsubst < "${FILE}" | kubectl ${TYPE} -f -
  done < <(find "${DIR}/../configs/userspace" -type "f" -name "*.yaml")

  mkdir -p "${SECRETS_DIR}"
  secret="$(kubectl get sa ${USERNAME} -o json --namespace ${USERNAME} | jq -r .secrets[].name)"
  kubectl get secret ${secret} -o json --namespace ${USERNAME} | jq -r '.data["ca.crt"]' | base64 -D > "${SECRETS_DIR}/ca.crt"
  kubectl get secret ${secret} -o json --namespace ${USERNAME} | jq -r '.data["token"]' | base64 -D > "${SECRETS_DIR}/user_token"

  CERTIFICATE="$(base64 < "${SECRETS_DIR}/ca.crt")" ENDPOINT="${ENDPOINT}" USERNAME="${USERNAME}" TOKEN="$(<"${SECRETS_DIR}/user_token")" envsubst < "${DIR}/../configs/kube.config.tmpl" > "${SECRETS_DIR}/kube.config"

  ENDPOINT="${ENDPOINT}" USERNAME="${USERNAME}" TOKEN="$(<"${SECRETS_DIR}/user_token")" envsubst < "${DIR}/install_script.sh.tmpl" > "${SECRETS_DIR}/install_script.sh"
  chmod +x "${SECRETS_DIR}/install_script.sh"

  cp "${DIR}/deploy.sh" "${SECRETS_DIR}/deploy.sh"

  (cd "${SECRETS_DIR}/../" && tar czf "${USERNAME}.tar.gz" "${USERNAME}")

  if [[ "${ADMIN}" == "true" ]]; then
    USERNAME="${USERNAME}" NAMESPACE="${USERNAME}" envsubst < "${DIR}/../configs/admin/cluster-admin.yaml" | kubectl apply -f -
  fi

done < <(jq -c '.[]' < "${DIR}/../configs/users.json")
