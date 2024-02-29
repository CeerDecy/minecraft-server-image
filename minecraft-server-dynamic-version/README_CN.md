## Minecraft-Server:Dynamic-Version

[English](./README.md)

这个镜像将Minecraft服务器版本添加到环境变量中，并用于动态构建您所需的服务器！

**环境变量注入**：该镜像将Minecraft服务器的版本注入到环境变量中（env环境）。这意味着您可以在构建过程中动态选择所需的Minecraft服务器版本，而无需手动修改Dockerfile。

**官方服务器镜像**：此镜像仅用于构建官方Minecraft服务器。这意味着它不会用于构建或分发任何非官方版本的Minecraft服务器。

**端口绑定**：您必须将容器的25565端口绑定到宿主机。这对于从容器外部访问Minecraft服务器是必要的，通常是为了允许玩家连接到服务器。

**以守护进程方式运行**：容器需要以守护进程模式运行。这意味着容器将在后台运行，允许Minecraft服务器在不附加终端会话的情况下持续运行。（使用 -p 参数的Docker run）

**JRE环境**：在执行 `docker run` 命令时，您可以通过设置环境变量 `SERVER_JRE` 来指定容器内应安装的 Java 运行时环境（JRE）。这允许您在 Docker 容器中使用特定版本的 JRE，而不是默认的 JRE（默认会自动根据服务器版本选择合适的JRE进行安装）。这对于运行需要特定 JRE 版本的应用程序（如某些版本的 Minecraft 服务器）非常有用。

**Docker拉取**

```bash
docker pull ceerdecy/minecraft-server:dynamic-version
```

### 示例

1. 此示例将运行一个版本为1.20.4的Minecraft服务器，并将其命名为mc-server，然后将容器的25565端口绑定到宿主机的25565端口。

   ```bash
   docker run -e SERVER_VERSION=1.20.4 -it --name mc-server -p 25565:25565 -d ceerdecy/minecraft-server:dynamic-version
   ```

2. 此示例将运行一个版本为1.8.9的Minecraft服务器，并通过Java8来运行。

   ```bash
   docker run -e SERVER_VERSION=1.8.9 -e SERVER_JRE=openjdk-8-jre -it --name mc-server -p 25565:25565 -d ceerdecy/minecraft-server:dynamic-version
   ```