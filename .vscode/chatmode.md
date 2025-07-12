# Chat Mode Configuration

This file contains settings and guidelines for configuring chat mode in the project.

## Purpose
Chat mode enables dynamic interaction with MCP servers and tools, allowing users to switch configurations and manage workflows efficiently.

## Configuration
- Use `mcp.json` to define server settings and tool groups.
- Separate configuration files can be created for different environments or workflows.

## Example
```jsonc
{
  "servers": {
    "github": {
      "url": "https://api.githubcopilot.com/mcp/"
    },
    "git-local": {
      "command": "uvx",
      "args": ["mcp-server-git"]
    }
  },
  "toolGroups": {
    "default": ["microsoft-docs"],
    "git-remove": ["github"],
    "git-local": ["git-local"],
    "demo": ["time", "echo", "duckduckgo"]
  }
}
```

## Notes
- Ensure the MCP server is running before using chat mode.
- Use environment variables or scripts to load specific configurations dynamically.

## Troubleshooting
If you encounter errors, verify:
- The MCP server is correctly configured and running.
- The selected configuration file matches your workflow requirements.
