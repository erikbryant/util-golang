package system

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"syscall"
)

// MyPath returns the file joined with the absolute path of the source file calling this
func MyPath(file string) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Cannot get caller information")
	}
	dir := path.Dir(filename)
	return path.Join(dir, file)
}

// CtrlT prints a debugging message when SIGUSR1 is sent to this process.
func CtrlT(str string, val *int, digits []int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)

	fmt.Println("$ kill -SIGUSR1", os.Getpid())

	go func() {
		for {
			<-c
			fmt.Println("^T] ", str, *val, digits)
		}
	}()
}
