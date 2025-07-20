#!/usr/bin/env bash 

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
set -x

# give the app some warm-up time - checking connections to CloudWatch and Redis
for i in {1..5}; do
    echo "Validating service liveness: attempt #$i..."
    curl --fail localhost:80 > index.html && echo "Service is live" && exit 0
    sleep 5
done;

echo "Service not live"
exit 1