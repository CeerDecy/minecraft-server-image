# Contributing to Minecraft Server Images

[中文](CONTRIBUTING_CN.md)

[Build Tutorial](https://etov.club/index.php/2024/02/build-minecraft-server-image/)

We welcome contributions to the Minecraft Server Images project! Whether you're a seasoned developer or a curious beginner, there are many ways to get involved and help make this project better. Here's a guide to help you get started.

## Code of Conduct

This project adheres to a code of conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to the project maintainers.

## How to Contribute

### Reporting Bugs

- Ensure the bug has not already been reported.
- Include as much detail as possible in your report.
- If possible, provide a live example or a reproduction of the problem.

### Suggesting Enhancements

- Use a clear and descriptive title for the issue.
- Provide a step-by-step description of the suggested enhancement.
- Explain the current behavior and the expected behavior.

### Your First Code Contribution

- Fork the repository on GitHub.
- Read the project documentation.
- Run the tests to ensure they all pass.
- Make your changes in a new branch.
- Ensure your changes follow the project's coding style.
- Push to your fork and submit a pull request.

### Pull Requests

- Ensure your PR is up-to-date with the base branch.
- Describe your changes clearly and concisely.
- Include the issue number in the PR title (e.g., "Closes #123").
- Follow the PR template provided in `.github/PULL_REQUEST_TEMPLATE.md`.

### Additional Resources

- [General GitHub documentation](https://help.github.com/)
- [GitHub pull request documentation](https://help.github.com/articles/about-pull-requests/)
- [Markdown guide](https://guides.github.com/features/mastering-markdown/)

## Dockerfile and Non-Official Image Guidelines

### Adding Your ID to the Dockerfile

When contributing to the project, please make sure to include your GitHub ID in the `AUTHORS` section of the Dockerfile. This helps to give credit to the contributors and maintain an accurate record of changes.

### Building Non-Official Images

If you are building a non-official Minecraft server image, please follow these guidelines:

- **Completeness**: Ensure that all necessary files such as mods, plugins, and other dependencies are included in the image. Do not assume that users will have these files or know where to obtain them.
- **Testing**: Thoroughly test your image to ensure that it works as expected. This includes starting the server, checking for errors, and verifying that all features are operational.
- **Naming Convention**: Use a clear and descriptive naming convention for your images. For server images, the format should be `minecraft-server-[server-type]-[game-version]`. For modpacks or integrated server images, use `minecraft-server-[modpack-name]-[modpack-version]`.

Examples:
- A vanilla Minecraft server image for version 1.16.4 would be named `minecraft-server-vanilla-1.16.4`.
- A modpack image named "SkyFactory 4" with version 4.0.0 would be named `minecraft-server-SkyFactory-4.0.0`.

### Additional Tips

- Keep your Dockerfile as simple and readable as possible. Avoid using unnecessary layers or complex commands.
- Document any special build instructions or usage notes in the image's README or within the Dockerfile itself.
- Always tag your images with the Minecraft version and your GitHub username (e.g., `yourusername/minecraft-server-vanilla-1.16.4`).

By following these guidelines, you help maintain the quality and consistency of the images in this repository. Thank you for your contributions!

## License

By contributing to this project, you agree that your contributions will be licensed under its [LICENSE](LICENSE).

---

Thank you for your interest in contributing! We look forward to your contributions.