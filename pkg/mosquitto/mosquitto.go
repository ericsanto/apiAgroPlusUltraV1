package mosquitto

import (
	"encoding/json"
	"fmt"
	"log"

	// mqtt "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/eclipse/paho.mqtt.golang"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	mystrings "github.com/ericsanto/apiAgroPlusUltraV1/pkg/my_strings"
)

type MessagePublicMQTT struct {
	Batch           string  `json:"lote"`
	PhenologicStage string  `json:"estagio_fenologico_atual"`
	Decision        bool    `json:"decisao"`
	Reason          string  `json:"motivo"`
	ETC             float64 `json:"etc"`
	IrrigationBlade float64 `json:"lamina_de_irrigacao_em_mm"`
	VolumePerPlant  float64 `json:"volume_por_planta_em_litros"`
}

type MosquittoInterface interface {
	FormatResponseDeepSeekInJSONForMqttBrokerPublisher(content, delimiterStart, delimiterEnd string) (MessagePublicMQTT, error)
	Publisher(topicName string, message interface{}) (bool, error)
	Disconnect(quiesce uint)
}

type OPTSMosquitto struct {
	UrlBrokerMosquitto string
	Username           string
	Password           string
	CleanSession       bool
	ClientID           string
}

type Mosquitto struct {
	clientMosquitto mqtt.Client
	jsonUtils       jsonutil.JsonUtilsInterface
}

func NewMosquittoBroker(opts OPTSMosquitto, jsonUtils jsonutil.JsonUtilsInterface) (MosquittoInterface, error) {
	opt := mqtt.NewClientOptions()
	opt.AddBroker(opts.UrlBrokerMosquitto)
	opt.Username = opts.Username
	opt.Password = opts.Password
	opt.CleanSession = opts.CleanSession
	opt.ClientID = opts.ClientID

	client := mqtt.NewClient(opt)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("erro ao criar cliente", token.Error())
		return nil, fmt.Errorf("%w: %v", myerror.ErrCreateClientMosquitto, token.Error())
	}

	return &Mosquitto{clientMosquitto: client, jsonUtils: jsonUtils}, nil

}

func (m *Mosquitto) FormatResponseDeepSeekInJSONForMqttBrokerPublisher(content, delimiterStart, delimiterEnd string) (MessagePublicMQTT, error) {

	messageForMqtt, err := mystrings.TakePartOfAText(content, delimiterStart, delimiterEnd)

	fmt.Println(messageForMqtt)

	var json MessagePublicMQTT

	if err != nil {
		return json, fmt.Errorf("%w", err)
	}

	err = m.jsonUtils.ConvertStringToJson(messageForMqtt, &json)

	if err != nil {
		return json, err
	}

	return json, nil
}

func (m *Mosquitto) Publisher(topicName string, message interface{}) (bool, error) {

	payload, err := json.Marshal(message)

	if err != nil {
		return false, fmt.Errorf("erro ao converter message para []byte %w", err)
	}

	publishMessageToken := m.clientMosquitto.Publish(topicName, 0, false, payload)
	publishMessageToken.Wait()

	if publishMessageToken.Error() != nil {
		return false, fmt.Errorf("%w: %v", myerror.ErrPublishMosquitto, publishMessageToken.Error())
	}

	return true, nil

}

func (m *Mosquitto) Disconnect(quiesce uint) {
	m.clientMosquitto.Disconnect(quiesce)
}
