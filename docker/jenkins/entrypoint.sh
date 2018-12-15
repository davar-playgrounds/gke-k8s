#!/bin/bash -e

if [[ ! -d /var/jenkins_home/plugins ]]; then
  echo "Installing Jenkins plugins..."
  install-plugins.sh \
      ansicolor \
      job-dsl \
      kubernetes-cd \
      credentials-binding \
      envinject \
      timestamper \
      pipeline-model-definition \
      build-pipeline-plugin \
      cloudbees-folder \
      filesystem_scm \
      workflow-aggregator \
      jdk-tool \
      git \
      git-parameter
fi

echo "Starting Jenkins"
/configs/create_init_job.sh & /sbin/tini -- /usr/local/bin/jenkins.sh $@
