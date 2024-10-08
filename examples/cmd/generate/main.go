package main

import (
	"context"
	"fmt"
	"os"

	coherego "github.com/wesleyfantinel/cohere-go"
	"github.com/wesleyfantinel/cohere-go/request"
	"github.com/wesleyfantinel/cohere-go/response"
)

func main() {

	client := coherego.New(os.Getenv("COHERE_API_KEY"))

	resp := &response.Generate{}
	err := client.Generate(
		context.Background(),
		&request.Generate{
			Prompt: "Please explain to me how LLMs work",
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
