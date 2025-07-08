package myerror

import (
	"errors"
	"fmt"
)

var ErrDuplicateSale = errors.New("já existe uma venda para este planting_id")
var ErrNotFoundSale = errors.New("não existe venda com id")
var ErrViolatedForeingKey = errors.New("nao existe")
var ErrNotFound = errors.New("não existe")
var ErrDuplicateKey = errors.New("já existe")
var ErrEnumInvalid = errors.New("está invalido")
var ErrStatusServiceUnavailable = errors.New("servico temporariamente fora do ar. Tente novamente mais tarde")
var ErrTimeOut = errors.New("tempo excedido")
var ErrImageSizeToLarge = errors.New("tamanho da imagem maior que 20MB")
var ErrUnsupportedImageType = errors.New("tipo da imagem nao e png, jpg ou jpeg")
var ErrNewCurrent = errors.New("erro ao criar cliente para buscar dados atuais do tempo")
var ErrCreateClientMosquitto = errors.New("erro ao criar cliente mosquitto")
var ErrSearchCurrentByCoordinatesOpenWeather = errors.New("erro ao buscar dados climaticos a partir das corrdenadas fornecidas")
var ErrGetUVSolarRadiationOpenWeather = errors.New("erro ao buscar radiacao solar")
var ErrConectDeepSeek = errors.New("erro ao tentar se conectar com api deepseek")
var ErrPublishMosquitto = errors.New("erro ao publicar mensagem no broker")
var ErrFarmNotFound = errors.New("nao existe fazenda com o id")
var ErrBatchAlreadyExists = errors.New("ja existe lote cadastrado com esse nome")

type ErrorApp struct {
	Message   interface{} `json:"message"`
	Code      int         `json:"code"`
	Timestamp string      `json:"timestamp"`
}

func NewError(message, timestamp string, code int) *ErrorApp {
	return &ErrorApp{Message: message, Code: code, Timestamp: timestamp}
}

func (e *ErrorApp) Error() string {
	return fmt.Sprintf("%d, %s, %s", e.Code, e.Message, e.Timestamp)
}

func InterpolationErrViolatedForeingKey(message string, id uint) string {
	return fmt.Sprintf("%s %d ", message, id)
}
