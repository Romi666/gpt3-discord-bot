package connection

import (
	gogpt "github.com/sashabaranov/go-gpt3"
	"os"
)

func InitConnectionGPT() *gogpt.Client {
	var API_KEY = os.Getenv("OPENAI_API_KEY")
	client := gogpt.NewClient(API_KEY)

	return client
}
