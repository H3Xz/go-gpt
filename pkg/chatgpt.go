package pkg

import (
	"context"
	"fmt"
	"os"

	"github.com/franciscoescher/goopenai"
	"github.com/joho/godotenv"
)

func test() {
	err := godotenv.Load()
	if err != nil {
		Log.Fatal(err)
	}

	apiKey := os.Getenv("APIKEY")
	organization := os.Getenv("ORGANIZATION")

	client := goopenai.NewClient(apiKey, organization)

	r := goopenai.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []goopenai.Message{
			{
				Role:    "user",
				Content: "Say this is a test!",
			},
		},
		Temperature: 0.7,
	}

	completions, err := client.CreateCompletions(context.Background(), r)
	if err != nil {
		panic(err)
	}

	fmt.Println(completions)
}
