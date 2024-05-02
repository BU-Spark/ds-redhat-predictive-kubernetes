#!/bin/bash

INTERVAL=${1:-1}

while true; do
  # Get current hour
  HOUR=$(date +%H)
  HOUR=$((10#$HOUR))

  MAX_PROB=100

  # Probabilities
  if ((HOUR < 4)); then
    PROBABILITY=5
  elif ((HOUR < 5)); then
    PROBABILITY=10
  elif ((HOUR < 6)); then
    PROBABILITY=20
  elif ((HOUR < 7)); then
    PROBABILITY=40
  elif ((HOUR < 8)); then
    PROBABILITY=80
  elif ((HOUR < 9)); then
    PROBABILITY=95
  elif ((HOUR < 10)); then
    PROBABILITY=70
  elif ((HOUR < 12)); then
    PROBABILITY=40
  elif ((HOUR < 13)); then
    PROBABILITY=50
  elif ((HOUR < 14)); then
    PROBABILITY=30
  elif ((HOUR < 16)); then
    PROBABILITY=20
  elif ((HOUR < 17)); then
    PROBABILITY=40
  elif ((HOUR < 19)); then
    PROBABILITY=50
  elif ((HOUR < 20)); then
    PROBABILITY=70
  elif ((HOUR < 22)); then
    PROBABILITY=40
  else
    PROBABILITY=20
  fi

  RANDOM_NUM=$((RANDOM % MAX_PROB))

  if [ $RANDOM_NUM -lt $PROBABILITY ]; then
    # echo "TEST"
    curl http://localhost:9999/traffic_test >> logs.txt 2>/dev/null &
  fi

  sleep $INTERVAL
done

