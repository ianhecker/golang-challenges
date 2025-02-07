#!/bin/bash

start_time=$(date +%s)

trap ctrl_c INT

ctrl_c() {
    end_time=$(date +%s)
    elapsed_time=$((end_time - start_time))

    hours=$((elapsed_time / 3600))
    minutes=$(((elapsed_time % 3600) / 60))
    seconds=$((elapsed_time % 60))

    printf "\nTimer stopped.\nElapsed time: %02d:%02d:%02d (hh:mm:ss)\n" "$hours" "$minutes" "$seconds"
    exit 0
}

echo "Timer started. Press Ctrl+C to stop."
while true; do
    sleep 1
done
