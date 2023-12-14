package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/repository"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api"
	"github.com/pikachu0310/hackathon-23winter/src/images"
	"net/http"
	"strings"
)

type TestResponse struct {
	ImageString []byte `json:"image_string"`
}

type Test2Response struct {
	GeneratedText string `json:"generated_text"`
}

// GET /api/v1/test
func (h *Handler) Test(c echo.Context) error {
	image, err := api.GenerateKemonoImage()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	kemonoParams := &repository.KemonoParams{
		ID:    uuid.New(),
		Image: images.TestKemonoImageFire,
	}

	_, err = h.repo.CreateKemono(c.Request().Context(), kemonoParams.ToKemono())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := TestResponse{
		ImageString: *image,
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/test/2
func (h *Handler) Test2(c echo.Context) error {
	var messages api.ChatMessages
	var contents api.MessageContents

	err := contents.AddText("この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。\nいま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。\n\n以下にケモノ「アクアフローラ」の特徴を載せます。\n名前:「アクアフローラ」\n能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「100」\n現在の体力:「100」\n攻撃力:「10」\n防御力:「10」\n\n以下にケモノ「ヒバナ」の特徴を載せます。\n名前:「ヒバナ」\n能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「20」\n現在の体力:「20」\n攻撃力:「5」\n防御力:「3」\n\nこれからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「8」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「8」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を生成してください。")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = messages.AddUserMessageContent(contents)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	generatedText, err := api.GenerateTextByGPT4(messages)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := Test2Response{
		GeneratedText: *generatedText,
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/v1/test/3
func (h *Handler) Test3(c echo.Context) error {
	var messages api.ChatMessages
	var contents api.MessageContents

	err := contents.AddText("この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。\nいま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。\n\n以下にケモノ「アクアフローラ」の特徴を載せます。\n名前:「アクアフローラ」\n能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「100」\n現在の体力:「100」\n攻撃力:「10」\n防御力:「10」\n\n以下にケモノ「ヒバナ」の特徴を載せます。\n名前:「ヒバナ」\n能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「20」\n現在の体力:「20」\n攻撃力:「5」\n防御力:「3」\n\nこれからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「8」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「8」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を生成してください。")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = messages.AddUserMessageContent(contents)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	generatedText, err := api.GenerateTextByGPT4(messages)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := Test2Response{
		GeneratedText: *generatedText,
	}

	return c.JSON(http.StatusOK, res)
}

func GenerateImagePrompt(concepts repository.Concepts) (api.ChatMessages, error) {
	/*
		以下の文章から、画像生成モデルDALL-E 3に画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。

		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。
		まずは、ケモノのキャラクターを生成してください。
		特徴は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		生成する際は、以下の事を注意して守ってください。
		- キャラクターは1体
		- ゲームの敵として出てきても違和感がない
		- そのままゲームに使える
		- デザイン用のカラーを含まない
	*/
	/*
		Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.
	*/
	/*
		完璧です！では、全く同じようにして、以下の文章から、画像生成モデルDALL-E 3に画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。

		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。
		まずは、ケモノのキャラクターを生成してください。
		特徴は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		生成する際は、以下の事を注意して守ってください。
		- キャラクターは1体
		- ゲームの敵として出てきても違和感がない
		- そのままゲームに使える
		- デザイン用のカラーを含まない
	*/

	var messages api.ChatMessages
	var userContent1 api.MessageContents
	err := userContent1.AddText("以下の文章から、画像生成モデルDALL-E 3に画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。\n\nあなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\nまずは、ケモノのキャラクターを生成してください。\n特徴は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\n生成する際は、以下の事を注意して守ってください。\n- キャラクターは1体\n- ゲームの敵として出てきても違和感がない\n- そのままゲームに使える\n- デザイン用のカラーを含まない")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！では、全く同じようにして、以下の文章から、画像生成モデルDALL-E 3に画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。\n\nあなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\nまずは、ケモノのキャラクターを生成してください。\n特徴は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\n生成する際は、以下の事を注意して守ってください。\n- キャラクターは1体\n- ゲームの敵として出てきても違和感がない\n- そのままゲームに使える\n- デザイン用のカラーを含まない")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\n生成する際は、以下の事を注意して守ってください。\n- キャラクターは1体\n- ゲームの敵として出てきても違和感がない\n- そのままゲームに使える\n- デザイン用のカラーを含まない")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 api.MessageContents
	err = userContent2.AddText(promptText)
	if err != nil {
		return nil, err
	}

	err = messages.AddUserMessageContent(userContent2)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// GET /api/v1/test/4
func (h *Handler) Test4(c echo.Context) error {
	kemonoId, err := uuid.Parse("00000000-0000-0000-0000-000000000001")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}
	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	prompt, err := GenerateImagePrompt(kemono.Concepts.Concepts())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	imagePrompt, err := api.GenerateTextByGPT4(prompt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.String(http.StatusOK, *imagePrompt)
}

func GenerateDescriptionPrompt(concepts repository.Concepts, imageBase64 string) (messages api.ChatMessages, err error) {
	/*
		提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、400文字程度で出力してください。

		画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		参考までに、このケモノは以下の特性を持っています。
		- 炎属性

		提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、400文字程度で出力してください。
	*/
	var contents api.MessageContents
	err = contents.AddImage(imageBase64)
	if err != nil {
		return
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、400文字程度で出力してください。\n\n画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\n参考までに、このケモノは以下の特性を持っています。")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\n提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、400文字程度で出力してください。")
	promptText := strings.Join(promptTexts, "\n")
	err = contents.AddText(promptText)
	if err != nil {
		return
	}
	err = messages.AddUserMessageContent(contents)
	if err != nil {
		return
	}

	return messages, nil
}

// GET /api/v1/test/5 画像から説明の作成
func (h *Handler) Test5(c echo.Context) error {
	kemonoID, err := uuid.Parse("00000000-0000-0000-0000-000000000001")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid kemonoID").SetInternal(err)
	}

	kemono, err := h.repo.GetKemono(c.Request().Context(), kemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	prompt, err := GenerateDescriptionPrompt(kemono.Concepts.Concepts(), api.ImageToBase64(kemono.Image))
	if err != nil {
		return err
	}

	kemonoDescription, err := api.GenerateTextByGPT4(prompt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	kemono.Description = kemonoDescription

	err = h.repo.UpdateKemono(c.Request().Context(), kemono)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.String(http.StatusOK, *kemonoDescription)
}
