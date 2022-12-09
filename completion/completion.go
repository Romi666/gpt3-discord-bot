package completion

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type Completion struct {
	Client *gogpt.Client
}

func NewCompletion(client *gogpt.Client) Completion {
	return Completion{
		Client: client,
	}
}

func (c *Completion) SendCompletionsRequest(ctx context.Context, prompt string) (response gogpt.CompletionResponse, err error) {
	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 2048,
		Prompt:    prompt,
	}
	response, err = c.Client.CreateCompletion(ctx, req)
	if err != nil {
		return gogpt.CompletionResponse{}, err
	}

	return response, nil
}
