pipeline {
  agent any

  stages {
    stage('Compile') {
      steps {
        dir("./services/${env.SERVICE_NAME}") {
          sh "./build.sh"
        }
      }
    }
  }
}