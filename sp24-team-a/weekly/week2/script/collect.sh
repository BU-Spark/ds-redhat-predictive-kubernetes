#!/bin/bash

DATA_FILE="./collected_data.csv"
INTERVAL=1 # seconds

if [ ! -s $DATA_FILE ]; then
    echo "Timestamp,Type,Name,CPU,Memory" > $DATA_FILE
fi

while true; do
    NOW=$(date '+%Y-%m-%d %H:%M:%S')

    kubectl top nodes --no-headers | while read line; do
        NAME=$(echo $line | awk '{print $1}')
        CPU=$(echo $line | awk '{print $2}')
        MEMORY=$(echo $line | awk '{print $4}')
        echo "$NOW,Node,$NAME,$CPU,$MEMORY" >> $DATA_FILE
    done

    kubectl top pods --all-namespaces --no-headers | while read line; do
        NAME=$(echo $line | awk '{print $2}')
        CPU=$(echo $line | awk '{print $3}')
        MEMORY=$(echo $line | awk '{print $4}')
        echo "$NOW,Pod,$NAME,$CPU,$MEMORY" >> $DATA_FILE
    done

    sleep $INTERVAL
done

