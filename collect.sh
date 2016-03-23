#!/bin/bash

declare -a cities=('Baker' 'Billings' 'Bozeman' 'Butte' 'Cut_Bank' 'Glasgow' 'Great_Falls' 'Havre' 'Helena' 'Kalispell' 'Lewistown' 'Miles_City' 'Missoula' 'Pinehurst_Ranch' 'Sidney' 'West_Glacier' 'West_Yellowstone' 'Wolf_Point')

while true; do
  echo "hello"
  for i in "${cities[@]}"
  do
    curl http://api.wunderground.com/api/2d3a55e50b9c3448/conditions/q/MT/Wolf_Point.json > temp.txt
    python gather.py
  done
  sleep 17280
done
