#!/bin/bash

# 定义文件路径
file_path="server.jar"

# 判断文件是否存在
if [ -f "$file_path" ]; then
    echo "server file exists, prepare running"
else
    echo "server file does not exist, prepare downloading"
    ./image-tools/download.sh "$SERVER_VERSION"
    chmod +x server.jar
fi

echo "Starting server..."
java -Xmx"$SERVER_XMX" -Xms"$SERVER_XMS" -jar "$SERVER_JAR" nogui