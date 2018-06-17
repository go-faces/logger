package logger

import "io"

// Interface provides an interface with a method set compatible with
// the Go log.Logger type.  It can be used to implement loggers using the
// with similar method signatures.
type Interface interface {
	// SetOutput sets the output destination for the logger.
	SetOutput(w io.Writer)

	// Output writes the output for a logging event. The string s contains
	// the text to print after the prefix specified by the flags of the
	// Logger. A newline is appended if the last character of s is not
	// already a newline. Calldepth is used to recover the PC and is
	// provided for generality, although at the moment on all pre-defined
	// paths it will be 2.
	Output(calldepth int, s string) error
	// Printf calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(format string, v ...interface{})
	// Print calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Print(v ...interface{})
	// Println calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Println(v ...interface{})
	// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
	Fatal(v ...interface{})
	// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
	Fatalf(format string, v ...interface{})
	// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
	Fatalln(v ...interface{})
	// Panic is equivalent to l.Print() followed by a call to panic().
	Panic(v ...interface{})
	// Panicf is equivalent to l.Printf() followed by a call to panic().
	Panicf(format string, v ...interface{})
	// Panicln is equivalent to l.Println() followed by a call to panic().
	Panicln(v ...interface{})
	// Flags returns the output flags for the logger.
	Flags() int
	// SetFlags sets the output flags for the logger.
	SetFlags(flag int)
	// Prefix returns the output prefix for the logger.
	Prefix() string
	// SetPrefix sets the output prefix for the logger.
	SetPrefix(prefix string)
}
