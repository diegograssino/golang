---
description: Rules for committing code to this repository
---
# Commit Rules

This repository strictly follows Conventional Commits to maintain a clean and readable history.

When committing code or proposing commits to the user, format the messages exactly as follows:
`<type>[optional scope]: <description>`

**Allowed Types:**
- `feat`: A new feature or a new learning project folder
- `fix`: A bug fix or correction in the code
- `docs`: Documentation only changes (e.g., updating the readme or adding comments)
- `style`: Changes that do not affect the meaning of the code (formatting, etc.)
- `refactor`: A code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests or correcting existing tests
- `chore`: Changes to the build process, `.gitignore`, or auxiliary IDE settings

**Examples:**
- `feat(hello-golang): initialize project with uuid package`
- `docs(readme): add new project to roadmap table`
- `fix(variables): correct type conversion in assignment`
