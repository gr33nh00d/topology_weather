#!/bin/bash

echo "hello"
cityID=0
for i in {1..2}
do
  # go run callOpenmap.go $i > output.txt
  cityID= sed -n "$ip" cityIDs.txt
  echo $cityID
  # go run /home/topology_weather/callOpenmap.go $cityID > output.txt
  # go run parse.go >> data.txt
done
