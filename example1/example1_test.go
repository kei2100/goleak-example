package example1

import (
	"testing"

	"go.uber.org/goleak"
)

func sendMessage(ch chan<- string, msg string) {
	ch <- msg
}

func TestOK(t *testing.T) {
	defer goleak.VerifyNone(t)
	go sendMessage(make(chan string, 1), "hello")
}

func TestNG(t *testing.T) {
	defer goleak.VerifyNone(t)
	go sendMessage(make(chan string, 0), "hello") // FAIL: TestFooNG leaks.go:78: found unexpected goroutines:
}
