package myerror

import (
	"errors"
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

var ErrDuplicateSale = errors.New("já existe uma venda para este planting_id")
var ErrNotFoundSale = errors.New("não existe venda com id")
var ErrViolatedForeingKey = errors.New("nao existe")
var ErrNotFound = errors.New("não existe")
var ErrDuplicateKey = errors.New("já existe")
var ErrEnumInvalid = errors.New("está invalido")

func InterpolationErrViolatedForeingKey(message string, id uint) string {
	return fmt.Sprintf("%s %d ", message, id)
}
