package api

import (
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
)

func GenerateKemonoPrompt(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoPromptPrompt(kemono.Concepts.Concepts())
	if err != nil {
		return nil, err
	}
	return GenerateTextByGPT4(prompt)
}

func GenerateKemonoDescription(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoDescriptionPrompt(kemono.Concepts.Concepts(), ImageToBase64(kemono.Image))
	if err != nil {
		return nil, err
	}
	return GenerateTextByGPT4(prompt)
}
