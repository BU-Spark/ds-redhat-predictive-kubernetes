#!/bin/bash

INTERVAL=${1:-1}

while true; do
  curl http://localhost:9999/traffic_test
  sleep $INTERVAL
done

