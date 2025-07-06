# Commit Instructions

These instructions define the required workflow and best practices for making commits in this project.


## MANDATORY GIT WORKFLOW RULES
- ALWAYS use PowerShell terminal and Windows command syntax for all git and code quality commands, regardless of the user's OS or default shell.
- NEVER use Git Bash or Unix shell syntax (such as `&&`, `;`, or Unix-style pipes) for any git or code quality operation. Use only PowerShell-compatible command separators (such as `;`) and syntax.
- When providing commit or workflow instructions, always show the correct PowerShell command format. For example, use `git add .; git commit -m "message"` instead of `git add . && git commit -m "message"`.
- These rules override all other shell/environment logic.
- All code must be reviewed and pass CI checks (lint, format, test, fuzz, benchmark) before merging.
- Use descriptive commit messages and PR descriptions.

---

**Use this file to guide all commit and workflow operations in this project.**
