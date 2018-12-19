pipeline {
  agent any

  stages {
    stage('Apply Config') {
      steps {
        sh "bash -c 'export \$(./scripts/base64_encode_env_files.sh \$(find . -name \"*.env\")); ./scripts/envsubst_ex.sh \"./configs/${env.CONFIG_PATH}\" | kubectl apply -f -'"
      }
    }
  }
}