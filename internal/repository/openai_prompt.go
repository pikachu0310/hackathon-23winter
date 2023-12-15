package repository

import (
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api"
	"strings"
)

func GenerateImagePrompt(concepts Concepts) (api.ChatMessages, error) {
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

func GenerateDescriptionPrompt(concepts Concepts, imageBase64 string) (messages api.ChatMessages, err error) {
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
