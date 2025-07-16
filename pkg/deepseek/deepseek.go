package deepseek

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cohesion-org/deepseek-go"

	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type PromptModel struct {
	CurrentTemperature         float64
	RelativeAirHumidity        float64
	WindSpeed                  float64
	WindDirection              float64
	SolarRadiation             float64
	SoilHumidity               float64
	SystemIrrigationEfficiency float64
	SpaceBetweenPlants         float64
	SpaceBetweenRows           float64
	AtmosphericPressure        float64
	StartDatePlanting          string
	SoilType                   string
	IrrigationType             string
	AgricultureCulture         string
	BatchName                  string
}

type LLMMethods interface {
	RequestRecommendationIrrigation(ctx context.Context, promptModel PromptModel) (string, error)
}

type DeepSeekInterface interface {
	CreatePrompt(promptModel PromptModel) string
	CreateRequest(content string) *deepseek.ChatCompletionRequest
	SendTheRequestAndHandleTheResponse(ctx context.Context, request *deepseek.ChatCompletionRequest) (*deepseek.ChatCompletionResponse, error)
}

type DeepSeek struct {
	apiKey         string
	deepseekClient *deepseek.Client
}

func NewDeepSeek(apiKey string) LLMMethods {
	return &DeepSeek{apiKey: apiKey,
		deepseekClient: deepseek.NewClient(apiKey)}
}

//func ConnectDeepSeek(temperatureMax, relativeAirHumidity, windSpeed, solarRadiation, soilHumidity float64, fenologicalStage, soilType, s)

func (ds *DeepSeek) CreatePrompt(promptModel PromptModel) string {

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
			"Tipo de solo: %s\n"+
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
		promptModel.BatchName, promptModel.AgricultureCulture, promptModel.StartDatePlanting,
		currentDate, promptModel.CurrentTemperature, promptModel.RelativeAirHumidity,
		promptModel.WindSpeed, promptModel.WindDirection, promptModel.SolarRadiation,
		promptModel.AtmosphericPressure, promptModel.SoilHumidity, promptModel.SpaceBetweenPlants,
		promptModel.SpaceBetweenRows, promptModel.SystemIrrigationEfficiency, promptModel.IrrigationType,
		promptModel.SoilType)

	return content

}

func (ds *DeepSeek) CreateRequest(content string) *deepseek.ChatCompletionRequest {

	request := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleUser, Content: content},
		},
	}

	return request

}

func (ds *DeepSeek) SendTheRequestAndHandleTheResponse(ctx context.Context, request *deepseek.ChatCompletionRequest) (*deepseek.ChatCompletionResponse, error) {

	response, err := ds.deepseekClient.CreateChatCompletion(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", myerror.ErrConectDeepSeek, err)
	}

	return response, nil

}

func (ds *DeepSeek) RequestRecommendationIrrigation(ctx context.Context, promptModel PromptModel) (string, error) {

	content := ds.CreatePrompt(promptModel)

	request := ds.CreateRequest(content)

	response, err := ds.SendTheRequestAndHandleTheResponse(ctx, request)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	// Print the response
	// fmt.Println("Response:", response.Choices[0].Message.Content)

	return response.Choices[0].Message.Content, nil
}
