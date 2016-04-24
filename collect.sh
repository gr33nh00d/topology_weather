#!/bin/bash

cityID=0
F="$1"
P="$2"
echo "$F"
echo "$P"
  while [[ true ]]; do
    rm output.txt
    touch output.txt
    for i in {1..144}
    do
      cityID="$(sed -n $i'p' cityIDs.txt)"
      go run $F $cityID >> output.txt
    done
    go run $P
    sleep 21600
  done
