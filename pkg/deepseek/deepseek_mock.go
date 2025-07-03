package deepseek

import (
	"context"

	"github.com/cohesion-org/deepseek-go"
	"github.com/stretchr/testify/mock"
)

type DeepSeekMock struct {
	mock.Mock
}

func (dsm *DeepSeekMock) CreateRequest(content string) *deepseek.ChatCompletionRequest {

	args := dsm.Called(content)

	return args.Get(0).(*deepseek.ChatCompletionRequest)
}

func (dsm *DeepSeekMock) SendTheRequestAndHandleTheResponse(request *deepseek.ChatCompletionRequest) (*deepseek.ChatCompletionResponse, error) {

	args := dsm.Called(request)

	return args.Get(0).(*deepseek.ChatCompletionResponse), args.Error(1)
}

func (dsm *DeepSeekMock) CreateChatCompletion(ctx context.Context, request *deepseek.ChatCompletionRequest) (*deepseek.ChatCompletionResponse, error) {

	args := dsm.Called(ctx, request)

	return args.Get(0).(*deepseek.ChatCompletionResponse), args.Error(1)

}
