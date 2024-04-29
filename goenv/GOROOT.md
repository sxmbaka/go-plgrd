`GOROOT` is an environment variable used in the Go programming language to specify the root directory where the Go SDK (Software Development Kit) is installed on your system. This variable signifies the base location of the Go installation.

When you install Go on your machine, whether using the official binary distribution or through a custom installation, Go assumes a default location for its installation directory. For example:

- On Linux and macOS, the default installation directory is typically `/usr/local/go`.
- On Windows, it's often `C:\Go`.

However, there are scenarios where you might want to install Go in a different location:

1. **Custom Installation Directory**:
   You may choose to install Go in a directory other than the default location. This could be for organizational preferences or to isolate different versions of Go.

2. **Multiple Go Versions**:
   If you need to maintain multiple versions of Go on your system, setting `GOROOT` allows you to switch between different installations easily.
   For example if you have installed GO to the location ~/Documents/go on Linux/MAC platform, then make below entry to the ~/$HOME/.profile file.

    ```bash
    export GOROOT=~/Documents/go
    export PATH=$PATH:$GOROOT/bin
    ```

Setting `GOROOT` tells the Go tools where to find the core files and libraries that make up the Go runtime, standard library, and development tools. This environment variable is essential for Go to function correctly because it helps the Go tools locate the necessary resources during compilation, linking, and execution of Go programs.

To summarize, `GOROOT` is used to specify the location of the Go SDK on your system, and it is crucial for correctly configuring your Go development environment.