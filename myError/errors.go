package myerror

import (
	"fmt"
)

type ErrorApp struct {
	Message   interface{} `json:"message"`
	Code      uint        `json:"code"`
	Timestamp string      `json:"timestamp"`
}

func NewError(message, timestamp string, code uint) *ErrorApp {
	return &ErrorApp{Message: message, Code: code, Timestamp: timestamp}
}

func (e *ErrorApp) Error() string {
	return fmt.Sprintf("%d, %s, %s", e.Code, e.Message, e.Timestamp)
}
