#!/usr/bin/env bash

set -eu

jenkins-cli() {
    java -jar "/var/jenkins_home/war/WEB-INF/jenkins-cli.jar" -s "http://localhost:8080" "$@"
}

jenkins-cli-auth() {
    jenkins-cli -auth "admin:test" "$@"
}

bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:8080)" != "200" ]]; do sleep 2; done'

jenkins-cli create-job seedjob < /configs/init-job-template.xml
jenkins-cli build seedjob

sleep 20

/configs/move_files.sh

jenkins-cli reload-configuration 2>/dev/null || true
jenkins-cli safe-restart 2>/dev/null || true

bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' -u admin:test localhost:8080)" != "200" ]]; do sleep 2; done'

touch /var/jenkins_started