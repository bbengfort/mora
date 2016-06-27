// Package moralog defines logging for mora
package moralog

import (
	"log"
	"os"
)

// Use golang's standard logger by default
// Access is not thread-safe: do not modify except in init()
var logger Logger = log.New(os.Stdout, "", log.LstdFlags)

// Logger mimics golang's standard Logger as an interface.
type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// SetLogger sets the logger that is used in mora. Call only from init()
func SetLogger(l Logger) {
	logger = l
}

// Fatal is equivalent to Print() followed by a call to os.Exit() with a non-zero exit code.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit() with a non-zero exit code.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit()) with a non-zero exit code.
func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
func Println(args ...interface{}) {
	logger.Println(args...)
}
