# Minecraft Server Images

This repository is used for automatically building and uploading Minecraft server images, which can be deployed on Docker or Kubernetes clusters.

**Automatic Build Process**: This repository contains an automated GitHub Actions workflow that triggers on each push to the repository, automatically building the specified version of the Minecraft server image and pushing it to [Docker Hub](https://hub.docker.com/r/ceerdecy/minecraft-server).

**Multi-platform Support**: By using Docker Buildx, we are able to build images for multiple platforms (such as Linux's amd64 and arm64 architectures), ensuring that you can find a suitable image version regardless of the hardware your cluster is running on.

**Configuration and Usage**: To deploy the Minecraft server in a Docker or Kubernetes environment, you need to configure the correct port mapping and environment variables. The default port inside the container is `25565`, which you need to bind to a port on the host machine.

**Environment Variables**: Here are some commonly used environment variables and their purposes:

- `SERVER_VERSION`: Specifies the version of the Minecraft server to build. (only dynamic-version)

**End User License Agreement (EULA)**: For ease of server setup, all images in this repository have the `EULA` set to `true` by default, <u>which means that by using the images from this repository, you are accepting the Minecraft End User License Agreement.</u>

**Maintenance and Updates**: This repository is regularly updated to support the latest versions of the Minecraft server. If you need a specific version of the image, you can create a branch or submit a request to add a new version.

**Contributions**: We welcome contributions from the community. If you find issues or have suggestions for improvement, please help us improve this project by submitting issues or pull requests.

**Disclaimer**: Please ensure you comply with Minecraft's terms of use and relevant laws and regulations.