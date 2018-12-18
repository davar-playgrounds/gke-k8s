pipeline {
  agent any

  stages {
    stage('Build & Compile') {
      steps {
        dir("./${env.CONTAINER_PATH}/${env.SERVICE_NAME}") {
          sh "docker build -t ${env.GCLOUD_HOSTNAME.trim()}/${env.GCLOUD_PROJECTNAME.trim()}/${env.SERVICE_NAME}:${env.IMAGE_TAG.trim()} ."
        }
      }
    }

    stage('Push') {
      steps {
        dir("./${env.CONTAINER_PATH}/${env.SERVICE_NAME}") {
          sh "gcloud docker -- push ${env.GCLOUD_HOSTNAME.trim()}/${env.GCLOUD_PROJECTNAME.trim()}/${env.SERVICE_NAME}:${env.IMAGE_TAG.trim()}"
        }
      }
    }
  }
}