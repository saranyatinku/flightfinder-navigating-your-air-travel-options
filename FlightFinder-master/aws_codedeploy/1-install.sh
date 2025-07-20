#!/usr/bin/env bash

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
set -x
yum update -y
yum install -y golang
mkdir -p /FlightFinder