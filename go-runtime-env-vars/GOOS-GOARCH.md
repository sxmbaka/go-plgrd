# GOOS and GOARCH

**Purpose**: These variables are used when cross-compiling Go code for different operating systems and architectures.

**Usage**: Set GOOS to your target operating system (e.g., "linux," "windows," "darwin") and GOARCH to your target architecture (e.g., "amd64," "arm") when running the go build or go install command.

**Benefits**: Cross-Platform Compatibility — These variables allow you to compile your Go code for different operating systems and architectures, increasing the portability of your applications.

```bash
# Cross-compile for Linux on a Windows machine (GOOS=linux, GOARCH=amd64)
GOOS=linux GOARCH=amd64 go build -o myapp-linux main.go
```