pipeline {
  agent any

  stages {
    stage('Apply Config') {
      steps {
        sh "bash -c 'export \$(./scripts/base64_encode_env_files.sh \$(find . -name \"*.env\")); ./scripts/envsubst_ex.sh \"./configs/${env.CONFIG_PATH}\" | kubectl apply -f -'"
      }
    }

    stage('Deploying...') {
      steps {
        script {
          if ( env.TYPE == "deployment" ) {
            sh "kubectl rollout status 'deployment/${env.NAME}' --namespace ${env.NAMESPACE} -w &"
          } else if (env.TYPE == "pod") {
            sh "while [[ ! \$(kubectl get pods --namespace ${env.NAMESPACE} | grep -w '${env.NAME}' | awk '{ print \$3 }') = 'Running' ]]; do echo 'Waiting for ${env.NAME}'; sleep 2; done"
          }
        }
      }
    }
  }
}