#!/bin/bash

cityID=0
F="$1"
P="$2"
echo "$F"
echo "$P"
for i in {1..144}
do
  cityID="$(sed -n $i'p' cityIDs.txt)"
  go run $F $cityID > output.txt
  go run $P >> data.txt
  echo $i
done
