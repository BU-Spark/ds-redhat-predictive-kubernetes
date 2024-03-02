#!/bin/bash

DATA_FILE="./data/collected_data_unlimited.csv"
INTERVAL=${1:-1} # seconds

if [ ! -s $DATA_FILE ]; then
    echo "Timestamp,Type,Name,CPU,Memory,Count" > $DATA_FILE
fi

while true; do
    NOW=$(date '+%Y-%m-%d %H:%M:%S')

    COUNT=$(kubectl top nodes --no-headers | wc -l)
    kubectl top nodes --no-headers | while read line; do
        NAME=$(echo $line | awk '{print $1}')
        CPU=$(echo $line | awk '{print $2}')
        MEMORY=$(echo $line | awk '{print $4}')
        echo "$NOW,Node,$NAME,$CPU,$MEMORY,$COUNT" >> $DATA_FILE
    done

    COUNT=$(kubectl top pods --namespace=default --no-headers | wc -l)
    kubectl top pods --namespace=default --no-headers | while read line; do
        NAME=$(echo $line | awk '{print $1}')
        CPU=$(echo $line | awk '{print $2}')
        MEMORY=$(echo $line | awk '{print $3}')
        echo "$NOW,Pod,$NAME,$CPU,$MEMORY,$COUNT" >> $DATA_FILE
    done

    sleep $INTERVAL
done

