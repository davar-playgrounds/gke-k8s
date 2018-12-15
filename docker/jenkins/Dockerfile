FROM jenkins/jenkins:lts-alpine

COPY entrypoint.sh /usr/local/bin/entrypoint.sh
COPY configs/* /configs/

ENV JAVA_OPTS="-Djenkins.install.runSetupWizard=false -Dorg.jenkinsci.plugins.durabletask.BourneShellScript.HEARTBEAT_CHECK_INTERVAL=3600"
ENV JENKINS_URL="http://jenkins.jenkins.svc.cluster.local"
ENV KUBECONFIG="/configs/kube.config"

USER root
ADD https://storage.googleapis.com/kubernetes-release/release/v1.13.1/bin/linux/amd64/kubectl /usr/local/bin/kubectl
RUN chown jenkins:jenkins /configs && \
    chmod +x /usr/local/bin/kubectl && \
    apk update -q && \
    apk add jq wget openrc docker gettext libintl curl bash --no-cache && \
    rc-update add docker boot && \
    git config --global user.email "jenkins@haddon.me" && \
    git config --global user.name "Jenkins"
# USER jenkins

ENTRYPOINT /usr/local/bin/entrypoint.sh
