# Go Project Setup

This guide provides instructions for setting up the Go development environment and installing necessary tools.

## 1. Installing Go

### macOS
The easiest way to install Go on macOS is using [Homebrew](https://brew.sh/):
```bash
brew install go
```

### Windows
You can install Go using [winget](https://learn.microsoft.com/en-us/windows/package-manager/winget/) or by downloading the installer from the [official website](https://go.dev/dl/).

Using winget (in Bash/Git Bash):
```bash
winget install GoLang.Go
```

## 2. Installing Development Tools

Run the following commands in your terminal (Bash) to install the required tools for an enhanced Go development experience:

```bash
# Go Language Server (LSP)
go install golang.org/x/tools/gopls@latest

# Generate unit tests
go install github.com/cweill/gotests/gotests@latest

# Modify struct field tags
go install github.com/fatih/gomodifytags@latest

# Generate method stubs for implementing an interface
go install github.com/josharian/impl@latest

# Run Go code snippets
go install github.com/haya14busa/goplay/cmd/goplay@latest

# Extract JSON schema of Go source file for outlining
go install github.com/ramya-rao-a/go-outline@latest

# Go debugger (Delve)
go install github.com/go-delve/delve/cmd/dlv@latest

# Go linter and static analysis
go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 3. Environment Configuration

Ensure that your Go binary path and tool binary path are added to your shell's configuration (e.g., `.bashrc`, `.zshrc`, or Windows Environment Variables).

### macOS / Linux (Bash/Zsh)
Add these lines to your `~/.bashrc` or `~/.zshrc`:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

### Windows (Bash/Git Bash)
If you are using Git Bash on Windows, add these to your `~/.bash_profile` or `~/.bashrc`:
```bash
export GOPATH=$HOME/go
export PATH=$PATH:"/c/Program Files/Go/bin":$GOPATH/bin
```

---
*Note: After installing, you may need to restart your terminal for the changes to take effect.*
