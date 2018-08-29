package error911

import "github.com/pkg/errors"

// Causer is defined as "causer" in "github.com/pkg/errors" but unexported
type Causer interface {
	Cause() error
}

// An interface used by the type E911 struct
type E911Interface interface {
	Cause() error
	Error() string
	First() error
	Push(msg ...interface{})
	Stacks() (error, string, errors.StackTrace)
}

// StackTracer is defined as "stackTracer" in "github.com/pkg/errors" but unexported
type StackTracer interface {
	StackTrace() errors.StackTrace
}
