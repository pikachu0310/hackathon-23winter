package api

import (
	"fmt"
	"github.com/pikachu0310/hackathon-23winter/internal/domains"
	"github.com/pikachu0310/hackathon-23winter/src/images"
	"strings"
)

func createKemonoPromptPrompt(concepts domains.Concepts) (messages ChatMessages, err error) {
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

	var userContent1 MessageContents
	err = userContent1.AddText("以下の文章から、画像生成モデルDALL-E 3に画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。\n\nあなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\nまずは、ケモノのキャラクターを生成してください。\n特徴は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\n生成する際は、以下の事を注意して守ってください。\n- キャラクターは1体\n- ゲームの敵として出てきても違和感がない\n- そのままゲームに使える\n- デザイン用のカラーを含まない")
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

	var userContent2 MessageContents
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

func createKemonoDescriptionPrompt(concepts domains.Concepts, image []byte) (messages ChatMessages, err error) {
	/*
		提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。

		画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		参考までに、このケモノは以下の特性を持っています。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。出力する際は、特徴だけを出力してください。
	*/
	/*
		この愛らしい生き物は、彩り鮮やかな森の中で目覚めた不思議な存在です。水属性を持ち、四足での敏捷な歩行が特徴のマスコットキャラクターで、見る者を魅了するそのかわいらしい表情と澄んだ瞳が印象的です。鮮やかな色合いの体毛は、森の生態系に完璧に溶け込みながらも、その明るい色彩で存在感を放ちます。 水の力を操る能力を秘めており、その姿は水辺で遊ぶ姿が想像されるほど親しみやすいです。
	*/
	/*
		完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。

		画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		参考までに、このケモノは以下の特性を持っています。
		- %s

		提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。出力する際は、特徴だけを出力してください。
	*/
	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return nil, err
	}
	err = userContent1.AddText("提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。\n\n画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\n参考までに、このケモノは以下の特性を持っています。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\n提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、100文字程度で出力してください。出力する際は、特徴だけを出力してください。")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("この愛らしい生き物は、彩り鮮やかな森の中で目覚めた不思議な存在です。水属性を持ち、四足での敏捷な歩行が特徴のマスコットキャラクターで、見る者を魅了するそのかわいらしい表情と澄んだ瞳が印象的です。鮮やかな色合いの体毛は、森の生態系に完璧に溶け込みながらも、その明るい色彩で存在感を放ちます。 水の力を操る能力を秘めており、その姿は水辺で遊ぶ姿が想像されるほど親しみやすいです。")

	var contents MessageContents
	err = contents.AddImage(image)
	if err != nil {
		return
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、100文字程度で出力してください。\n\n画像に書かれているのは、僕が考えたゲームのキャラクターです。このキャラクターは、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームのキャラクターです。あなたには、この世界で生息しているこの画像のキャラクターがどういった生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\n参考までに、このケモノは以下の特性を持っています。")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\n提供した画像を読み込んで、画像に書かれているキャラクターの性質や性格や、特徴やできることを推測して、200文字程度で出力してください。出力する際は、特徴だけを出力してください。")
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

func createKemonoStatusPrompt(description *string, concepts domains.Concepts, image []byte) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。
		ケモノのキャラクターのステータスパラメーターを考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。
		ケモノのキャラクターのステータスパラメーターを出力する際は、以下のフォーマットに従ってください。
		```
		MaxHP=100
		Attack=10
		Defence=5
		```

		ケモノのキャラクターの特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクターが持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		それでは、考えたケモノのキャラクターのステータスパラメーターを以下のフォーマットに従って出力してください。
		```
		MaxHP=100
		Attack=10
		Defence=5
		```
	*/
	/*
		MaxHP=70
		Attack=8
		Defence=6
	*/
	/*
		完璧です！まったく同じようにして、以下のケモノのステータスパラメーターも考えてください。

		ケモノのキャラクターの特徴は以下の通りです。
		- %s

		ケモノのキャラクターが持つ概念は以下の通りです。
		- %s

		それでは、考えたケモノのキャラクターのステータスパラメーターを以下のフォーマットに従って出力してください。
		```
		MaxHP=100
		Attack=10
		Defence=5
		```
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。\nケモノのキャラクターのステータスパラメーターを考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。\nケモノのキャラクターのステータスパラメーターを出力する際は、以下のフォーマットに従ってください。\n```\nMaxHP=100\nAttack=10\nDefence=5\n```\n\nケモノのキャラクターの特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクターが持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 炎属性\n\nそれでは、考えたケモノのキャラクターのステータスパラメーターを以下のフォーマットに従って出力してください。\n```\nMaxHP=100\nAttack=10\nDefence=5\n```")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("MaxHP=70\nAttack=8\nDefence=6")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！まったく同じようにして、以下のケモノのステータスパラメーターも考えてください。\n\nケモノのキャラクターの特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *description))
	promptTexts = append(promptTexts, "\nケモノのキャラクターが持つ概念は以下の通りです。")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nそれでは、考えたケモノのキャラクターのステータスパラメーターを以下のフォーマットに従って出力してください。\n```\nMaxHP=100\nAttack=10\nDefence=5\n```")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 MessageContents
	err = userContent2.AddImage(image)
	if err != nil {
		return
	}
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

func createKemonoCharacterChipPrompt(description *string, concepts domains.Concepts, image []byte) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。

		動物の表
		| 00     | 01         | 02   | 03  | 04  | 05   | 06       | 07           |
		| ------ | ---------- | ---- | --- | --- | ---- | -------- | ------------ |
		| キツネ | フェレット | ネコ | 魚  | 犬  | ヘビ | コウモリ | カーバンクル |

		色の表
		| 00  | 01  | 02  | 03  | 04       | 05  | 06   | 07  | 08  | 09  | 10  | 11     | 12  |
		| --- | --- | --- | --- | -------- | --- | ---- | --- | --- | --- | --- | ------ | --- |
		| 白  | 灰  | 黒  | 赤  | オレンジ | 黄  | 黄緑 | 緑  | 水  | 青  | 紫  | ピンク | 茶  |

		ケモノのキャラクターの特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクターが持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。
		出力する際は、以下のフォーマットで出力してください。
		```
		動物=00
		色=04
		```
	*/
	/*
		動物=00
		色=08
	*/
	/*
		完璧です！まったく同じようにして、以下のケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。

		動物の表
		| 00     | 01         | 02   | 03  | 04  | 05   | 06       | 07           |
		| ------ | ---------- | ---- | --- | --- | ---- | -------- | ------------ |
		| キツネ | フェレット | ネコ | 魚  | 犬  | ヘビ | コウモリ | カーバンクル |

		色の表
		| 00  | 01  | 02  | 03  | 04       | 05  | 06   | 07  | 08  | 09  | 10  | 11     | 12  |
		| --- | --- | --- | --- | -------- | --- | ---- | --- | --- | --- | --- | ------ | --- |
		| 白  | 灰  | 黒  | 赤  | オレンジ | 黄  | 黄緑 | 緑  | 水  | 青  | 紫  | ピンク | 茶  |

		ケモノのキャラクターの特徴は以下の通りです。
		- %s

		ケモノのキャラクターが持つ概念は以下の通りです。
		- %s

		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。
		出力する際は、以下のフォーマットで出力してください。
		```
		動物=00
		色=04
		```
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。\n\n動物の表\n| 00     | 01         | 02   | 03  | 04  | 05   | 06       | 07           |\n| ------ | ---------- | ---- | --- | --- | ---- | -------- | ------------ |\n| キツネ | フェレット | ネコ | 魚  | 犬  | ヘビ | コウモリ | カーバンクル |\n\n色の表\n| 00  | 01  | 02  | 03  | 04       | 05  | 06   | 07  | 08  | 09  | 10  | 11     | 12  |\n| --- | --- | --- | --- | -------- | --- | ---- | --- | --- | --- | --- | ------ | --- |\n| 白  | 灰  | 黒  | 赤  | オレンジ | 黄  | 黄緑 | 緑  | 水  | 青  | 紫  | ピンク | 茶  |\n\nケモノのキャラクターの特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクターが持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。\n出力する際は、以下のフォーマットで出力してください。\n```\n動物=00\n色=04\n```")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("動物=00\n色=08")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！まったく同じようにして、以下のケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。\n\n動物の表\n| 00     | 01         | 02   | 03  | 04  | 05   | 06       | 07           |\n| ------ | ---------- | ---- | --- | --- | ---- | -------- | ------------ |\n| キツネ | フェレット | ネコ | 魚  | 犬  | ヘビ | コウモリ | カーバンクル |\n\n色の表\n| 00  | 01  | 02  | 03  | 04       | 05  | 06   | 07  | 08  | 09  | 10  | 11     | 12  |\n| --- | --- | --- | --- | -------- | --- | ---- | --- | --- | --- | --- | ------ | --- |\n| 白  | 灰  | 黒  | 赤  | オレンジ | 黄  | 黄緑 | 緑  | 水  | 青  | 紫  | ピンク | 茶  |\n\nケモノのキャラクターの特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *description))
	promptTexts = append(promptTexts, "\nケモノのキャラクターが持つ概念は以下の通りです。")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターが、以下の表のどの動物に最も近いか、そして以下の表のどの色に最も近いかを判断して、数字を出力してください。\n出力する際は、以下のフォーマットで出力してください。\n```\n動物=00\n色=04")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 MessageContents
	err = userContent2.AddImage(image)
	if err != nil {
		return
	}
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

func createKemonoNamePrompt(description *string, concepts domains.Concepts, image []byte) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターの名前を考えて出力してください。
		ケモノのキャラクターの名前を考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。
		ケモノのキャラクターの名前を出力する際は、名前を1個だけ出力してください。

		ケモノのキャラクターの特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクターが持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		では、このケモノキャラクターの名前を考えて出力してください。名前を出力する際は、名前だけを出力してください。
	*/
	/*
		アクアリン
	*/
	/*
		完璧です！まったく同じようにして、以下のケモノのキャラクターの名前も考えてください。
		提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターの名前を考えて出力してください。
		ケモノのキャラクターの名前を考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。
		ケモノのキャラクターの名前を出力する際は、名前を1個だけ出力してください。

		ケモノのキャラクターの特徴は以下の通りです。
		- %s

		ケモノのキャラクターが持つ概念は以下の通りです。
		- %s

		では、このケモノキャラクターの名前を考えて出力してください。名前を出力する際は、名前だけを出力してください。
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターの名前を考えて出力してください。\nケモノのキャラクターの名前を考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。\nケモノのキャラクターの名前を出力する際は、名前を1個だけ出力してください。\n\nケモノのキャラクターの特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクターが持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 炎属性\n\nでは、このケモノキャラクターの名前を考えて出力してください。名前を出力する際は、名前だけを出力してください。")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("アクアリン")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！まったく同じようにして、以下のケモノのキャラクターの名前も考えてください。\n提供した画像を読み込んで、提供した画像に書かれているケモノのキャラクターの名前を考えて出力してください。\nケモノのキャラクターの名前を考える際は、画像のキャラクターの見た目と、以下に書くキャラクターの特徴や概念を参考にしてください。\nケモノのキャラクターの名前を出力する際は、名前を1個だけ出力してください。\n\nケモノのキャラクターの特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *description))
	promptTexts = append(promptTexts, "\nケモノのキャラクターが持つ概念は以下の通りです。")
	for _, concept := range concepts {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nでは、このケモノキャラクターの名前を考えて出力してください。名前を出力する際は、名前だけを出力してください。")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 MessageContents
	err = userContent2.AddImage(image)
	if err != nil {
		return
	}
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
