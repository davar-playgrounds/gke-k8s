#!/usr/bin/env bash

bash -c 'while [[ "$(curl -s -o /dev/null -w ''%{http_code}'' localhost:8080)" != "200" ]]; do sleep 2; done'

java -jar /var/jenkins_home/war/WEB-INF/jenkins-cli.jar -s http://localhost:8080 create-job seedjob < /configs/init-job-template.xml
java -jar /var/jenkins_home/war/WEB-INF/jenkins-cli.jar -s http://localhost:8080 build seedjob
