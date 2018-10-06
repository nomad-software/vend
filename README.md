# Vend

**A small command line utility for fully vendoring module dependencies**

---

## Why?

[Because Google have a different idea of what vendoring means than the rest of us.](https://github.com/golang/go/issues/26366) If you use the built-in `go mod vendor` command, it will only cherry pick certain files for inclusion in the vendor folder.

This tool copies the entire dependency tree into the vendor folder like every other package manager does and how every sane developer would expect it to work.

This package expects that the new [module system](https://github.com/golang/go/wiki/Modules) is being used.

## Example

```
vend
```

## Help

Run the following command for help.

```
vend -help
```
