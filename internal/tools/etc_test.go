package tools

import (
	"os"
	"testing"

	"github.com/bouk/monkey"
)

func assertPanic(t *testing.T, f func(), message string) {
	fakeExit := func(int) {
		panic("os.Exit called")
	}
	patch := monkey.Patch(os.Exit, fakeExit)
	defer patch.Unpatch()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf(message)
		}
	}()
	f()
}
