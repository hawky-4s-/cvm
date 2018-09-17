package internal

import "strings"

type ExecutionError struct {
	executionErrors []error
}

func (e *ExecutionError) Error() string {
	if e.executionErrors != nil {
		errStrings := make([]string, len(e.executionErrors))

		for _, err := range e.executionErrors {
			errStrings = append(errStrings, err.Error())
		}
		return strings.Join(errStrings[:], "\n")
	}

	return ""
}

func NewExecutionError(errors ...error) *ExecutionError {
	return &ExecutionError{executionErrors: errors}
}
