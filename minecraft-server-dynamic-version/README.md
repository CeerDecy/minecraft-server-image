## Minecraft-Server:Dynamic-Version

[中文](./README_CN.md)

This image adds the Minecraft Server version to the env environment and is used to dynamically build the Server you need! 

**Environment Variable Injection**: This image injects the version of the Minecraft Server into the environment variables (env environment). This means you can dynamically select the version of the Minecraft Server you need during the build process without manually modifying the Dockerfile.

**Official Server Image**: This image is only for building official Minecraft servers. This implies that it will not be used to build or distribute any unofficial versions of the Minecraft Server.

**Port Binding**: You must bind the container's port 25565 to the host machine. This is necessary for the Minecraft server to be accessible from outside the container, typically for players to connect to the server.

**Running as a Daemon**: The container needs to be run in daemon mode. This means that the container will run in the background, allowing the Minecraft server to operate continuously without being attached to a terminal session. (Docker run with -p)

**Docker pull**

```bash
docker pull ceerdecy/minecraft-server:dynamic-version
```

### Example

This example will run a Minecraft server with version 1.20.4 and name it mc-server, then bind the container's port 25565 to the host's port 25565.

```bash
docker run -e SERVER_VERSION=1.20.4 -it --name mc-server -p 25565:25565 -d ceerdecy/minecraft-server:dynamic-version
```

