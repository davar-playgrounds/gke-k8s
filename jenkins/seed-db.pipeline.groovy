pipeline {
  agent any

  stages {
    stage('Deploy seed') {
      steps {
        build(
            job: "/Deployments/airports-seed",
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
        sh "csvjson ./data/csv/airports.csv > ./data/json/airports.json"
        stash(name: 'data', includes: "data/json/*")
      }
    }

    stage('Deploying') {
      steps {
        sh "while [[ ! \$(kubectl get pods --namespace ${NAMESPACE} | grep -w 'airports-seed' | awk '{ print \$3 }') = 'Running' ]]; do echo 'Waiting for airports-seed'; sleep 2; done"
      }
    }

    stage('Upload data') {
      steps {
        unstash 'data'

        sh "kubectl cp ./data/json/airports.json ${NAMESPACE}/airports-seed:/app/data.json"
      }
    }
  }
}