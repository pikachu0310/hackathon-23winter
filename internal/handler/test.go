package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/repository"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api"
	"github.com/pikachu0310/hackathon-23winter/src/images"
	"net/http"
)

type TestResponse struct {
	ImageString []byte `json:"image_string"`
}

type Test2Response struct {
	GeneratedText string `json:"generated_text"`
}

// GET /api/v1/test
func (h *Handler) Test(c echo.Context) error {
	image, err := api.GenerateKemonoImage("test")
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

	prompt, err := repository.GenerateImagePrompt(kemono.Concepts.Concepts())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}
	imagePrompt, err := api.GenerateTextByGPT4(prompt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError).SetInternal(err)
	}

	return c.String(http.StatusOK, *imagePrompt)
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

	prompt, err := repository.GenerateDescriptionPrompt(kemono.Concepts.Concepts(), api.ImageToBase64(kemono.Image))
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
