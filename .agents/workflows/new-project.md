---
description: Create a new Go learning project
---
# Creating a New Go Project

When creating a new project in this repository, you MUST follow these steps exactly:

1. **Create the Folder**: Create a new folder in the root of the repository. The name MUST follow the format `[Number]_[kebab-case-name]` (e.g., `2_variables-and-types`). The number should increment based on the last project in the repository.

// turbo
2. **Initialize Go Module**: Navigate into the new folder and initialize a Go module with the name matching the folder.
`cd [folder-name] && go mod init [folder-name]`

3. **Structure**: Create a `main.go` file inside the new folder.

4. **Project README**: Create a `readme.md` file inside the new folder. This readme MUST serve as documentation for the project. Provide a `## Description` of what the project does, and a `## Requirements` section listing the exact steps a learner needs to fulfill to reconstruct the project from scratch.

5. **Update Root README**: You MUST update the "Projects & Roadmap" table in the `readme.md` in the root of the repository. Add the new project name, relative folder link, description, and the exact cross-platform run command (`cd [folder-name] && go run .`).
