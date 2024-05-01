<img src="assets\images\go-logo-blue.png">
<hr>

### [`General`](./general-conventions/)

1. [Hello World in Go](./general/hello-world.md)
2. [Semicolons in Go](./general-conventions/semicolons.md)
3. [Short Declarations](./general/short-declarations.md)<br>
(Using the `:=` operator)
4. [iota in Go](./general/iota.go)
5. [Enumerations in Go](./general/enums.md)

### [`GO Runtime Environment Variables`](./go-runtime-env-vars/)

<details>
    <summary><b>What are Go Environment Variables?</b></summary>

The `go env` command prints the value of environment variables used by the Go tools. These variables are used to configure the behavior of the Go tools and the Go runtime. These runtime environment variables are used to control the behavior of the Go runtime, and the Go tools. The Go tools include the `go` command, the `gofmt` command, and the `godoc` command. The Go runtime is the part of the Go toolchain that executes Go programs.

**Manipulating Go runtime environment variables:**
- Show all environment variables
    ```bash
    go env
    ```
- Show a specific environment variable
    ```bash
    go env [GOPATH]
    ```
- Set an environment variable to a value
    ```bash
    go env -w [GOBIN]=[path/to/directory]
    ```
- Reset an environment variable's value
    ```bash
    go env -u [GOBIN]
    ```
- Get help on a specific subcommand
    ```bash
    # tells you about the -w flag
    go help env -w
    ```
- Get more help
    ```bash
    # tells you about the go env command
    go help env
    ```
    ```bash
    # tells you about all the Go environment variables
    go help environment
    ```

Read further at:
- [A whirlwind tour of Go runtime environment variables - Dave Cheney](https://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables)
- [Go your own way - Sourav Choudhary](https://medium.com/@souravchoudhary0306/go-your-own-way-customizing-go-with-environment-variables-3e47c880fe34)
- [Go environment variables explained in 5 mins - GolangDojo (YT)](https://youtu.be/Ut-NLq6d694?si=pjV1xE1R5Jycsc7r)
- [Golang Environment and Golang Command - Wahyu Eko Hadi Saputro](https://wahyu-ehs.medium.com/golang-environment-and-golang-command-1fdcbc145f32)
- [Go runtime environment variables - Golang Docs](https://pkg.go.dev/runtime#hdr-Environment_Variables)
- [The Go's official env.go file](https://go.dev/src/cmd/go/internal/envcmd/env.go)

The following environment variables (`$name` or `%name%`, depending on the host operating system) control the run-time behavior of Go programs. The meanings and use may change from release to release. 

This arrangement lists the Go runtime environment variables in alphabetical order for easier reference and lookup.
</details>

1. [GOCACHE and GOTMPDIR](./go-runtime-env-vars/GOCACHE-GOTMPDIR.md)
2. [GOGC](./go-runtime-env-vars/GOGC.md)
3. [GO111MODULE](./go-runtime-env-vars/GO111MODULE.md)
4. [GOARCH and GOOS](./go-runtime-env-vars/GOOS-GOARCH.md)
5. [GOBIN](./go-runtime-env-vars/GOBIN.md)
6. [GOENV](./go-runtime-env-vars/GOENV.md)
7. [GOFLAGS](./go-runtime-env-vars/GOFLAGS.md)
8. [GOMAXPROCS](./go-runtime-env-vars/GOMAXPROCS.md)
9. [GOMOD](./go-runtime-env-vars/GOMOD.md)
10. [GOPATH](./go-runtime-env-vars/GOPATH.md)
11. [GOPROXY](./go-runtime-env-vars/GOPROXY.md)
12. [GOROOT](./go-runtime-env-vars/GOROOT.md)
13. [GOTRACEBACK](./go-runtime-env-vars/GOTRACEBACK.md)
14. [GODEBUG](./go-runtime-env-vars/GODEBUG.md)


### [`Variables`](./variables/)

1. [Variable Types](./variables/variable-types.go)
2. [Operations on variables](./variables/operations.go)<br>
(You might need to study [Operators](./operators/) in case of any anomalies)
3. [Type Conversions](./variables/type-conversions.go)
4. [Type Inference](./variables/type-inference.go)
5. [Type Inference](./variables/type-inference.go)<br>
(Read further at <https://go.dev/blog/type-inference>)
6. [Variable Naming Conventions](./variables/variable-naming-conventions.go)
7. [Variable Scopes](./variables/variable-scopes.go)
8. [Variable Shadowing](./variables/shadowing-variables.go)
9. [Constants in Go](./variables/constants.go)<br>
(Read further at <https://golangbyexample.com/constant-golang/>)

__Tip__: _Create a `playground` directory in the `root` directory and then test these examples yourself._