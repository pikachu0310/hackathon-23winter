package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/pkg/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	apiServerURL = "https://api.tringpt.com/api"
)

var (
	model = "supertrin_highpres"
)

// RequestBody defines the structure of the API request body.
type RequestBody struct {
	Text            string   `json:"text"`
	Length          int      `json:"length"`                      // トークン数（1～300）
	Temperature     float64  `json:"temperature"`                 // ランダム度（0～2.5）
	TopP            float64  `json:"top_p"`                       // Top Pサンプリング（0.01～1.0）
	RepPen          float64  `json:"rep_pen"`                     // 繰り返しペナルティ（1.0～2.0）
	TopK            *int     `json:"top_k,omitempty"`             // Top Kサンプリング（1～500）
	TopA            *float64 `json:"top_a,omitempty"`             // Top Aサンプリング（0～1.0）
	TailFree        *float64 `json:"tailfree,omitempty"`          // Tail-freeサンプリング（0.01～1.0）
	RepPenRange     *int     `json:"rep_pen_range,omitempty"`     // 繰り返しペナルティ範囲（0～2048）
	RepPenSlope     *float64 `json:"rep_pen_slope,omitempty"`     // 繰り返しペナルティの傾斜（0.01～10）
	RepPenPres      *int     `json:"rep_pen_pres,omitempty"`      // コンテキスト依存の繰り返しペナルティ（0～100）
	TypicalP        *float64 `json:"typical_p,omitempty"`         // Typicalサンプリング（0.01～1.0）
	BadWords        *string  `json:"badwords,omitempty"`          // 禁止ワード（区切り文字: <<|>>）
	LogitBias       *string  `json:"logit_bias,omitempty"`        // トークンの出現率調整（区切り文字: <<|>>）
	LogitBiasValues *string  `json:"logit_bias_values,omitempty"` // logit spaceにおける確率（区切り文字: |）
	StopTokens      *string  `json:"stoptokens,omitempty"`        // 出力打ち切りシーケンス（区切り文字: <<|>>）
	Model           *string  `json:"model,omitempty"`             // 使用するモデル（例: supertrin_highpres）
}

// ResponseData defines the structure for the API response data.
type ResponseData struct {
	Data []string `json:"data"`
}

func GetGeneratedText(text string) (*string, error) {
	reqBody := NewRequestBody(text)
	return GenerateText(reqBody)
}

func NewRequestBody(text string) *RequestBody {
	return &RequestBody{
		Text:        text,
		Length:      20,
		Temperature: 0.5,
		TopP:        0.7,
		RepPen:      2.0,
	}
}

// GenerateText sends a request to the AI writing API and returns the response.
func GenerateText(reqBody *RequestBody) (*string, error) {
	jsonData, err := json.Marshal(reqBody)
	fmt.Printf("Request body: %s\n", string(jsonData))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiServerURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+config.GetNovelistAPIKey())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%#v\n", string(body))

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, err
	}

	if len(responseData.Data) != 1 {
		log.Fatal("responseData.Data is not 1 length" + strconv.Itoa(len(responseData.Data)))
	}

	return &responseData.Data[0], nil
}
