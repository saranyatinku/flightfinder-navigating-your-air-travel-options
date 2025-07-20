#!/usr/bin/env bash

command /sbin/time -V > /dev/null
[[ $? -ne 0 ]] && echo "You need to install 'time' eg. sudo pacman -S time. Exiting now" && exit 1


go build . 
echo -n "[`date`] " >> bench_history # -n = no newline
/sbin/time -f '%E' 2>&1 ./finder_cli -flights_data=../../assets  < test_cases > /dev/null | tee --append bench_history # %E = wall time
rm finder_cli