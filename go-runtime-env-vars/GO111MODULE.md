# GO111MODULE

The `GO111MODULE` environment variable is used to control the behavior of Go modules, which were introduced in Go 1.11 as a way to manage dependencies and versioning within Go projects. This variable influences how Go handles module-aware mode, affecting whether and how Go modules are used in a project.

### Values of `GO111MODULE`

The `GO111MODULE` variable can be set to one of three values:

1. **`GO111MODULE=off`**:
   - Disables Go modules entirely.
   - When `GO111MODULE=off`, Go operates in the GOPATH mode.
   - Go does not use or require a `go.mod` file.
   - Dependencies are managed using the traditional GOPATH-based approach.

2. **`GO111MODULE=on`**:
   - Enables Go modules for all projects.
   - When `GO111MODULE=on`, Go always uses Go modules, regardless of whether a `go.mod` file is present in the project directory.
   - This is useful for enforcing module usage across all projects.

3. **`GO111MODULE=auto`** (default behavior):
   - This is the default mode when Go modules are detected in the project directory.
   - Go uses modules if the current directory contains a `go.mod` file or is below a directory containing a `go.mod` file.
   - If no `go.mod` file is found, Go falls back to GOPATH mode.
   - This mode is flexible and allows projects to use Go modules only when needed.

### Working with `GO111MODULE`

- **Enabling Go Modules for a Project**:
  - To start using Go modules in a project, navigate to the project directory and initialize modules using `go mod init`.
  - This will create a `go.mod` file that defines the project's module and its dependencies.

- **Using `GO111MODULE=auto`**:
  - This is the recommended setting for projects using Go modules.
  - It allows Go to automatically determine whether to use modules based on the presence of a `go.mod` file.
  - If you're working on a project with Go modules, set `GO111MODULE=auto` in your environment or shell configuration.

### Why Use Go Modules?

- **Dependency Management**:
  - Go modules provide a standardized way to manage dependencies and versioning.
  - Modules specify the exact versions of dependencies required by a project, ensuring reproducibility across builds.

- **Versioning**:
  - Modules allow projects to specify which versions of dependencies are compatible, preventing version conflicts and ensuring reliable builds.

- **Simplicity and Maintainability**:
  - Go modules simplify dependency management compared to the GOPATH approach.
  - They make it easier to share and reuse code across projects while maintaining versioning and compatibility.

### Best Practices

- **Use `GO111MODULE=auto`**:
  - For projects using Go modules, set `GO111MODULE=auto` in your environment or shell configuration.
  - This ensures that Go modules are used when needed, without requiring explicit configuration for each project.

- **Explicitly Initialize Modules**:
  - Always initialize Go modules (`go mod init`) in a project directory to define the project's module and dependencies.

- **Update Dependencies**:
  - Use `go get` or `go mod tidy` to manage dependencies and update module versions as needed.
  - Keep `go.mod` and `go.sum` files up to date to reflect the current state of dependencies.

### Conclusion

The `GO111MODULE` environment variable plays a crucial role in how Go handles module-aware mode and dependency management. It provides flexibility in managing dependencies and allows projects to adopt Go modules seamlessly. By understanding and leveraging `GO111MODULE` effectively, developers can ensure reliable and maintainable Go projects with clear dependency management.