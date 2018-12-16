pipeline {
  agent any

  stages {
    stage('Build & Compile') {
      steps {
        dir("./services/${env.SERVICE_NAME}") {
          sh "./build.sh ${env.IMAGE_TAG}"
        }
      }
    }
  }
}