package api

import (
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
)

func GenerateKemonoPrompt(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoPromptPrompt(kemono.Concepts.Concepts())
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4(prompt)
}

func GenerateKemonoImage(kemono *domains.Kemono) ([]byte, error) {
	return generateKemonoImage(kemono.Prompt)
}

func GenerateKemonoDescription(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoDescriptionPrompt(kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4(prompt)
}

func GenerateKemonoStatus(kemono *domains.Kemono) (*domains.KemonoStatus, error) {
	prompt, err := createKemonoStatusPrompt(kemono.Description, kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	kemonoStatusText, err := generateTextByGPT4(prompt)
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
	kemonoCharacterChipText, err := generateTextByGPT4(prompt)
	if err != nil {
		return nil, err
	}

	return domains.ParseKemonoCharacterChip(kemonoCharacterChipText)
}

func GenerateKemonoName(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoNamePrompt(kemono.Description, kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4(prompt)
}

func BreedKemonoPrompt(kemono1 *domains.Kemono, kemono2 *domains.Kemono) (*string, error) {
	prompt, err := generateBreedKemonoPromptPrompt(kemono1, kemono2)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4(prompt)
}

func BreedKemonoDescription(kemono1 *domains.Kemono, kemono2 *domains.Kemono, kemono3 *domains.Kemono) (*string, error) {
	prompt, err := generateBreedKemonoDescriptionPrompt(kemono1, kemono2, kemono3)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4(prompt)
}

func BreedKemonoStatus(kemono1 *domains.Kemono, kemono2 *domains.Kemono, kemono3 *domains.Kemono) (*domains.KemonoStatus, error) {
	prompt, err := generateBreedKemonoStatusPrompt(kemono1, kemono2, kemono3)
	if err != nil {
		return nil, err
	}
	kemonoStatusText, err := generateTextByGPT4(prompt)
	if err != nil {
		return nil, err
	}
	return domains.ParseKemonoStatus(kemonoStatusText)
}

func BreedKemonoConcepts(kemono1 *domains.Kemono, kemono2 *domains.Kemono, kemono3 *domains.Kemono) (*domains.Concepts, error) {
	prompt, err := generateBreedKemonoConceptsPrompt(kemono1, kemono2, kemono3)
	if err != nil {
		return nil, err
	}
	kemonoConceptsText, err := generateTextByGPT4(prompt)
	if err != nil {
		return nil, err
	}
	return domains.ParseKemonoConcepts(kemonoConceptsText)
}
