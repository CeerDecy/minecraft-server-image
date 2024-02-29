FROM ubuntu:latest
LABEL authors="ceerdecy"

WORKDIR /server

COPY . .

ENV SERVER_XMX 1024M
ENV SERVER_XMS 1024M
ENV SERVER_JAR minecraft_server.1.20.3.jar

RUN sed -i s/deb.debian.org/mirrors.aliyun.com/g /etc/apt/sources.list  \
    && apt-get update -y  \
    && apt-get install -y openjdk-17-jre curl  \
    && rm -rf /var/lib/apt/lists/*

RUN chmod +x run.sh && chmod +x minecraft_server.1.20.3.jar

EXPOSE 25565

ENTRYPOINT ["./run.sh"]