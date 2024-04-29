# GOBIN

**What is GOBIN?**

- `GOBIN` is an environment variable used in the Go programming language.
- It determines the location where compiled Go application binaries (executables) are placed after building the main package with the `go build` or `go install` commands.

**Default Behavior:**

- If `GOBIN` is not explicitly set, Go uses the following default directories:
    - **With `GOPATH` set:** `$GOPATH/bin`
    - **Without `GOPATH` set:** `$HOME/go/bin` (assuming you have a Go installation directory named `go` in your home directory)

**Why Use GOBIN?**

- You might want to customize the installation location for binaries for various reasons:
    - **Project-Specific Binaries:** For a project, you might want compiled binaries to go into a separate directory within your project structure.
    - **Centralized Binaries:** You could set `GOBIN` to a central location on your system to have all Go binaries in one place.
    - **Different User Accounts:** If you work on multiple user accounts, setting `GOBIN` per user can keep binaries organized for each account.

**Setting GOBIN**

1. **Temporary Setting:**
   - Use the `export GOBIN=desired/location` command in your terminal to set `GOBIN` for the current session.
2. **Persistent Setting:**
   - Edit your shell profile file (e.g., `~/.bashrc` or `~/.zshrc`) and add the line `export GOBIN=desired/location`.
   - This makes the setting apply every time you open a new terminal window.

**Adding GOBIN to PATH**

- To run compiled Go binaries without specifying their full path, you need to add the `GOBIN` directory to your system's `PATH` environment variable.
- Edit your shell profile file and append `:$GOBIN` to the existing `PATH` variable.
- Example: `export PATH=$PATH:$GOBIN`

**Example (Centralized Binaries):**

1. Set `GOBIN` to a central directory like `~/bin`:
   - Add the line `export GOBIN=$HOME/bin` to your shell profile.
2. Add `$HOME/bin` to your PATH:
   - Append `:$HOME/bin` to the `PATH` variable in your profile.

**Important Notes:**

- Make sure the directory you choose for `GOBIN` exists and has write permissions for your user.
- Setting `GOBIN` affects only the location where `go install` places binaries. Manually built binaries (with `go build`) will go to the current working directory by default.
- Go introduced modules in Go 1.11, which are the preferred way to manage dependencies. `GOPATH` is still used for downloaded dependencies, but `GOBIN` retains its functionality for installing binaries.
