package api

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/pkg/config"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api/openai"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api/openai/openai_api"
	"log"
	"net/http"
)

const (
	openaiServerUrl = "https://api.openai.com/v1"
)

var (
	client *openai.ClientWithResponses

	createImageModel   openai_api.CreateImageRequest_Model
	createImageN       = 1
	createImageQuality = openai_api.Standard
	createImageFormat  = openai_api.CreateImageRequestResponseFormatB64Json
	createImageSize    = openai_api.CreateImageRequestSizeN1024x1024
	createImageStyle   = openai_api.Vivid
	createImageUser    = "pikachu0310"

	createChatCompletionModel     openai_api.CreateChatCompletionRequest_Model
	createChatCompletionMaxTokens = 4096
	createChatCompletionN         = 1
	createChatCompletionStream    = false
)

type ChatMessage openai_api.ChatCompletionRequestMessage

func init() {
	var err error
	client, err = openai.NewClientWithResponses(openaiServerUrl, openai.WithRequestEditorFn(addAPIKeyToRequest(config.GetOpenaiAPIKey())))
	if err != nil {
		log.Panic(err)
	}

	err = createImageModel.FromCreateImageRequestModel1(openai_api.CreateImageRequestModel1DallE3)
	if err != nil {
		log.Panic(err)
	}

	err = createChatCompletionModel.FromCreateChatCompletionRequestModel1(openai_api.CreateChatCompletionRequestModel1Gpt4VisionPreview)
	if err != nil {
		log.Panic(err)
	}
}

type ChatMessages []openai_api.ChatCompletionRequestMessage
type MessageContents openai_api.ChatCompletionRequestUserMessageContent1

func (chatMessage *ChatMessages) AddUserMessageContent(messageContents MessageContents) error {
	var content openai_api.ChatCompletionRequestUserMessage_Content
	err := content.FromChatCompletionRequestUserMessageContent1(messageContents)
	if err != nil {
		return err
	}
	userMessage := openai_api.ChatCompletionRequestUserMessage{
		Content: content,
		Role:    openai_api.ChatCompletionRequestUserMessageRoleUser,
	}
	var message openai_api.ChatCompletionRequestMessage
	err = message.FromChatCompletionRequestUserMessage(userMessage)
	if err != nil {
		return err
	}
	*chatMessage = append(*chatMessage, message)
	return nil
}

func (chatMessage *ChatMessages) AddAssistantMessageContent(text string) error {
	assistantMessage := openai_api.ChatCompletionRequestAssistantMessage{
		Content: &text,
		Role:    openai_api.ChatCompletionRequestAssistantMessageRoleAssistant,
	}

	var message openai_api.ChatCompletionRequestMessage
	err := message.FromChatCompletionRequestAssistantMessage(assistantMessage)
	if err != nil {
		return err
	}

	*chatMessage = append(*chatMessage, message)
	return nil
}

func (chatMessage *ChatMessages) AddSystemMessage(text string) error {
	systemMessage := openai_api.ChatCompletionRequestSystemMessage{
		Content: text,
		Role:    openai_api.System,
	}

	var message openai_api.ChatCompletionRequestMessage
	err := message.FromChatCompletionRequestSystemMessage(systemMessage)
	if err != nil {
		return err
	}

	*chatMessage = append(*chatMessage, message)
	return nil
}

func (messageContents *MessageContents) AddText(text string) error {
	var messageContentPart openai_api.ChatCompletionRequestMessageContentPart
	var messageContentPartText openai_api.ChatCompletionRequestMessageContentPartText
	messageContentPartText.Text = text
	messageContentPartText.Type = openai_api.ChatCompletionRequestMessageContentPartTextTypeText

	err := messageContentPart.FromChatCompletionRequestMessageContentPartText(messageContentPartText)
	if err != nil {
		return err
	}

	*messageContents = append(*messageContents, messageContentPart)
	return nil
}

func (messageContents *MessageContents) AddImage(imageBase64 string) error {
	var messageContentPart openai_api.ChatCompletionRequestMessageContentPart
	var messageContentPartImage openai_api.ChatCompletionRequestMessageContentPartImage
	messageContentPartImage.ImageUrl = struct {
		Detail *openai_api.ChatCompletionRequestMessageContentPartImageImageUrlDetail `json:"detail,omitempty"`
		Url    string                                                                 `json:"url"`
	}{Url: imageBase64}
	messageContentPartImage.Type = openai_api.ImageUrl

	err := messageContentPart.FromChatCompletionRequestMessageContentPartImage(messageContentPartImage)
	if err != nil {
		return err
	}

	*messageContents = append(*messageContents, messageContentPart)
	return nil
}

func addAPIKeyToRequest(apiKey string) openai.RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "Bearer "+apiKey)
		return nil
	}
}

func GenerateKemonoImage() (*[]byte, error) {
	req := openai_api.CreateImageRequest{
		Model:          &createImageModel,
		N:              &createImageN,
		Prompt:         "Create a digital illustration of a single cute, four-legged kemomimi character that would fit as a mascot in a vibrant game world inhabited by various creatures. The character has just awakened in a lush forest, reminiscent of a Pokémon in style. It should have bright, eye-catching colors indicative of its water element, suggesting magical abilities. The character must look fitting to encounter as a potential friendly or enemy entity in a fantasy game environment, and it should be designed ready to use in a game without any additional color palette indicators or design drafts.",
		Quality:        &createImageQuality,
		ResponseFormat: &createImageFormat,
		Size:           &createImageSize,
		Style:          &createImageStyle,
		User:           &createImageUser,
	}
	res, err := client.CreateImageWithResponse(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if res.HTTPResponse.StatusCode != 200 {
		return nil, fmt.Errorf("failed to generate kemono image: %s", string(res.Body))
	}

	base64String := *res.JSON200.Data[0].B64Json

	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func GenerateTextByGPT4(messages ChatMessages) (*string, error) {
	req := openai_api.CreateChatCompletionRequest{
		Model:     createChatCompletionModel,
		MaxTokens: &createChatCompletionMaxTokens,
		Messages:  messages,
		N:         &createChatCompletionN,
		Stream:    &createChatCompletionStream,
	}

	res, err := client.CreateChatCompletionWithResponse(context.Background(), req)
	if err != nil {
		return nil, err
	}
	if res.HTTPResponse.StatusCode != 200 {
		return nil, fmt.Errorf("failed to generate text by GPT-4: %s", string(res.Body))
	}

	responseMessages := res.JSON200
	latestResponseMessage := responseMessages.Choices[len(responseMessages.Choices)-1]

	return latestResponseMessage.Message.Content, nil
}