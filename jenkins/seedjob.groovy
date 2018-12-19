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

[
  [ name: "airports-db", path: "airports-app/airports-db.yaml", type: "deployment" ],
  [ name: "airports-service", path: "airports-app/airports-service.yaml", type: "deployment" ],
  [ name: "airports-seed", path: "airports-app/airports-seed.yaml", type: "pod" ],
  [ name: "countries-db", path: "airports-app/airports-db.yaml", type: "deployment" ],
  [ name: "countries-service", path: "airports-app/airports-service.yaml", type: "deployment" ],
  [ name: "countries-seed", path: "airports-app/airports-seed.yaml", type: "pod" ],
  [ name: "runways-db", path: "airports-app/airports-db.yaml", type: "deployment" ],
  [ name: "runways-service", path: "airports-app/airports-service.yaml", type: "deployment" ],
  [ name: "runways-seed", path: "airports-app/airports-seed.yaml", type: "pod" ],
  [ name: "runways-country-service", path: "airports-app/runways-country-service.yaml", type: "deployment" ],
  [ name: "frontend", path: "airports-app/frontend-service.yaml", type: "deployment" ]
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
      env("NAME", "${environment.name}")
      env("TYPE", "${environment.type}")
    }

    definition {
      cpsScmFlowDefinition {
        scm(scmConfiguration('${GIT_TAG_NAME}'))
        scriptPath("./jenkins/apply-config.pipeline.groovy")
      }
    }
  }
}

[
  "airports", "countries", "runways"
].each { environment ->
  pipelineJob("Seed/${environment}") {
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
      env("SERVICE_NAME", "${environment}")
    }

    definition {
      cpsScmFlowDefinition {
        scm(scmConfiguration('${GIT_TAG_NAME}'))
        scriptPath("./jenkins/seed-db.pipeline.groovy")
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

pipelineJob("Deployments/DeployAll") {
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
      scriptPath("./jenkins/deploy-all-services.pipeline.groovy")
    }
  }
}