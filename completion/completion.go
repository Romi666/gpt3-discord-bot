package completion

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type Completion struct {
	Prompt string
}

func NewCompletion(prompt string) Completion {
	return Completion{
		Prompt: prompt,
	}
}


func (c *Completion) SendCompletionsRequest(ctx context.Context, client *gogpt.Client) (response gogpt.CompletionResponse, err error) {
	req := gogpt.CompletionRequest{
		Model: "text-davinci-003",
		MaxTokens: 2048,
		Prompt:    c.Prompt,
	}
	response, err = client.CreateCompletion(ctx, req)
	if err != nil {
		return gogpt.CompletionResponse{}, err
	}

	return response, nil
}
