# Vend

**A small command line utility for fully vendoring module dependencies**

---

## Why?

[Because Google have a different idea of what vendoring means than the rest
of us.](https://github.com/golang/go/issues/26366) If you use the built-in
`go mod vendor` command, it will only cherry pick certain files for inclusion
in the vendor folder. This can cause problems when using
[Cgo](https://blog.golang.org/c-go-cgo) because it [ignores C files that are
not in the package
directory](https://github.com/golang/go/issues/26366#issuecomment-405683150).
Tests and examples for dependencies are ignored too.

This tool copies the entire dependency tree into the vendor folder like every
other package manager does and how every sane developer would expect it to
work. It can be used safely in the `$GOPATH` or elsewhere.

This package expects that the new [module
system](https://github.com/golang/go/wiki/Modules) [introduced in
v1.11](https://golang.org/doc/go1.11) is being used.

## What does it do?

This tool fully copies all files from your project's imported dependencies
into the `vendor` folder. This allows you to:

1. Always have access to _all_ files in your dependencies, even if they go offline
2. Always be able to build your project on a disconnected computer
3. Always be able to run the tests or benchmarks of all your dependencies

## Supported Go versions

* v1.11+

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

## Caveats

* [Go mod replace directives](https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive) are intentionally not supported because there is no consensus how to support them.
