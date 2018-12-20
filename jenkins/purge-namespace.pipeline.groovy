pipeline {
  agent any

  stages {
    stage('Deleting') {
      steps {
        sh "kubectl delete daemonsets,replicasets,services,deployments,pods,rc --all --namespace ${env.NAMESPACE}"
      }
    }
  }
}