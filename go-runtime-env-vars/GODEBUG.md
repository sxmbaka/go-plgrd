# GODEBUG

The `GODEBUG` environment variable in Go is used to enable or configure various debugging-related features within the Go runtime. It provides a way to control runtime-specific behaviors and obtain additional debugging information when running Go programs.

Here's an overview of its purpose and usage:

**Purpose**:
- Enable specific debugging features or behaviors within the Go runtime.
- Obtain additional runtime information and diagnostics.

**Usage**:
- Set `GODEBUG` to a comma-separated list of key-value pairs to enable specific debug features.
- The format is `key1=value1,key2=value2,...`.

**Common Debugging Features**:
- **`gctrace=1`**: Enables garbage collection (GC) trace output, providing information about GC cycles.
- **`allocfreetrace=1`**: Enables allocation and free trace output, useful for tracking memory allocations and deallocations.
- **`schedtrace=1000`**: Enables scheduler tracing with a specified interval (in microseconds).
- **`traceback=1`**: Enables stack traces for all goroutine transitions.
- **`cgocheck=1`**: Enables checks for invalid use of cgo.
- **`invalidptr=1`**: Enables panics for invalid memory accesses.

**Example**:
```bash
# Enable GC trace output and scheduler tracing
GODEBUG=gctrace=1,schedtrace=1000 go run main.go
```

In this example:
- `gctrace=1` enables GC trace output, providing insights into garbage collection cycles.
- `schedtrace=1000` enables scheduler tracing, generating trace information at intervals of 1000 microseconds.

**When to Use**:
- Use `GODEBUG` when you need detailed runtime information or debugging output.
- Helpful for diagnosing performance issues related to memory management, concurrency, and scheduling.
- Use cautiously as enabling certain debugging features can impact performance.

The `GODEBUG` environment variable is a powerful tool for gaining insights into the behavior of Go programs at runtime and diagnosing complex issues. However, it's recommended to refer to the official documentation for specific key-value pairs and their implications on program behavior.

## Further Reading
https://dave.cheney.net/tag/godebug