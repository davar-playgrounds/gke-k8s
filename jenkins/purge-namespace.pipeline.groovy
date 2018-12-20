pipeline {
  agent any

  stages {
    stage('Deleting') {
      steps {
        sh "kubectl delete daemonsets,replicasets,services,deployments,pods,rc,networkpolicies,hpa --all --namespace ${env.NAMESPACE}"
      }
    }
  }
}