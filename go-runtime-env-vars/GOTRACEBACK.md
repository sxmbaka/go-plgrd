## GOTRACEBACK

**Purpose**: This environment variable controls the level of detail in the panic trace generated when a Go program encounters a runtime panic.

**Usage**: Set `GOTRACEBACK` to adjust the verbosity of panic output.

### Usage in Go 1.5

In Go 1.5, `GOTRACEBACK` had four valid values:

- `GOTRACEBACK=0`: Suppresses all tracebacks, showing only the panic message.
- `GOTRACEBACK=1`: Default behavior, shows stack traces for all goroutines but suppresses runtime-related stack frames.
- `GOTRACEBACK=2`: Similar to `1`, but includes frames related to the runtime (revealing runtime-created goroutines).
- `GOTRACEBACK=crash`: Same as `2`, but triggers a segfault instead of exiting cleanly.

### Example

```bash
# Enable detailed panic traces
GOTRACEBACK=all go run main.go
```

### Changes in Go 1.6 and Later

Starting with Go 1.6, the interpretation of `GOTRACEBACK` values changed:

- `GOTRACEBACK=none`: Suppresses all tracebacks, showing only the panic message.
- `GOTRACEBACK=single` (new default): Prints the stack trace only for the goroutine believed to have caused the panic.
- `GOTRACEBACK=all`: Shows stack traces for all goroutines but suppresses runtime-related frames.
- `GOTRACEBACK=system`: Similar to `all`, but includes frames related to the runtime.

For compatibility:
- `GOTRACEBACK=0` in Go 1.6+ maps to `none`.
- `GOTRACEBACK=1` maps to `all`.
- `GOTRACEBACK=2` maps to `system`.

### Example Program

```go
package main

func main() {
    panic("kerboom")
}
```

Compiling and running this program with `GOTRACEBACK=0` (or `none` in Go 1.6+) will suppress all goroutine stack traces, displaying only the panic message.

### Go 1.22.2 Update

For Go 1.22.2, ensure that you refer to the latest Go documentation for any potential updates or changes in behavior related to `GOTRACEBACK`. Always use the appropriate `GOTRACEBACK` setting based on your debugging or diagnostic needs when encountering panics or unexpected program crashes in Go.