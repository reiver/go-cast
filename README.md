# go-cast

Package cast provides tools for safely converting from one type to another, without information loss,
for the Go programming language.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-cast

[![GoDoc](https://godoc.org/github.com/reiver/go-cast?status.svg)](https://godoc.org/github.com/reiver/go-cast)


## Example

For an example, let's look at converting to an `int64`.

You can safely convert a `uint8`, `uint16`, `uint32`, plus an `int`, `int8`, `int16`, `int32`, `int64`,
into an `int64` without information loss.

In other words, if `v` (in the following example code) is any of the types in the list just mentioned, then the following
code would successfully return an `int64` of the same numeric value.

```go
i64, err := cast.Int64(v)
```

## More Examples

There is something similar for the other Go built-in types. I.e.,:

```go
b, err := cast.Bool(v)
```

```go
c64, err := cast.Complex64(v)
```
```go
c128, err := cast.Complex128(v)
```

```go
f32, err := cast.Float32(v)
```
```go
f64, err := cast.Float64(v)
```

```go
i, err := cast.Int(v)
```
```go
i8, err := cast.Int8(v)
```
```go
i16, err := cast.Int16(v)
```
```go
i32, err := cast.Int32(v)
```
```go
i64, err := cast.Int64(v)
```

```go
s, err := cast.String(v)
```

```go
t, err := cast.Time(v)
```

```go
u, err := cast.Uint(v)
```
```go
u8, err := cast.Uint8(v)
```
```go
u16, err := cast.Uint16(v)
```
```go
u32, err := cast.Uint32(v)
```
```go
u64, err := cast.Uint64(v)
```

## Advanced Example

In addition to these types built into Go (that were in the list), this package also supports converting
custom types, as long as they implement a special method.

For int64, the special method is:

```go
interface {
	Int64() (int64, error)
}
```

For example:
```go
type Always5 struct{}

func (receiver Always5) Int64() (int64, error) {
	return 5, nil
}

// ...

var v Always5

// ...

i64, err := cast.Int64(v)
```
