# GOFLAGS

**Purpose**: This variable allows you to specify additional flags for the go command.

**Usage**: Set GOFLAGS to include additional command-line flags when running go commands.

## Examples uses of `GOFLAGS`:
1. **Enabling Verbose Output:**
    ```bash
    GOFLAGS=-v 
    go build main.go
    ```
    This command compiles the `main.go` program with the `-v` flag, which stands for “verbose.” It provides detailed information about the build process, including which files are being compiled and linked.

2. **Setting Output Directory:**
    ```bash
    GOFLAGS=-o myapp 
    go build main.go
    ```
    Here, the `-o` flag is used to specify the output file’s name as `myapp`. This can be handy when you want to customize the name of the generated binary.

3. **Building with Specific Tags:**
    ```bash
    GOFLAGS=-tags=mytag 
    go build main.go
    ```
    You can use the `-tags` flag to build your Go program with specific build tags. Build tags are a way to include or exclude certain parts of your code during compilation based on conditions defined in your source code.

4. **Cross-Compiling with `GOOS` and `GOARCH`:**
    ```bash
    GOFLAGS="-v -ldflags '-s -w' -trimpath" 
    GOOS=linux 
    GOARCH=amd64 
    go build main.go
    ```
    In this example, `GOFLAGS` is used to pass a combination of flags to customize the build process. It includes the `-v` flag for verbose output, `-ldflags` to specify linker flags, and `-trimpath` to trim the source path. Additionally, `GOOS` and `GOARCH` are set to cross-compile the program for Linux on an AMD64 architecture.

5. **Modifying Build Mode:**
    ```bash
    GOFLAGS="-buildmode=shared" 
    go build main.go
    ```
    Here, `GOFLAGS` is used to specify the build mode as “shared.” This can be useful when you want to create a shared library or plugin using Go.

6. **Running Tests with Parallelism:**
    ```bash
    GOFLAGS="-p 2" 
    go test ./…
    ```
    You can control the parallelism of your Go test runs using the `-p` flag. In this example, `GOFLAGS` is set to run tests with a maximum of 2 parallel processes.

7. **Running Benchmarks with Custom Timing:**
    ```bash
    GOFLAGS="-benchtime=2s" 
    go test -bench=. ./…
    ```
    The `-benchtime` flag allows you to specify the duration for which benchmark tests should run. Here, `GOFLAGS` is used to set the benchmark time to 2 seconds.

8. **Using Custom Environment Variables:**
    ```bash
    GOFLAGS="GOARCH=arm64 GOOS=linux" 
    go build main.go
    ```
    You can use `GOFLAGS` to set custom environment variables for the Go build process. In this case, it specifies the `GOARCH` and `GOOS` variables to cross-compile for ARM64 architecture on Linux.