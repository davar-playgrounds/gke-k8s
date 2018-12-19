pipeline {
  agent any

  stages {
    stage('Deploy seed') {
      steps {
        build(
          job: "/Deployments/${env.SERVICE_NAME}-seed",
          parameters: [
            [ $class: 'StringParameterValue', name: 'GIT_TAG_NAME', value: "${env.GIT_TAG_NAME}" ],
            [ $class: 'StringParameterValue', name: 'NAMESPACE', value: "${env.NAMESPACE}" ]
          ],
          propagate: true,
          wait: true
        )
      }
    }

    stage('Convert data') {
      agent {
        docker {
          image 'jehrhardt/csvkit:latest'
        }
      }

      steps {
        sh "mkdir -p ./data/json"
        sh "csvjson ./data/csv/${env.SERVICE_NAME}.csv > ./data/json/${env.SERVICE_NAME}.json"
        stash(name: 'data', includes: "data/json/*")
      }
    }

    stage('Upload data') {
      steps {
        unstash 'data'

        sh "kubectl cp ./data/json/${env.SERVICE_NAME}.json ${env.NAMESPACE}/${env.SERVICE_NAME}-seed:/app/data.json"
      }
    }
  }
}