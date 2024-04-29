# Hello World
In this tutorial, we will see how to setup Go environment and create a simple hello world program. We will also see how to run the program using different commands.
So far here is the value of [GOROOT](../goenv/GOROOT.md), [GOPATH](../goenv/GOPATH.md), and [GOBIN](../goenv/GOBIN.md).
```bash
echo $GOROOT will give installed directory for GO
echo $GOPATH will give ~/Desktop/go
echo $GOBIN will give ~/Desktop/go/bin
```
A typical program has `.go` file extension. Now let’s create a “Hello World” program. For that first create a directory named hello outside `$GOPATH/src`.  Since the project is getting creating outside `$GOPATH/src` we also need to create a `go.mod` file with import path as `sample.com/hello` .For now, let’s create a simple module so that we can see how the hello world program looks like in go. Use the below command for that
```bash
go mod init sample.com/hello
```
It will create `go.mod` file as below.
```bash
module sample.com/hello
go 1.14
```
Now create file `hello.go` with below contents.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

There are three different commands to run the above GO program. First cd to the ‘hello’ directory which contains the `hello.go` file.

### go install
It will compile the program and place the binary at the `$GOBIN` folder. Type the `go install` command, make sure you are in the hello directory which contains the hello.go file.

It will create a binary named `hello` in the `$GOBIN` directory. 

> The name of the binary is same as last part of the import path of the module. The import path of the module is sample.com/hello and last part of import path is simply hello. Hence binary name will be hello in our case. 

Type hello on the terminal, it will give below output:
```bash
Hello World
```
Output of `which hello` will be `~/Desktop/go/bin/hello `. Remember our `$GOBIN` path was `~/Desktop/go/bin`. Hence the binary `hello` was copied to this directory.

### go build
It will compile the program and place the binary in the current working directory. 
Type the command `go build`. It will create the binary named `hello` same as impbeside hello.go file. To run the binary type `./hello` on the terminal. It will give below output
```bash
Hello World
```
### go run
The `go run hello.go` command will compile and then execute the binary. Type the command `go run hello.go`.
It will output:
```bash
Hello World
```