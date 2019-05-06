package maven

import (
	"fmt"
)

//import "calculator-master/errors"

type ErrorResponse struct {
	code    string
	message string
	cause   string
}

const CUSTOMER_ERRORS_PREFIX = "MVN"

var (
	DependencyFetchError = defineError(1, "Could'nt fetch dependency")
	FileCreateError = defineError(2, "Could'nt create file path")
	FileWriteError = defineError(3, "Could'nt write file")
)

func defineError(index int, msg string) func(error) ErrorResponse {
	return func(err error) ErrorResponse {
		return ErrorResponse{
			code:    fmt.Sprintf("%s-%03d", CUSTOMER_ERRORS_PREFIX, index),
			message: msg,
			cause:   err.Error(),
		}
	}
}
