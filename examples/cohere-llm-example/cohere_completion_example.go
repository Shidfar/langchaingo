package main

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/shidfar/langchaingo/llms"
	"gitlab.com/shidfar/langchaingo/llms/cohere"
)

func main() {
	llm, err := cohere.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	input := "The first man to walk on the moon"
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(completion)

	inputToken := llms.CountTokens("", input)
	outputToken := llms.CountTokens("", completion)

	fmt.Printf("%v/%v\n", inputToken, outputToken)
}
