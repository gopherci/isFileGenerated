# isFileGenerated

`isFileGenerated` checks if a file is generated. A file is generated if

- [github.com/shurcooL/go/analysis#IsFileGenerated](https://godoc.org/github.com/shurcooL/go/analysis#IsFileGenerated) returns true, or
- path contains a directory labelled `testdata`.

Sets the exit status 0 if file is generated, 1 if not generated, or > 1 for other error.

# Install

```bash
go install github.com/gopherci/isFileGenerated
```

# Usage

```bash
isFileGenerated <root directory> <path to file>
```

# Example

```bash
isFileGenerated $GOPATH/src/github.com/hydroflame/fuzzi fuzzi/fuzzy.go; echo $?
github.com/hydroflame/fuzzi/fuzzi/fuzzy.go was not generated
1
```

```bash
isFileGenerated $GOPATH/src/github.com/hydroflame/fuzzi vendor/github.com/go-kit/kit/endpoint/endpoint.go; echo $?
github.com/hydroflame/fuzzi/vendor/github.com/go-kit/kit/endpoint/endpoint.go was generated
0
```

