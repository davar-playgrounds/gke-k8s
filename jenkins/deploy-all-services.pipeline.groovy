node {
  [
    [
      name: "configs",
      group: [ "Deployments/seed-config" ]
    ], [
      name: "databases",
      group: [ "Deployments/airports-db", "Deployments/countries-db", "Deployments/runways-db" ]
    ], [
      name: "seeds",
      group: [ "Seed/airports", "Seed/countries", "Seed/runways" ]
    ], [
      name: "services",
      group: [ "Deployments/airports", "Deployments/countries", "Deployments/runways", "Deployments/runways-country", "Deployments/frontend" ]
    ]
  ].each { deploymentGroup ->
    stage( "${deploymentGroup.name}" ) {
      parallel deploymentGroup.group.collect { String job ->
        def tasks = [:]
        tasks["${job}"] = {
          stage("${job}") {
            build(
              job: "${job}",
              parameters: [
                [ $class: 'StringParameterValue', name: 'GIT_TAG_NAME', value: "${env.GIT_TAG_NAME}" ],
                [ $class: 'StringParameterValue', name: 'NAMESPACE', value: "${env.NAMESPACE}" ]
              ],
              propagate: true,
              wait: true
            )
          }
        }
        return tasks
      }.inject { result, i ->
        return result + i
      }
    }
  }
}