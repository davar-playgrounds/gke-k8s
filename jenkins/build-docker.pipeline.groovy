pipeline {
  agent any

  stages {
    stage('Login To Docker') {
      steps {
        sh "set +x; docker login -u '${env.GCLOUD_USER_NAME.trim()}' -p '${env.GCLOUD_ACCESS_TOKEN.trim()}' https://${env.GCLOUD_HOSTNAME.trim()}"
      }
    }

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
          sh "docker push ${env.GCLOUD_HOSTNAME.trim()}/${env.GCLOUD_PROJECTNAME.trim()}/${env.SERVICE_NAME}:${env.IMAGE_TAG.trim()}"
        }
      }
    }
  }
}