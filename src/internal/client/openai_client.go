package client

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"llm-client/models"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient(token string) *OpenAIClient {
	return &OpenAIClient{
		client: openai.NewClient(token),
	}
}

func (o *OpenAIClient) Ask(ctx context.Context, request models.Request) (*models.Response, error) {
	resp, err := o.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: request.Model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: request.Question,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	return &models.Response{
		Answer: resp.Choices[0].Message.Content,
	}, nil
}
