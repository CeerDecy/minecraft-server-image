# Minecraft Server Images

本仓库用于自动构建和上传Minecraft服务器镜像，可以用于在部署在Docker或Kubernetes集群上。

**自动构建流程**：本仓库包含了一个自动化的GitHub Actions工作流，它在每次推送到仓库时触发，自动构建指定版本的Minecraft服务器镜像，并将其推送到[Docker Hub](https://hub.docker.com/r/ceerdecy/minecraft-server)中。

**多平台支持**：通过使用Docker Buildx，我们能够为多种平台（如Linux的amd64和arm64架构）构建镜像，确保无论您的集群运行在何种硬件上，都能找到合适的镜像版本。

**配置和使用**：为了在Docker或Kubernetes环境中部署Minecraft服务器，您需要配置正确的端口映射和环境变量。默认容器内的端口为`25565`，您需要将此端口与宿主机的某个端口进行绑定。

**环境变量**：以下是一些常用的环境变量及其用途：

- `SERVER_VERSION`：指定要构建的Minecraft服务器版本。（仅在dynamic-version可用）

**用户许可协议**：为了便于快速搭建服务器，本仓库所有镜像`EULA`默认为`true`，<u>这意味着一旦您使用本仓库中的镜像默认接受Minecraft的最终用户许可协议。</u>

**维护和更新**：本仓库定期更新，以支持最新的Minecraft服务器版本。如果您需要特定版本的镜像，可以创建一个分支或提交请求来添加新版本。

**贡献**：我们欢迎社区的贡献。如果您发现问题或有改进建议，请通过提交问题或拉取请求来帮助我们改进这个项目。

**免责声明**：请确保您遵守Minecraft的使用条款和相关法律法规。

