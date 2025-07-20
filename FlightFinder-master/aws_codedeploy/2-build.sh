#!/usr/bin/env bash

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
set -x
cd /FlightFinder/
ls -la
go build -o flight-finder cmd/finder_web/main.go
