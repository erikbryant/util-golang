package system

import (
	"fmt"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"
	"text/tabwriter"
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
	// Format in tab-separated columns with a tab stop of 8.
	w := new(tabwriter.Writer)
	w.Init(os.Stderr, 0, 8, 2, ' ', 0)

	nGo := runtime.NumGoroutine()
	buf := make([]byte, nGo*1024)
	runtime.Stack(buf, true)
	lines := strings.Split(string(buf), "\n")

	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "goroutine ") {
			record := strings.TrimSpace(lines[i][9 : len(lines[i])-1])
			if i+1 < len(lines) {
				i++
				record += "\t" + strings.TrimSpace(lines[i])
			}
			if i+1 < len(lines) {
				i++
				record += "\t" + strings.TrimSpace(lines[i])
			}
			fmt.Fprintln(w, record)
		}
	}
	w.Flush()
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
