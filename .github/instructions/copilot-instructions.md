

# Copilot/AI Code Generation Instructions

These instructions define how Copilot and other AI code generation tools should operate in this project.

Copilot must always follow all requirements and best practices in the other instructions files in this directory.

Only include Copilot/AI-specific workflow, meta-instructions, and integration rules here.


# MANDATORY GIT WORKFLOW RULES
- ALWAYS use PowerShell terminal and Windows command syntax for all git and code quality commands, regardless of the user's OS or default shell.
- ALWAYS run `golangci-lint run --fix` and `gofumpt -l -w .` before every commit, and fix all warnings and errors before continuing with the commit, as part of a single chained command. Use `golangci-lint run --fix` to automatically fix all fixable lint issues.
- NEVER use Git Bash or Unix shell syntax for any git or code quality operation.
- The canonical commit command is:
  ```
  golangci-lint run --fix; gofumpt -l -w .; git add -A; git commit -m "..."; git push
  ```
- These rules override all other shell/environment logic.


---

**Use this file to guide all Copilot/AI code generation in this project.**
