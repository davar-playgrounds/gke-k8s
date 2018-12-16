folder("Deployments") {
    displayName("Kubernetes Deployments")
}

folder("Build") {
    displayName("Build Containers")
}

Closure scmConfiguration(String gitUrl = 'https://github.com/mhaddon/gke-k8s', String branch = "*/master") {
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

["airports", "countries", "frontend", "runways", "runways-country"].each { goService ->
  pipelineJob("Build/${goService}") {
    parameters {
      stringParam("IMAGE_TAG", "latest", "Tag of docker image")
    }

    environmentVariables {
      env("SERVICE_NAME", "${goService}")
    }

    definition {
      cpsScmFlowDefinition {
        scm(scmConfiguration())
        scriptPath("./jenkins/build-go-service.pipeline.groovy")
      }
    }
  }
}