# Commit Instructions

These instructions define the required workflow and best practices for making commits in this project.

## MANDATORY GIT WORKFLOW RULES
- ALWAYS use PowerShell terminal and Windows command syntax for all git and code quality commands, regardless of the user's OS or default shell.
- ALWAYS run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit, and fix all warnings and errors before continuing with the commit, as part of a single chained command. Use `golangci-lint run --fix` to automatically fix all fixable lint issues.
- NEVER use Git Bash or Unix shell syntax for any git or code quality operation.
- The canonical commit command is:
  ```
  golangci-lint run --fix; gofumpt -l -w .; git add -A; git commit -m "..."; git push
  ```
- These rules override all other shell/environment logic.
- All code must be reviewed and pass CI checks (lint, format, test, fuzz, benchmark) before merging.
- Use descriptive commit messages and PR descriptions.
- Always check for the latest Go, linter, and dependency versions before starting new work.

---

**Use this file to guide all commit and workflow operations in this project.**
