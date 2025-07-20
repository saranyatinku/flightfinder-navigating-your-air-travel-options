#!/usr/bin/env bash

#
# CodeDeploy helper script
#

set -e # fail fast


# upload application to this S3 location
S3_BUCKET="codedeploy-input-artifacts" # bucket must already exist
S3_KEY="FlightFinder.zip" # existing object will be overridden

# then deploy the application here
APP_NAME="FlightFinder" # application in CodeDeploy must exist
DEPLOYMENT_GROUP_NAME="FlightFinderDeployGroup" # deployment group in CodeDeploy must exist

# configure aws cli
export AWS_PROFILE=rozneg
export AWS_DEFAULT_REGION=us-east-1

# do the job!
echo "Deploying FlightFinder to AWS CodeDeploy"
echo

echo "go vet ."
go vet . && echo "OK"
echo 

echo "Package & Push to bucket s3://${S3_BUCKET}/${S3_KEY}..."
aws deploy push --application-name ${APP_NAME} --s3-location s3://${S3_BUCKET}/${S3_KEY} --ignore-hidden-files && echo "OK"
echo

echo "Deploy to ${APP_NAME}/${DEPLOYMENT_GROUP_NAME}..."
aws deploy create-deployment --application-name ${APP_NAME} --s3-location bucket=${S3_BUCKET},key=${S3_KEY},bundleType=zip --deployment-group-name ${DEPLOYMENT_GROUP_NAME} && echo "OK"

echo
echo "Done."