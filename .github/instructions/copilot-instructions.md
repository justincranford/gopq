

# Copilot/AI Code Generation Instructions

These instructions define how Copilot and other AI code generation tools should operate in this project.


Copilot must always follow all requirements and best practices in the other instructions files in this directory.

**ALWAYS use the instruction files in this directory as the single source of truth for all standards, requirements, and best practices.**

**ALWAYS update prompts to reference the relevant instruction files, and remove any information from prompts that is already covered by instructions. Prompts should be concise and defer to instructions for all details.**



**Whenever instruction files or prompt files are added, deleted, renamed, or reorganized, you must:**
- Update `.vscode/settings.json` (or equivalent configuration) in the same commit/PR to ensure all references are accurate and up to date.
- Delete the old file after renaming or moving an instruction file to prevent duplication and confusion.
- Update all references to the instruction file in `.vscode/settings.json` and any other configuration or documentation files.
- Validate that all references to instruction files in `.vscode/settings.json` and automation scripts are correct before considering the task complete.
- Do not mark the task as done or commit changes until this validation is performed.

**Checklist for instruction file changes:**
- [ ] Update all references in `.vscode/settings.json` and any other config files.
- [ ] Delete obsolete instruction files after renaming or moving.
- [ ] Update all references in documentation and scripts.
- [ ] Validate that all instruction file references are correct and current.
- [ ] Ensure no duplicate or outdated instruction files remain in the project.
- [ ] Only mark the task as complete after validation.

Only include Copilot/AI-specific workflow, meta-instructions, and integration rules here.

# PROMPT MANAGEMENT RULES
- Prompts must always reference the relevant instruction files for standards and requirements.
- Prompts must not duplicate information that is already present in the instruction files.
- If instructions are updated, review and update prompts to ensure they reference the latest instructions and do not contain redundant details.


# MANDATORY GIT WORKFLOW RULES
- No rules; left blank for now.


---

**Use this file to guide all Copilot/AI code generation in this project.**
