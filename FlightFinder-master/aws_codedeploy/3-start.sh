#!/usr/bin/env bash 

# This file is intended for AWS CodeDeploy service
# Referenced by "appspec.yml"
set -x
cd /FlightFinder/
setsid ./flight-finder \
    -flights_data=./assets \
    -web_data=./web \
    -port=80 \
    -aws_xray=true \
    -aws_region=us-east-1 \
    -redis_addr=flightfinderrediscache.mufbi3.ng.0001.use1.cache.amazonaws.com:6379 \
    -redis_pass="" \
    >/FlightFinder/flight-finder.logs 2>&1 < /FlightFinder/flight-finder.logs & # run as daemon
