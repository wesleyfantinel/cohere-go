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

	resp := &response.Embed{}
	inputType := model.EmbedSearchQuery
	err := client.Embed(
		context.Background(),
		&request.Embed{
			Texts: []string{
				"Hello world",
				"Ciao mondo",
			},
			Model:     model.EmbedModelEnglishV30,
			InputType: &inputType,
		},
		resp,
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("resp: %#v\n", resp)

}
