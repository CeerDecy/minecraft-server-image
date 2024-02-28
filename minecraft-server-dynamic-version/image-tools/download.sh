#!/bin/bash
version=$1

output=$(./image-tools/version-tool-linux-amd64 -version "$version")

IFS=' ' read -ra array <<< "$output"

if [[ "${array[0]}" == "ok" ]]; then
    echo "downloading ${array[1]}"
    curl "${array[1]}" -o "server.jar"
else
    echo "$output"
    exit 1
fi

echo  "Building version $version"