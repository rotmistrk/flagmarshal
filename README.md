# flagmarshal
## SYNOPSIS

A simple golang marshaller from commandline to a struct

```go
ParseFlags(structptr interface{}) error
```

## DESCRIPTION

Very simple implementation to marshal commandline arguments to a struct.

ParseFlags uses flag package at the back.

For a struct passed to the ParseXXX function, for any field that has non-empty "flag" tag, 
it adds one or more flags to parse.  

- If the field is not exportable, an error will be returned and no parsing will be performed

- If the field type is not supported, an error will be returned and no parsing will be performed

## TAGS

flag:
    comma-separated list of synonymous flags to add to parser

help:
    help text to use for the flags

## EXAMPLE
```go
type MyArgs struct {
    StrVal string `flag:"s,str" help:"string value"`
    IntVal int    `flag:"i,int" help:"integer value"`
    ...
}

func main() {
    args := MyArgs{intVal: 5}  // default values
    err := flagmarshal.ParseFlags(&args)
    ...
}
```

## TYPES

Types supported are limited to the backend: flag package.
Since flag package does not callback on each argument, but parses first to a map,
lists are not supported.

- string
- uint64
- int64
- uint
- int
- float64
- float
- bool



```
