# logger.Interface (alpha)

An interface compatible with built-in type `log.Logger`.  Use this
interface instead of a concrete type like `log.Logger` when you need logging. This
will allow you to swap implementatations when needed.

```go
type Interface interface {
	SetOutput(w io.Writer)
	Output(calldepth int, s string) error
	Printf(format string, v ...interface{})
	Print(v ...interface{})
	Println(v ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
	Flags() int
	SetFlags(flag int)
	Prefix() string
	SetPrefix(prefix string)
}
```
## Compatible APIs

- Go's Built-in `log.Logger` API
- Logrus Logging API - https://github.com/sirupsen/logrus
- If you know of others, add here.
