# 贡献于 Minecraft 服务器镜像

[English](CONTRIBUTING.md)

我们欢迎对 Minecraft 服务器镜像项目的贡献！无论你是一位经验丰富的开发者还是一个好奇的初学者，都有许多方式可以让你参与进来，帮助这个项目变得更好。以下是帮助你开始的一个指南。

## 行为守则

本项目遵循一套行为守则。通过参与，你需要遵守这一守则。请向项目维护者报告任何不可接受的行为。

## 如何贡献

### 报告问题

- 确保该问题尚未被报告。
- 在你的报告中包含尽可能多的细节。
- 如果可能，提供一个实时示例或问题的复现。

### 提出增强建议

- 使用清晰且具有描述性的标题。
- 提供增强建议的逐步描述。
- 解释当前的行为和预期的行为。

### 你的第一次代码贡献

- 在 GitHub 上 fork 仓库。
- 阅读项目文档。
- 运行测试以确保所有测试通过。
- 在新分支中进行你的更改。
- 确保你的更改遵循项目的编码风格。
- 推送到你的 fork 并提交一个 pull request。

### Pull Requests

- 确保你的 PR 与基础分支保持最新。
- 清晰且简洁地描述你的更改。
- 在 PR 标题中包含问题编号（例如，“Closes #123”）。
- 遵循 `.github/PULL_REQUEST_TEMPLATE.md` 中提供的 PR 模板。

### 附加资源

- [GitHub 通用文档](https://help.github.com/)
- [GitHub pull request 文档](https://help.github.com/articles/about-pull-requests/)
- [Markdown 指南](https://guides.github.com/features/mastering-markdown/)

## Dockerfile 和非官方镜像指南

### 在 Dockerfile 中添加你的 ID

在为项目贡献时，请确保在 Dockerfile 的 `AUTHORS` 部分中包含你的 GitHub ID。这有助于给贡献者以信用，并保持更改记录的准确性。

### 构建非官方镜像

如果你正在构建一个非官方的 Minecraft 服务器镜像，请遵循以下指南：

- **完整性**：确保所有必要的文件，如 mods、插件和其他依赖项都包含在镜像中。不要假设用户会有这些文件或知道在哪里获取它们。
- **测试**：彻底测试你的镜像以确保它按预期工作。这包括启动服务器、检查错误以及验证所有功能是否正常运行。
- **命名规则**：为你的镜像使用清晰且具有描述性的命名规则。对于服务器镜像，格式应该是 `minecraft-server-[服务器类型]-[游戏版本]`。对于 modpack 或集成服务器镜像，使用 `minecraft-server-[modpack名称]-[modpack版本]`。

示例：
- 版本 1.16.4 的纯 Minecraft 服务器镜像将被命名为 `minecraft-server-vanilla-1.16.4`。
- 名为 "SkyFactory 4" 的 modpack 版本 4.0.0 的镜像将被命名为 `minecraft-server-SkyFactory-4.0.0`。

### 附加提示

- 尽可能保持你的 Dockerfile 简单且易于阅读。避免使用不必要的层或复杂命令。
- 在镜像的 README 或 Dockerfile 本身中记录任何特殊的构建指令或使用说明。
- 始终使用 Minecraft 版本和你的 GitHub 用户名来标记你的镜像（例如，`yourusername/minecraft-server-vanilla-1.16.4`）。

通过遵循这些指南，你将帮助维护此仓库中镜像的质量和一致性。感谢你的贡献！

## 许可证

通过为这个项目贡献，你同意你的贡献将根据其 [LICENSE](LICENSE) 许可证进行许可。

---

感谢你对贡献的兴趣！我们期待你的贡献。