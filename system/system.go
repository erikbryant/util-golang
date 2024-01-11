package system

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
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

// dumpDebug prints a debugging message to stderr
func dumpDebug() {
	nGo := runtime.NumGoroutine()
	buf := make([]byte, nGo*1024)
	runtime.Stack(buf, true)
	lines := strings.Split(string(buf), "\n")
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "goroutine ") {
			fmt.Fprintf(os.Stderr, "%16s", lines[i][9:len(lines[i])-1])
			if i+1 < len(lines) {
				fmt.Fprintf(os.Stderr, " %25s", lines[i+1])
				i++
			}
			if i+1 < len(lines) {
				fmt.Fprintf(os.Stderr, "  %s", lines[i+1])
				i++
			}
			fmt.Fprintf(os.Stderr, "\n")
		}
	}
}

// InstallDebug installs a signal handler that prints a debugging message to stderr when SIGUSR1 is received
func InstallDebug() int {
	// To signal a debugging dump:
	//   $ kill -SIGUSR1 <pid>

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)

	go func() {
		for {
			<-c
			dumpDebug()
		}
	}()

	return os.Getpid()
}
