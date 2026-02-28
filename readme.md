# Learning Go Through Projects

Welcome to this repository for learning the Go programming language through practical, isolated projects!

Each core concept is organized into its own numbered directory. Inside each directory, you'll find a `readme.md` file that provides a description of the project and the requirements needed to construct it from scratch for learning purposes.

## Quick Setup (If you don't want to read it all)

### 1. Install Antigravity IDE

If you don't already have an IDE installed, we highly recommend setting up **Antigravity** (though you can use VS Code, Cursor, or any specialized editor of your taste!):

**macOS:**
```bash
brew tap google/antigravity && brew install --cask antigravity
```

**Windows:**
Download and run the installer from the [Antigravity releases page](https://github.com/google/antigravity).

### 2. Configure the Golang IDE Profile (Antigravity Only)

Once Antigravity is installed, instantly apply the Golang configuration profile (custom themes, extensions, formatting) straight into the IDE using your terminal:

**macOS:**
```bash
open "antigravity://profile/github/ae1804a7bbc9851299a28e06df5174e6"
```

**Windows (Git Bash / CMD):**
```bash
start antigravity://profile/github/ae1804a7bbc9851299a28e06df5174e6
```

### 3. Install Go & Development Tools

Copy and paste the appropriate bash script into your terminal to instantly install Go, the required Nerd Font, all essential Go development tools, and configure your path environment variables.

**macOS:**
```bash
brew install font-code-new-roman-nerd-font go && \
go install golang.org/x/tools/gopls@latest github.com/cweill/gotests/gotests@latest github.com/fatih/gomodifytags@latest github.com/josharian/impl@latest github.com/haya14busa/goplay/cmd/goplay@latest github.com/ramya-rao-a/go-outline@latest github.com/go-delve/delve/cmd/dlv@latest honnef.co/go/tools/cmd/staticcheck@latest && \
echo -e "\nexport GOPATH=\$HOME/go\nexport PATH=\$PATH:/usr/local/go/bin:\$GOPATH/bin" >> ~/.zshrc && source ~/.zshrc
```

**Windows (Git Bash):**
```bash
winget install GoLang.Go && \
go install golang.org/x/tools/gopls@latest github.com/cweill/gotests/gotests@latest github.com/fatih/gomodifytags@latest github.com/josharian/impl@latest github.com/haya14busa/goplay/cmd/goplay@latest github.com/ramya-rao-a/go-outline@latest github.com/go-delve/delve/cmd/dlv@latest honnef.co/go/tools/cmd/staticcheck@latest && \
echo -e '\nexport GOPATH=$HOME/go\nexport PATH=$PATH:"/c/Program Files/Go/bin":$GOPATH/bin' >> ~/.bash_profile && source ~/.bash_profile
```

---

## Projects & Roadmap

This repository is organized by folders, each containing a different project or stage of the learning path:

| Project         | Folder        | Description                                                       | Run Command |
| :-------------- | :------------ | :---------------------------------------------------------------- | :---------- |
| **Hello Golang** | [1_hello-golang](./1_hello-golang) | Basic "Hello Golang" implementation using the Google UUID package. | `cd 1_hello-golang && go run .` |

---

## Step-by-Step Setup (if you want to see detail of what are you doing)

### 1. Installing Go

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

### 2. Installing Development Tools

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

### 3. Environment Configuration

Ensure that your Go binary path and tool binary path are added to your shell's configuration (e.g., `.bashrc`, `.zshrc`, or Windows Environment Variables).

### macOS / Linux (Bash/Zsh)

Run the following commands to append these lines to your `~/.zshrc` (or replace `.zshrc` with `.bashrc` if you use Bash):

```bash
echo 'export GOPATH=$HOME/go' >> ~/.zshrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.zshrc
```

### Windows (Bash/Git Bash)

If you are using Git Bash on Windows, run the following commands to append these lines to your `~/.bash_profile` (or replace `.bash_profile` with `.bashrc`):

```bash
echo 'export GOPATH=$HOME/go' >> ~/.bash_profile
echo 'export PATH=$PATH:"/c/Program Files/Go/bin":$GOPATH/bin' >> ~/.bash_profile
```

### 4. IDE Profile Setup (VS Code, Antigravity & Cursor)

For a fully optimized Go development experience, you can import the custom **Golang IDE Profile**. This profile contains custom themes, formatting rules, UI preferences, and only the essential extensions for Go (without any frontend clutter).

**Import URL:** `https://gist.github.com/diegograssino/ae1804a7bbc9851299a28e06df5174e6`

### Manual Installation
1. Open your IDE (VS Code, Antigravity, or Cursor).
2. Go to **File > Preferences > Profiles > Import Profile...** (or search "Profiles: Import Profile" in the Command Palette `Ctrl+Shift+P` / `Cmd+Shift+P`).
3. Paste the URL above when prompted and press Enter.
4. Follow the prompts to create the new profile named **Golang**.
> **Note:** This profile includes highly opinionated configurations regarding the UI look and feel (such as specific themes, typography, and color customizations). Once imported, you can always revert the aesthetic choices back to your own defaults within the IDE's settings while still retaining all the necessary Golang extensions and language formatters.

### Optional: Install CodeNewRoman Nerd Font

If you want the exact typographic experience configured in the profile, you should install **CodeNewRoman Nerd Font Mono**.

**macOS:**
```bash
brew install --cask font-code-new-roman-nerd-font
```

**Windows:**
1. Download `CodeNewRoman.zip` from the [Nerd Fonts Releases page](https://github.com/ryanoasis/nerd-fonts/releases).
2. Extract the archive.
3. Select all the extracted `.ttf` files, right-click, and choose **"Install"**.

---

_Note: After installing, you may need to restart your terminal for the changes to take effect._
