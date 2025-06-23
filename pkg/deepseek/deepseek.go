package deepseek

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cohesion-org/deepseek-go"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

var (
	DEEPSEEK_API_KEY = os.Getenv("DEEPSEEK_API_KEY")
)

//func ConnectDeepSeek(temperatureMax, relativeAirHumidity, windSpeed, solarRadiation, soilHumidity float64, fenologicalStage, soilType, s)

func CreatePromptDeepSeek(currentTemperature, relativeAirHumidity, windSpeed, windDirection, solarRadiation, soilHumidity,
	systemIrrigationEfficiency, spaceBetweenPlants, spaceBetweenRows, atmosphericPressure float64,
	startDatePlanting, soilType, irrigationType, agricultureCulture, batchName string) string {

	currentDate := time.Now().String()

	content := fmt.Sprintf(
		"Lote: %s\n"+
			"Cultura: %s. Estágio fenológico: baseado na data do inicio da plantacao. Determine de acordo com a cultura. Data da plantacao: %s Data atual: %s\n"+
			"Temperatura: %.2f °C\n"+
			"Umidade relativa do ar: %.2f%%. \n"+
			"Velocidade do vento: %.2f km/h com direção %.2f °.\n"+
			"Radiação líquida: %.2f MJ/m²/dia. Pressão atmosférica ao nível local: %.2f kPa.\n"+
			"humidade do solo  em %.2f%%"+
			"Espaçamento entre plantas: %.2f m. Espaçamento entre linhas: %.2f m.\n"+
			"Eficiência do sistema de irrigação: %.2f%%.\n"+
			"Tipo de irrigacao: %s\n"+
			"Pergunta: Devo irrigar a plantação nas próximas 24h? Se sim, indique aproximadamente quantos milímetros de água são necessários para irrigar adequadamente. "+
			"1. Calcule o estágio fenológico atual considerando a data de plantio e o ciclo da cultura\n"+
			"2. Calcule a evapotranspiração da cultura (ETc) considerando o estágio fenológico\n"+
			"3. Avalie a necessidade de irrigação com base na ETc e umidade do solo\n"+
			"4. Se irrigação for necessária, calcule:\n"+
			"   - Lâmina de irrigação em mm\n"+
			"   - Volume por planta em litros\n\n"+
			"FORMATO DE SAÍDA ESPERADO:\n"+
			"SAIDA FINAL\n"+
			"- lote: [nome]\n"+
			"- estagio_fenologico_atual: [estágio calculado]\n"+
			"- decisao: [true/false]\n"+
			"- motivo: [ex: \"Umidade do solo (X%%) abaixo da necessidade da cultura neste estágio (Y%%)\"]\n"+
			"- ETc: [valor]\n"+
			"- lamina_de_irrigacao_mm: [valor]\n"+
			"- volume_por_planta_litros: [valor]\n"+
			"Eu nao quero que vc escreva nada abaixo do texto da saida final. Se tiver observacoes, escreva antes da saida final. Alias, quero que essa saida seja formatada em formato json e vc remova o titulo SAIDA FINAL, so me envie o json sem titulo e sem nenhum acento nas palavras. Mas antes desse json, eu quero as explicacoes das decisoes tomadas\n",
		batchName, agricultureCulture, startDatePlanting, currentDate, currentTemperature, relativeAirHumidity, windSpeed,
		windDirection, solarRadiation, atmosphericPressure, soilHumidity, spaceBetweenPlants, spaceBetweenRows,
		systemIrrigationEfficiency, irrigationType)

	return content

}

func ConnectDeepSeek(secretKey string) *deepseek.Client {

	client := deepseek.NewClient(secretKey)

	return client

}

func CreateRequestDeepSeek(content string) *deepseek.ChatCompletionRequest {

	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleUser, Content: content},
		},
	}

	return request

}

func SendTheRequestAndHandleTheResponse(client *deepseek.Client, request *deepseek.ChatCompletionRequest) (*deepseek.ChatCompletionResponse, error) {

	ctx := context.Background()
	response, err := client.CreateChatCompletion(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", myerror.ErrConectDeepSeek, err)
	}

	return response, nil

}

func RequestRecommendationIrrigationDeepSeek(currentTemperature, relativeAirHumidity, windSpeed, windDirection, solarRadiation, soilHumidity,
	systemIrrigationEfficiency, spaceBetweenPlants, spaceBetweenRows, atmosphericPressure float64,
	startDatePlanting, soilType, irrigationType, agricultureCulture, batchName string) (string, error) {

	content := CreatePromptDeepSeek(currentTemperature, relativeAirHumidity, windSpeed, windDirection, solarRadiation, soilHumidity,
		systemIrrigationEfficiency, spaceBetweenPlants, spaceBetweenRows, atmosphericPressure, startDatePlanting,
		soilType, irrigationType, agricultureCulture, batchName)

	client := ConnectDeepSeek(DEEPSEEK_API_KEY)

	request := CreateRequestDeepSeek(content)

	response, err := SendTheRequestAndHandleTheResponse(client, request)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	// Print the response
	// fmt.Println("Response:", response.Choices[0].Message.Content)

	return response.Choices[0].Message.Content, nil
}
