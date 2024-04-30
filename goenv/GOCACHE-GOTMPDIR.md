# GOCACHE and GOTMPDIR

**Purpose**: These variables manage Go’s build cache and temporary file locations.

**Usage**: Customize GOCACHE and GOTMPDIR to specify cache and temporary file storage.

**Benefits** :
- Customizing cache and temporary storage locations to improve build performance.
- Managing build artifacts efficiently.

```bash
# Customize the cache directory
GOCACHE=/path/to/cache  GOTMPDIR=/path/to/temp go build main.go
```
## How GOCACHE improves build performance ??
Basically GoCACHE stores and manages compiled package objects to improve build performance. It serves as a cache for intermediate build artifacts, preventing redundant compilation and improving the efficiency of the build process.

### Here’s how the Go build cache works:

1. **Compilation**: When you build a Go application or package, the Go compiler compiles the source code into object files (with a `.o` extension).

2. **Caching**: The Go build cache stores these compiled object files, along with related metadata, in a cache directory. By default, this cache directory is located in the user’s home directory (`$HOME/go/pkg`). However, you can customize the cache location using the `GOCACHE` environment variable.

3. **Reuse**: The next time you build the same package or import the same dependency, Go checks the cache directory to see if the compiled object files and metadata are already present. If they are, Go reuses these cached files instead of recompiling the source code. This significantly speeds up the build process, especially for large projects with many dependencies.

4. **Invalidation**: The cache is aware of changes in source code or build flags. If anything has changed, the cache is invalidated for that specific build, and Go will recompile the affected parts of the code.

### The Go build cache offers several benefits:

- **Faster Build**s: By reusing cached object files, Go builds are much faster, especially in large projects or in scenarios where dependencies don’t change frequently.

- **Reduced CPU and Memory Usage**: Caching reduces the CPU and memory resources required for compiling the same code repeatedly.

- **Consistent Builds**: The cache ensures that the same code always produces the same compiled results, improving build consistency.

- **Isolation**: Cached files are stored separately for different Go versions and different build configurations, preventing conflicts.

- **Customization**: You can customize the cache location and other cache-related settings using environment variables like `GOCACHE`.

### What are build artifacts ??

Build artifacts are the files and outputs generated as a result of compiling and building a software project. These artifacts are created during the build process and typically include:

1. **Executable Binaries**: For compiled languages like Go, C, C++, and Rust, the primary build artifact is the executable binary. This is the file that contains the machine code instructions that the computer’s processor can execute. It’s the final output of the build process, representing the compiled version of your source code that can be run as a program.

2. **Library Files**: In some cases, especially in languages like C and C++, build artifacts include shared libraries (e.g., DLLs on Windows, shared objects on Linux) or static libraries (e.g., .a files on Unix-based systems). These libraries contain reusable code that can be linked to other programs or libraries.

3. **Intermediate Files**: During the compilation process, various intermediate files are generated. These files are not meant to be directly executed but are crucial for creating the final executable. Examples include object files (with .o or .obj extensions) and preprocessed source code.

4. **Configuration Files**: In certain build systems, configuration files may be generated as build artifacts. These files can contain information about build settings, environment variables, or metadata related to the build.

5. **Documentation**: Build processes can also generate documentation as artifacts. This can include API documentation, user manuals, or any other documentation associated with the software project.

6. **Test Results**: When running tests as part of the build process, the results of those tests can be considered build artifacts. These results may include test logs, reports, and coverage information.

7. **Packaged Distributions**: In some cases, the build process produces packaged distributions of the software, such as installer files (e.g., .msi on Windows, .dmg on macOS), tarballs, zip files, or container images. These distributions are often used for deployment and distribution purposes.

Build artifacts are essential in the software development lifecycle because they represent the tangible output of the development and build processes. These artifacts are typically stored, versioned, and distributed as part of software releases. They allow developers to share, deploy, and run their software on various environments, making them a critical component of software development and delivery.