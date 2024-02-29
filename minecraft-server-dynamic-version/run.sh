#!/bin/bash

# 定义文件路径
file_path="server.jar"
jre=""

# 判断文件是否存在
if [ -f "$file_path" ]; then
    echo "server file exists, prepare running"
else
    echo "server file does not exist, prepare downloading"
    ./image-tools/download.sh "$SERVER_VERSION"
    jre="$SERVER_JRE"
    if [ -z "$jre" ]; then
      echo "jre is empty, parsing"
      jre=$(./image-tools/parse-jre.sh "$SERVER_VERSION")
    fi
    sed -i s/deb.debian.org/mirrors.aliyun.com/g /etc/apt/sources.list  \
        && apt-get update -y  \
        && apt-get install -y "$jre" \
        && rm -rf /var/lib/apt/lists/*
    chmod +x server.jar
    echo "init jre"
fi

echo "Starting server..."
java -Xmx"$SERVER_XMX" -Xms"$SERVER_XMS" -jar "$SERVER_JAR" nogui