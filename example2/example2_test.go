package example2

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

func TestOK(t *testing.T) {
	go sendMessage(make(chan string, 1), "hello")
}

func TestNG(t *testing.T) {
	go sendMessage(make(chan string, 0), "hello")

	// After the test passes, TestMain reports the following errors
	//
	// goleak: Errors on successful test run: found unexpected goroutines:
	//[Goroutine 7 in state chan send, with github.com/kei2100/goleak-example/example2.sendMessage on top of the stack:
}
