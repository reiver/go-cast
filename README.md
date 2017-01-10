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

This package also supports converting from **custom types**, as long as that _custom type_ implement a special method.

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

## More Advanced Example

The other special methods are:

```go
interface {
	Bool() (bool, error)
}
```

```go
interface {
	 Complex64() (complex64, error)
}
```
```go
interface {
	 Complex128() (complex128, error)
}
```

```go
interface {
	 Float32() (float32, error)
}
```
```go
interface {
	 Float64() (float64, error)
}
```

```go
interface {
	 Int() (int, error)
}
```
```go
interface {
	 Int8() (int8, error)
}
```
```go
interface {
	 Int16() (int16, error)
}
```
```go
interface {
	 Int32() (int32, error)
}
```
```go
interface {
	 Int64() (int64, error)
}
```

```go
interface {
	 String() (String, error)
}
```

```go
interface {
	 Time() (time.Time, error)
}
```

```go
interface {
	 Uint() (uint, error)
}
```
```go
interface {
	 Uint8() (uint8, error)
}
```
```go
interface {
	 Uint16() (uint16, error)
}
```
```go
interface {
	 Uint32() (uint32, error)
}
```
```go
interface {
	 Uint64() (uint64, error)
}
```
