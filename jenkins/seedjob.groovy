import groovy.json.JsonSlurper

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

final List userspaces = readFileFromWorkspace('configs/users').split(',')
final String deployments = readFileFromWorkspace('configs/deployments.json')

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
    displayName("${environment.name}")

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

(new JsonSlurper().parse(deployments.getBytes())).each { environment ->
  pipelineJob("Deployments/${environment.name}") {
    displayName("${environment.name}")

    parameters {
      gitParam('GIT_TAG_NAME') {
        description('Git tag or branch of project repo')
        type('BRANCH_TAG')
        sortMode('ASCENDING')
        defaultValue('origin/master')
      }

      choiceParam('NAMESPACE', userspaces)
    }

    environmentVariables {
      env("CONFIG_PATH", "${environment.path}")
      env("NAME", "${environment.name}")
      env("TYPE", "${environment.type}")
      env("MATCH", "${environment.match ?: "Running"}")
      env("ENV", "${(environment.env ?: []).join(";")}")
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
    displayName("${environment}")

    parameters {
      gitParam('GIT_TAG_NAME') {
        description('Git tag or branch of project repo')
        type('BRANCH_TAG')
        sortMode('ASCENDING')
        defaultValue('origin/master')
      }

      choiceParam('NAMESPACE', userspaces)
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

pipelineJob("BuildAll") {
  displayName("Build All")

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

pipelineJob("DeployAll") {
  displayName("Deploy All")

  parameters {
    gitParam('GIT_TAG_NAME') {
      description('Git tag or branch of project repo')
      type('BRANCH_TAG')
      sortMode('ASCENDING')
      defaultValue('origin/master')
    }

    choiceParam('NAMESPACE', userspaces)
  }

  definition {
    cpsScmFlowDefinition {
      scm(scmConfiguration('${GIT_TAG_NAME}'))
      scriptPath("./jenkins/deploy-all-services.pipeline.groovy")
    }
  }
}

pipelineJob("DeleteAll") {
  displayName("Delete All")

  parameters {
    gitParam('GIT_TAG_NAME') {
      description('Git tag or branch of project repo')
      type('BRANCH_TAG')
      sortMode('ASCENDING')
      defaultValue('origin/master')
    }

    choiceParam('NAMESPACE', userspaces)
  }

  definition {
    cpsScmFlowDefinition {
      scm(scmConfiguration('${GIT_TAG_NAME}'))
      scriptPath("./jenkins/purge-namespace.pipeline.groovy")
    }
  }
}