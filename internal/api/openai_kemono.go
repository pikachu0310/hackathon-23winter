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
	prompt, err := createKemonoDescriptionPrompt(kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	return GenerateTextByGPT4(prompt)
}

func GenerateKemonoStatus(kemono *domains.Kemono) (*domains.KemonoStatus, error) {
	prompt, err := createKemonoStatusPrompt(kemono.Description, kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	kemonoStatusText, err := GenerateTextByGPT4(prompt)
	if err != nil {
		return nil, err
	}
	return domains.ParseKemonoStatus(kemonoStatusText)
}

func GenerateKemonoCharacterChip(kemono *domains.Kemono) (*domains.KemonoCharacterChip, error) {
	prompt, err := createKemonoCharacterChipPrompt(kemono.Description, kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	kemonoCharacterChipText, err := GenerateTextByGPT4(prompt)
	if err != nil {
		return nil, err
	}

	return domains.ParseKemonoCharacterChip(kemonoCharacterChipText)
}
