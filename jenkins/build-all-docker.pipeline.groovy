String[] BuildJobs=[]

node {
  stage('Get Jobs') {
    String AllJobs = sh(returnStdout: true, script: "awk '/string/' \${JENKINS_HOME}/javaposse.jobdsl.plugin.ExecuteDslScripts.xml | sed 's/<[^>]*>//g' | xargs").trim()
    BuildJobs = sh(returnStdout: true, script: "echo '${AllJobs}' | xargs -n 1 | awk '/services\\/|docker\\//'").split('\n')
  }

  parallel BuildJobs.collect { String job ->
    def tasks = [:]
    tasks["${job}"] = {
      stage("${job}") {
        build(
          job: "${job}",
          parameters: [
            [ $class: 'StringParameterValue', name: 'GIT_TAG_NAME', value: "${env.GIT_TAG_NAME}" ],
            [ $class: 'StringParameterValue', name: 'IMAGE_TAG', value: "${env.IMAGE_TAG}" ]
          ],
          propagate: false,
          wait: true
        )
      }
    }
    return tasks
  }.inject { result, i ->
    return result + i
  }
}