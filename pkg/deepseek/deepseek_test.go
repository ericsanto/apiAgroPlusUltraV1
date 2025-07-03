package deepseek

import (
	"testing"
	"time"

	"github.com/cohesion-org/deepseek-go"
	"github.com/stretchr/testify/assert"
)

func TestCreatePromptModel_Success(t *testing.T) {

	deepseek := &DeepSeek{}

	promptModel := PromptModel{
		CurrentTemperature:         20,
		RelativeAirHumidity:        10,
		WindSpeed:                  8,
		WindDirection:              190,
		SolarRadiation:             6,
		SoilHumidity:               80,
		SystemIrrigationEfficiency: 80,
		SpaceBetweenPlants:         0.50,
		SpaceBetweenRows:           0.20,
		AtmosphericPressure:        30,
		StartDatePlanting:          time.Now().String(),
		SoilType:                   "argiloso",
		IrrigationType:             "aspersao",
		AgricultureCulture:         "milho",
		BatchName:                  "lote 5",
	}

	result := deepseek.CreatePrompt(promptModel)

	assert.NotEmpty(t, result)
	assert.Contains(t, result, "milho")
	assert.Contains(t, result, "aspersao")
	assert.Contains(t, result, "argiloso")
	assert.Contains(t, result, "lote 5")
	assert.Contains(t, result, "Devo irrigar")

}

func TestCreateRequest_Success(t *testing.T) {

	mock := new(DeepSeekMock)

	content := "teste"

	requestExpected := &deepseek.ChatCompletionRequest{
		Model: deepseek.DeepSeekChat,
		Messages: []deepseek.ChatCompletionMessage{
			{Role: deepseek.ChatMessageRoleUser, Content: content},
		},
	}

	mock.On("CreateRequest", content).Return(requestExpected)

	request := mock.CreateRequest(content)

	assert.Equal(t, requestExpected, request)
	assert.EqualValues(t, requestExpected.Model, request.Model)
	assert.EqualValues(t, requestExpected.Messages[0].Role, request.Messages[0].Role)
	assert.EqualValues(t, requestExpected.Messages[0].Content, request.Messages[0].Content)

	mock.AssertExpectations(t)
}
