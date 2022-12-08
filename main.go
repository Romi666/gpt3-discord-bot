package main

import (
	"context"
	"fmt"
	"github.com/Romi666/gpt3-discord-bot/completion"
	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env files : %v", err)
		return
	}
}
func main() {
	c := gogpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := completion.NewCompletion("apa itu bilangan prima?")
	resp, err := req.SendCompletionsRequest(ctx, c)
	if err != nil {
		log.Fatalf("%+v", err)
	}
	fmt.Println(resp.Choices[0].Text)
}
