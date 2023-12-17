package api

import (
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
)

func GenerateKemonoPrompt(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoPromptPrompt(kemono.Concepts.Concepts())
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4Vision(prompt)
}

func GenerateKemonoImage(kemono *domains.Kemono) ([]byte, error) {
	return generateKemonoImage(kemono.Prompt)
}

func GenerateKemonoDescription(kemono *domains.Kemono) (*string, error) {
	prompt, err := createKemonoDescriptionPrompt(kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4Vision(prompt)
}

func GenerateKemonoStatus(kemono *domains.Kemono) (*domains.KemonoStatus, error) {
	prompt, err := createKemonoStatusPrompt(kemono.Description, kemono.Concepts.Concepts(), kemono.Image)
	if err != nil {
		return nil, err
	}
	kemonoStatusText, err := generateTextByGPT4Vision(prompt)
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
	kemonoCharacterChipText, err := generateTextByGPT4Vision(prompt)
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
	return generateTextByGPT4Vision(prompt)
}

func GenerateBattleText(attacker *domains.Kemono, defender *domains.Kemono, damage int) (*string, error) {
	prompt, err := generateBattleTextPrompt(attacker, defender, damage)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT3Dot5(prompt)
}

func BreedKemonoPrompt(kemonoParent1 *domains.Kemono, kemonoParent2 *domains.Kemono) (*string, error) {
	prompt, err := generateBreedKemonoPromptPrompt(kemonoParent1, kemonoParent2)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4Vision(prompt)
}

func BreedKemonoDescription(kemonoParent1 *domains.Kemono, kemonoParent2 *domains.Kemono, kemono *domains.Kemono) (*string, error) {
	prompt, err := generateBreedKemonoDescriptionPrompt(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return nil, err
	}
	return generateTextByGPT4Vision(prompt)
}

func BreedKemonoStatus(kemonoParent1 *domains.Kemono, kemonoParent2 *domains.Kemono, kemono *domains.Kemono) (*domains.KemonoStatus, error) {
	prompt, err := generateBreedKemonoStatusPrompt(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return nil, err
	}
	kemonoStatusText, err := generateTextByGPT4Vision(prompt)
	if err != nil {
		return nil, err
	}
	return domains.ParseKemonoStatus(kemonoStatusText)
}

func BreedKemonoConcepts(kemonoParent1 *domains.Kemono, kemonoParent2 *domains.Kemono, kemono *domains.Kemono) (*domains.Concepts, error) {
	prompt, err := generateBreedKemonoConceptsPrompt(kemonoParent1, kemonoParent2, kemono)
	if err != nil {
		return nil, err
	}
	kemonoConceptsText, err := generateTextByGPT4Vision(prompt)
	if err != nil {
		return nil, err
	}
	return domains.ParseKemonoConcepts(kemonoConceptsText)
}
