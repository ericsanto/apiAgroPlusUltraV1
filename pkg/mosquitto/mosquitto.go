package mosquitto

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// mqtt "github.com/eclipse/paho.mqtt.golang"
	mqtt "github.com/eclipse/paho.mqtt.golang"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/pkg/jsonutil"
	mystrings "github.com/ericsanto/apiAgroPlusUltraV1/pkg/my_strings"
)

var (
	URL_BROKER_MOSQUITTO      = os.Getenv("URL_BROKER_MOSQUITTO")
	PASSWORD_BROKER_MOSQUITTO = os.Getenv("PASSWORD_BROKER_MOSQUITTO")
	USERNAME_BROKER_MOSQUITTO = os.Getenv("USERNAME_BROKER_MOSQUITTO")
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

func FormatResponseDeepSeekInJSONForMqttBrokerPublisher(content, delimiterStart, delimiterEnd string) (MessagePublicMQTT, error) {

	messageForMqtt, err := mystrings.TakePartOfAText(content, delimiterStart, delimiterEnd)

	fmt.Println(messageForMqtt)

	var json MessagePublicMQTT

	if err != nil {
		return json, fmt.Errorf("%w", err)
	}

	err = jsonutil.ConvertStringToJson(messageForMqtt, &json)

	if err != nil {
		return json, err
	}

	return json, nil
}

func CreateClient() (mqtt.Client, error) {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(URL_BROKER_MOSQUITTO)
	opts.Username = USERNAME_BROKER_MOSQUITTO
	opts.Password = PASSWORD_BROKER_MOSQUITTO
	opts.CleanSession = true
	opts.ClientID = "clientId-teste1"

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Println("erro ao criar cliente", token.Error())
		return nil, fmt.Errorf("%w: %v", myerror.ErrCreateClientMosquitto, token.Error())
	}

	return client, nil
}

func Publisher(client mqtt.Client, topicName string, message interface{}) (bool, error) {

	payload, err := json.Marshal(message)

	if err != nil {
		return false, fmt.Errorf("erro ao converter message para []byte %w", err)
	}

	publishMessageToken := client.Publish(topicName, 0, false, payload)
	publishMessageToken.Wait()

	if publishMessageToken.Error() != nil {
		return false, fmt.Errorf("%w: %v", myerror.ErrPublishMosquitto, publishMessageToken.Error())
	}

	return true, nil

}
