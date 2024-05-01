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
- Get more help
    ```bash
    go help env
    ```

Read further at:
- [A whirlwind tour of Go runtime environment variables - Dave Cheney](https://dave.cheney.net/2015/11/29/a-whirlwind-tour-of-gos-runtime-environment-variables)
- [Go your own way - Sourav Choudhary](https://medium.com/@souravchoudhary0306/go-your-own-way-customizing-go-with-environment-variables-3e47c880fe34)
- [The Go's official env.go file](https://go.dev/src/cmd/go/internal/envcmd/env.go)
- [Go environment variables explained in 5 mins - GolangDojo (YT)](https://youtu.be/Ut-NLq6d694?si=pjV1xE1R5Jycsc7r)
- [Golang Environment and Golang Command - Wahyu Eko Hadi Saputro](https://wahyu-ehs.medium.com/golang-environment-and-golang-command-1fdcbc145f32)

The following environment variables (`$name` or `%name%`, depending on the host operating system) control the run-time behavior of Go programs. The meanings and use may change from release to release.
</details>

1. [GOROOT](./go-runtime-env-vars/GOROOT.md)
2. [GOPATH](./go-runtime-env-vars/GOPATH.md)
3. [GOBIN](./go-runtime-env-vars/GOBIN.md)
4. [GO111MODULE](./go-runtime-env-vars/GO111MODULE.md)
5. [GOMOD](./go-runtime-env-vars/GOMOD.md)
6. [GOOS and GOARCH](./go-runtime-env-vars/GOOS-GOARCH.md)
7. [GOCACHE and GOTMPDIR](./go-runtime-env-vars/GOCACHE-GOTMPDIR.md)
8. [GOPROXY](./go-runtime-env-vars/GOPROXY.md)
9. [GOFLAGS](./go-runtime-env-vars/GOFLAGS.md)
10. [GOENV](./go-runtime-env-vars/GOENV.md)
11. [GOGC](./go-runtime-env-vars/GOGC.md)

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