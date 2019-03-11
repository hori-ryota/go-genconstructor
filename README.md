# go-genconstructor

Constructor generator for Go.

## Usage

```go
    //genconstructor [-p]
    type Foo struct {
        key string `required:"[constValue]"`
    }
```

with `go generate` command

```go
    //go:generate go-genconstructor
```

### Example

def
[./_example/person.go](./_example/person.go)

generated
[./_example/example_constructor_gen.go](./_example/example_constructor_gen.go)


## Installation

```sh
$ go get github.com/hori-ryota/go-genconstructor
```
