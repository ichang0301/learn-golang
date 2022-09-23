package errors

const (
	ErrUnsupportedType = ErrorFormat("unsupported type")
	ErrNoSuchElement   = ErrorFormat("no such element")
	ErrEmptyList       = ErrorFormat("empty slice/array")
)

// ErrorFormat is the format that prints errors occurred
type ErrorFormat string

// Error returns a error context for implementing of error
func (e ErrorFormat) Error() string { // implement 'error' interface. : https://go.dev/blog/error-handling-and-go
	return string(e)
}
