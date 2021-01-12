package client

import "fmt"

// ResponseError is the representation of a response error.
type ResponseError struct {
	Code    int
	Message string
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Service responden with Code: %d, Message: %s", e.Code, e.Message)
}
