package main

import (
	"context"
	"fmt"
	"os"

	coherego "github.com/wesleyfantinel/cohere-go"
	"github.com/wesleyfantinel/cohere-go/model"
	"github.com/wesleyfantinel/cohere-go/request"
	"github.com/wesleyfantinel/cohere-go/response"
)

func main() {
	client := coherego.New(os.Getenv("COHERE_API_KEY"))

	resp := &response.Chat{}
	err := client.ChatStream(
		context.Background(),
		&request.Chat{
			ChatHistory: []model.ChatMessage{
				{
					Role:    model.ChatMessageRoleUser,
					Message: "Hello, how are you?",
				},
				{
					Role:    model.ChatMessageRoleChatbot,
					Message: "I'm good, how are you?",
				},
			},
			Message: "I'm very happy!",
		},
		resp,
		func(resp *response.Chat) {
			if resp.EventType == model.EventTypeTextGeneration {
				fmt.Printf("%s", resp.Text)
			}
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\nresp: %#v\n", resp)

}
