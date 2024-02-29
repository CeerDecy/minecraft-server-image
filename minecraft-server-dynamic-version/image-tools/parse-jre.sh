#!/bin/bash
tag="$1"
version="${tag#*.}"
version="${version%.*}"
# 根据不同的条件输出不同的值
if (( version >= 1 && version <= 6 )); then
  echo "openjdk-8-jre-headless"
elif (( version >= 7 && version <= 12 )); then
  echo "openjdk-8-jre-headless"
elif (( version >= 13 && version <= 16 )); then
  echo "openjdk-11-jre-headless"
elif (( version >=17 && version <= 18 )); then
  echo "openjdk-17-jre"
else
  echo "openjdk-17-jre"
fi