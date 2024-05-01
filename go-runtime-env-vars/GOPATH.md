# GOPATH

The discussion for `GOPATH` will revolve around whether you are using GO version `1.13` higher or lower. In version `1.13`, GO introduced a new feature for dependency management called GO modules. Let’s first understand the legacy behavior of `GOPATH` and then we will discuss what has changed with respect to `GOPATH` after GO version `1.13`.

## Before GO version 1.13

The `GOPATH` env variable was used for resolving go imports statements as well as it also specifies the location of your GO workspace. It is the root of GO workspace. A dependency cannot be imported if it is not inside `GOPATH`. Hence earlier version required all your source code to be inside `GOPATH`. So basically `GOPATH` is used for

1. Resolving imports. A dependency cannot be imported if it is not inside GOPATH. Hence it still requires your code to be inside `GOPATH`
2. Packages are installed in the directory `$GOPATH/pkg/$GOOS_$GOARCH`.
3. Store compiled application binary in `$GOPATH/bin` (This can be overridden by setting $GOBIN to a different path)
It contains the following folders

### src –
Source file location. It contains `.go` and other source files
When installing any dependency package using `go get`, all the package files are stored in this location.

### pkg –
1. Stores the compilation output of your actual source code present in `src`. directory. It contains `.a` files
1. Basically, it contains the GO packages compiled from the `src/` directory. They are then used at link time to create binary executables which are then placed in the `bin/` directory.
2. It is a good idea to compile a package once and used it to create different binary executables.
3. Each pair of Operating System and Architecture will have its own subdirectory in `pkg`. (eg: `pkg/GOOS_GOARCH`)
### bin – 
Location of executables built by GO. It contains binary executables of your GO programs.

## With Go version 1.13 or later

In version `1.13`, GO announced a new feature of dependency management called GO modules. We will learn about this feature in upcoming tutorials. For now, you can imagine that with the new feature, GO doesn’t require putting all GO code in the Go workspace or in the `$GOAPTH/src` directory. You can create the directory anywhere and put your Go program there. Note that all legacy behavior of GOPATH is still valid for version `1.13` and higher.  Also note that with this new GO module feature, GO programs can be run in two ways. 

Using modules: when using modules, GOPATH is not used for resolving imports. However, when using modules for running GO programs, `GOPATH` will be used for
1. Store download source code in `$GOPATH/pkg/mod` directory.
2. Store compiled application binary in `$GOPATH/bin` (This can be overridden by setting `$GOBIN` to a different path)
3. Not using modules:  It is still possible to run a GO program in the legacy ways even with version `1.13` or later. When not using modules while running GO programs, `$GOPATH` behavior is same as earlier versions which is the same as mentioned above.
   
## Set Up GOPATH

If this env variable is unset, it defaults `$HOME/go` on Unix systems and `%USERPROFILE%\go` on Windows. If your workspace location is `~/Desktop/go`, then make below entry to the `~/$HOME/.profile` file. It is good idea to always have the `GOPATH` set up as `GOPATH` is used even with the introduction of modules.
```bash
export GOPATH=~/Desktop/go
```