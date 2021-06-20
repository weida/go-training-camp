#!/bin/bash

datasize=(  10 20 50 100 200 1000  5000)
n=1000000

for size in ${datasize[@]}:
do
    echo "datasize = $size not using Pipeline"
    ./redis-benchmark -q -r 10000 -c 1 -n 100000 -d $size  -t get,set
    echo "datasize = $size using Pipeline -P 16 "
    ./redis-benchmark -q -r 10000 -c 1 -n 100000 -d $size -P 16 -t get,set

done

