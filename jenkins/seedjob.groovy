folder("Deployments") {
    displayName("Kubernetes Deployments")
}

folder("Seed") {
  displayName("Seed data")
}

folder("Build") {
    displayName("Build Containers")
}

folder("Build/services") {
  displayName("Build Go Services")
}

folder("Build/docker") {
  displayName("Build Generic Docker Containers")
}

Closure scmConfiguration(String branch = "*/master", String gitUrl = 'https://github.com/mhaddon/gke-k8s') {
  return {
    gitSCM {
      branches {
        branchSpec {
          name(branch)
        }
      }
      userRemoteConfigs {
        userRemoteConfig {
          credentialsId(null)
          name('origin')
          url("${gitUrl}.git")
          refspec('+refs/heads/*:refs/remotes/origin/*')
        }
      }
      doGenerateSubmoduleConfigurations(false)
      browser {
        gitWeb {
          repoUrl(gitUrl)
        }
      }
      gitTool('')
    }
  }
}

[
  [ name: "airports", path: "services" ],
  [ name: "countries", path: "services" ],
  [ name: "frontend", path: "services" ],
  [ name: "runways", path: "services" ],
  [ name: "runways-country", path: "services" ],
  [ name: "jenkins", path: "docker" ],
  [ name: "mongo-seed", path: "docker" ]
].each { environment ->
  pipelineJob("Build/${environment.path}/${environment.name}") {
    parameters {
      gitParam('GIT_TAG_NAME') {
        description('Git tag or branch of project repo')
        type('BRANCH_TAG')
        sortMode('ASCENDING')
        defaultValue('origin/master')
      }

      stringParam("IMAGE_TAG", "latest", "Tag of docker image")
    }

    environmentVariables {
      env("SERVICE_NAME", "${environment.name}")
      env("CONTAINER_PATH", "${environment.path}")
    }

    definition {
      cpsScmFlowDefinition {
        scm(scmConfiguration('${GIT_TAG_NAME}'))
        scriptPath("./jenkins/build-docker.pipeline.groovy")
      }
    }
  }
}

pipelineJob("Build/BuildAll") {
  parameters {
    gitParam('GIT_TAG_NAME') {
      description('Git tag or branch of project repo')
      type('BRANCH_TAG')
      sortMode('ASCENDING')
      defaultValue('origin/master')
    }

    stringParam("IMAGE_TAG", "latest", "Tag of docker image")
  }

  definition {
    cpsScmFlowDefinition {
      scm(scmConfiguration('${GIT_TAG_NAME}'))
      scriptPath("./jenkins/build-all-docker.pipeline.groovy")
    }
  }
}

[
  [ name: "airports-db", path: "airports-app/airports-db.yaml" ],
  [ name: "airports-service", path: "airports-app/airports-service.yaml" ],
  [ name: "airports-seed", path: "airports-app/airports-seed.yaml" ],
  [ name: "countries-db", path: "airports-app/airports-db.yaml" ],
  [ name: "countries-service", path: "airports-app/airports-service.yaml" ],
  [ name: "countries-seed", path: "airports-app/airports-seed.yaml" ],
  [ name: "runways-db", path: "airports-app/airports-db.yaml" ],
  [ name: "runways-service", path: "airports-app/airports-service.yaml" ],
  [ name: "runways-seed", path: "airports-app/airports-seed.yaml" ],
  [ name: "runways-country-service", path: "airports-app/runways-country-service.yaml" ],
  [ name: "frontend", path: "airports-app/frontend-service.yaml" ]
].each { environment ->
  pipelineJob("Deployments/${environment.name}") {
    parameters {
      gitParam('GIT_TAG_NAME') {
        description('Git tag or branch of project repo')
        type('BRANCH_TAG')
        sortMode('ASCENDING')
        defaultValue('origin/master')
      }

      stringParam("NAMESPACE", "michael", "Namespace to deploy to")
    }

    environmentVariables {
      env("CONFIG_PATH", "${environment.path}")
    }

    definition {
      cpsScmFlowDefinition {
        scm(scmConfiguration('${GIT_TAG_NAME}'))
        scriptPath("./jenkins/apply-config.pipeline.groovy")
      }
    }
  }
}

pipelineJob("Seed/airports") {
  parameters {
    gitParam('GIT_TAG_NAME') {
      description('Git tag or branch of project repo')
      type('BRANCH_TAG')
      sortMode('ASCENDING')
      defaultValue('origin/master')
    }

    stringParam("NAMESPACE", "michael", "Namespace to deploy to")
  }

  definition {
    cpsScmFlowDefinition {
      scm(scmConfiguration('${GIT_TAG_NAME}'))
      scriptPath("./jenkins/seed-db.pipeline.groovy")
    }
  }
}