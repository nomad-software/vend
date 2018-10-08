# Vend

**A small command line utility for fully vendoring module dependencies**

---

## Why?

[Because Google have a different idea of what vendoring means than the rest of us.](https://github.com/golang/go/issues/26366) If you use the built-in `go mod vendor` command, it will only cherry pick certain files for inclusion in the vendor folder. This can cause problems when using [Cgo](https://blog.golang.org/c-go-cgo) because it [ignores C files that are not in the package directory](https://github.com/golang/go/issues/26366#issuecomment-405683150). Tests and examples for dependencies are ignored too.

This tool copies the entire dependency tree into the vendor folder like every other package manager does and how every sane developer would expect it to work. It can be used safely in the `$GOPATH` or elsewhere.

This package expects that the new [module system](https://github.com/golang/go/wiki/Modules) is being used, if not it will attempt to create and update a module.

## Install

```
$ go get github.com/nomad-software/vend
```

## Usage

```
$ cd $GOPATH/mypackage
$ vend
```

## Help

Run the following command for help.

```
$ vend -help
```
