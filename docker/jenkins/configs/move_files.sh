#!/bin/bash -e

mkdir -p /var/jenkins_home
cp /configs/config.xml /var/jenkins_home/config.xml
cp /configs/users.xml /var/jenkins_home/users/config.xml

mkdir -p /var/jenkins_home/users/admin
cp /configs/admin_user.xml /var/jenkins_home/users/admin/config.xml