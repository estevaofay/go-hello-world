# Compile it

`go install github.com/xFayre/go-hello-world`

or

`go install .`

## FYI (excerpt from the official docs)

The install directory is controlled by the GOPATH and GOBIN environment variables. If GOBIN is set, binaries are installed to that directory. If GOPATH is set, binaries are installed to the bin subdirectory of the first directory in the GOPATH list. Otherwise, binaries are installed to the bin subdirectory of the default GOPATH ($HOME/go or %USERPROFILE%\go).

---

`$GOPATH/bin` or inside your `GOBIN` directory

On my local machine that is inside *~/golib/bin/go-hello-world* and that is an executable

# Run it

`go run .`

or

`~/golib/go-hello-world`