## Minecraft-Server:Dynamic-Version

[English](./README.md)

这个镜像将Minecraft服务器版本添加到环境变量中，并用于动态构建您所需的服务器！

**环境变量注入**：该镜像将Minecraft服务器的版本注入到环境变量中（env环境）。这意味着您可以在构建过程中动态选择所需的Minecraft服务器版本，而无需手动修改Dockerfile。

**官方服务器镜像**：此镜像仅用于构建官方Minecraft服务器。这意味着它不会用于构建或分发任何非官方版本的Minecraft服务器。

**端口绑定**：您必须将容器的25565端口绑定到宿主机。这对于从容器外部访问Minecraft服务器是必要的，通常是为了允许玩家连接到服务器。

**以守护进程方式运行**：容器需要以守护进程模式运行。这意味着容器将在后台运行，允许Minecraft服务器在不附加终端会话的情况下持续运行。（使用 -p 参数的Docker run）

**Docker拉取**

```bash
docker pull ceerdecy/minecraft-server:dynamic-version
```

### 示例

此示例将运行一个版本为1.20.4的Minecraft服务器，并将其命名为mc-server，然后将容器的25565端口绑定到宿主机的25565端口。

```bash
docker run -e SERVER_VERSION=1.20.4 -it --name mc-server -p 25565:25565 -d ceerdecy/minecraft-server:dynamic-version
```