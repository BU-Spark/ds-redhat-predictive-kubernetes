#!/bin/bash

INTERVAL=${1:-1}

while true; do
  # Get current minute
  MINUTE=$(date +%M)
  MINUTE=$((10#$MINUTE))

  # Probability: 0~10: 1/6, 10~20: 1/5, 20~30: 1/4, 30~40: 1/3, 40~50: 1/2, 50~60: 1/1
  if ((MINUTE < 10)); then
    PROBABILITY=6
  elif ((MINUTE < 20)); then
    PROBABILITY=5
  elif ((MINUTE < 30)); then
    PROBABILITY=4
  elif ((MINUTE < 40)); then
    PROBABILITY=3
  elif ((MINUTE < 50)); then
    PROBABILITY=2
  else
    PROBABILITY=1
  fi

  RANDOM_NUM=$((RANDOM % PROBABILITY))

  if [ $RANDOM_NUM -eq 0 ]; then
    # echo "TEST"
    curl http://localhost:9999/traffic_test
  fi

  sleep $INTERVAL
done

