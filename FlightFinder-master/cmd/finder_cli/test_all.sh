#!/usr/bin/env bash

go build . # -gcflags=-B .
./finder_cli -flights_data=../../assets < test_cases
rm finder_cli