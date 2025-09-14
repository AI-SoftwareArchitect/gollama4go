package main

import (
	"fmt"
	"gollama4go/entities"
	"gollama4go/prompt_builder"
)

func main() {
	g4g := entities.NewG4GClient(
		"http://127.0.0.1:11434",
		20,
	)
	models, err := g4g.ListOllamaModels()
	if err != nil {
		return
	}
	fmt.Println(models)

	pBuilder := prompt_builder.PromptBuilder{
		prompt_builder.PromptType(2),
	}

}
