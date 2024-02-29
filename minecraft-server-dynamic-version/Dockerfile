FROM ubuntu:latest
LABEL authors="ceerdecy"

WORKDIR /server

COPY . .

ENV SERVER_XMX 1024M
ENV SERVER_XMS 1024M
ENV SERVER_JAR server.jar

RUN sed -i s/deb.debian.org/mirrors.aliyun.com/g /etc/apt/sources.list  \
    && apt-get update -y  \
    && apt-get install -y curl  \
    && rm -rf /var/lib/apt/lists/*

RUN chmod +x image-tools/download.sh image-tools/parse-jre.sh image-tools/version-tool-linux-amd64 run.sh

EXPOSE 25565

ENTRYPOINT ["./run.sh"]