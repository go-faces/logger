package logger

import (
	"io"
	"log"
	"testing"
)

type testLogger struct {
}

func TestBuildtimeCheck(t *testing.T) {
	var _ Interface = &log.Logger{}
}
func TestRuntimeAssertion(t *testing.T) {
	var a Interface = &testInterface{}
	if _, ok := a.(*testInterface); !ok {
		t.Fatalf("failed to assert built-in type log.Logger")
	}
}

type testInterface struct{}

func (i *testInterface) SetOutput(w io.Writer)                  {}
func (i *testInterface) Output(calldepth int, s string) error   { return nil }
func (i *testInterface) Printf(format string, v ...interface{}) {}
func (i *testInterface) Print(v ...interface{})                 {}
func (i *testInterface) Println(v ...interface{})               {}
func (i *testInterface) Fatal(v ...interface{})                 {}
func (i *testInterface) Fatalf(format string, v ...interface{}) {}
func (i *testInterface) Fatalln(v ...interface{})               {}
func (i *testInterface) Panic(v ...interface{})                 {}
func (i *testInterface) Panicf(format string, v ...interface{}) {}
func (i *testInterface) Panicln(v ...interface{})               {}
func (i *testInterface) Flags() int                             { return 0 }
func (i *testInterface) SetFlags(flag int)                      {}
func (i *testInterface) Prefix() string                         { return "" }
func (i *testInterface) SetPrefix(prefix string)                {}
