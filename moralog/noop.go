package moralog

import "log"

// NoopLogger implements empty functions for logging instead of writing to
// Discard or another writer, thus improving the performance of stifle logging.
type NoopLogger struct {
	*log.Logger
}

// Fatal is a noop
func (l *NoopLogger) Fatal(args ...interface{}) {}

// Fatalf is a noop
func (l *NoopLogger) Fatalf(format string, args ...interface{}) {}

// Fatalln is a noop
func (l *NoopLogger) Fatalln(args ...interface{}) {}

// Print is a noop
func (l *NoopLogger) Print(args ...interface{}) {}

// Printf is a noop
func (l *NoopLogger) Printf(format string, args ...interface{}) {}

// Println is a noop
func (l *NoopLogger) Println(args ...interface{}) {}
