package myerror

import "errors"

var ErrKafkaMessageTimeout = errors.New("timeout ao esperar mensagem")
