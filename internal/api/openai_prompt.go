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

func generateBreedKemonoPromptPrompt(kemono1 domains.Kemono, kemono2 domains.Kemono) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像2枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しようと思っています。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、特徴や詳細からケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		ケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 炎属性

		ケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		また、プロンプトを生成する際、以下の事を守らせるような文章を追加してください。
		- 画面中央にケモノが1体だけいる
		- ケモノだけが画像に含まれる
		- ゲームの敵として出てきても違和感がなく、そのままゲームに使える
		- デザイン用のカラーを含まない
		- 文字を含まない

		以上の2匹のケモノキャラクターの特徴や詳細から、この2匹のケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。プロンプトは非常に長くても大丈夫です。正確に上記の要素で生成できるように、正確で長いプロンプトを出力してください。
	*/
	/*
		Imagine a single, enchanting creature that is a perfect fusion of two mystical kemomono characters, one infused with the essence of water and the other with the essence of fire. This new character should be a delightful blend of both, featuring a four-legged form that radiates the charm of a cute, approachable mascot. Its design should showcase a magnificent coat that dances with vivid hues of blue and orange, harmoniously representing its dual elemental nature. The creature's ears and tail should be distinctively fluffy, with the tail artistically merging aqueous patterns and fiery designs, a true testament to its combined powers. Its eyes should gleam with a gentle yet unwavering spirit, a mirror to its kind demeanor and its fierce abilities to conjure both water and fire magic. As a potential adversary in a game, this character would be both a guide and a challenger to adventurers, its appearance bringing life to the heart of the forest it protects. The image should center on this kemomono, poised and ready for inclusion in a gaming world, devoid of any extraneous design markers, color palettes, or text, to ensure an uninterrupted experience for the player.
	*/
	/*
		完璧です！まったく同じようにして、先ほどとは異なるケモノの画像と特徴をまた提供するので、画像を生成するためのプロンプトを考えて出力してください。

		提供した画像2枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しようと思っています。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、特徴や詳細からケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		ケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 炎属性

		ケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		また、プロンプトを生成する際、以下の事を守らせるような文章を追加してください。
		- 画面中央にケモノが1体だけいる
		- ケモノだけが画像に含まれる
		- ゲームの敵として出てきても違和感がなく、そのままゲームに使える
		- デザイン用のカラーを含まない
		- 文字を含まない

		以上の2匹のケモノキャラクターの特徴や詳細から、この2匹のケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。プロンプトは非常に長くても大丈夫です。正確に上記の要素で生成できるように、正確で長いプロンプトを出力してください。
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddImage(images.TestKemonoImageFire)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像2枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しようと思っています。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、特徴や詳細からケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。\n\nケモノのキャラクター1匹目の特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクター1匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\nケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。\n- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.\n\nケモノのキャラクター2匹目の特徴は以下の通りです。\n- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。\n\nケモノのキャラクター2匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 炎属性\n\nケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。\n- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.\n\nまた、プロンプトを生成する際、以下の事を守らせるような文章を追加してください。\n- 画面中央にケモノが1体だけいる\n- ケモノだけが画像に含まれる\n- ゲームの敵として出てきても違和感がなく、そのままゲームに使える\n- デザイン用のカラーを含まない\n- 文字を含まない\n\n以上の2匹のケモノキャラクターの特徴や詳細から、この2匹のケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。プロンプトは非常に長くても大丈夫です。正確に上記の要素で生成できるように、正確で長いプロンプトを出力してください。")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("Imagine a single, enchanting creature that is a perfect fusion of two mystical kemomono characters, one infused with the essence of water and the other with the essence of fire.")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！まったく同じようにして、先ほどとは異なるケモノの画像と特徴をまた提供するので、画像を生成するためのプロンプトを考えて出力してください。\n\n提供した画像2枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しようと思っています。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、特徴や詳細からケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。\n\nケモノのキャラクター1匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono1.Description)
	promptTexts = append(promptTexts, "\nケモノのキャラクター1匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono1.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono1.Prompt)

	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono2.Description)
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono2.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono2.Prompt)
	promptTexts = append(promptTexts, "\nまた、プロンプトを生成する際、以下の事を守らせるような文章を追加してください。\n- 画面中央にケモノが1体だけいる\n- ケモノだけが画像に含まれる\n- ゲームの敵として出てきても違和感がなく、そのままゲームに使える\n- デザイン用のカラーを含まない\n- 文字を含まない\n\n以上の2匹のケモノキャラクターの特徴や詳細から、この2匹のケモノのキャラクターを配合させた後どんな姿になるかを考え、画像生成モデルDALL-E 3に配合されたケモノキャラクターの画像を生成させるためのプロンプトを生成して、そのプロンプトだけを出力してください。プロンプトは非常に長くても大丈夫です。正確に上記の要素で生成できるように、正確で長いプロンプトを出力してください。")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 MessageContents
	err = userContent2.AddImage(kemono1.Image)
	if err != nil {
		return
	}
	err = userContent2.AddImage(kemono2.Image)
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

func generateBreedKemonoDescriptionPrompt(kemono1 domains.Kemono, kemono2 domains.Kemono, kemono3 domains.Kemono) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちだけが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		ケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 炎属性

		ケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。
		- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.

		以上の2匹のケモノキャラクターの特徴や詳細と、配合後のケモノの姿から、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持ち、どのような戦闘能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力してください。300文字程度で出力してください。出力する際は、説明文を日本語で、説明文だけを出力してください。
	*/
	/*
		この新たなケモノキャラクターは、水と炎の属性を併せ持つ森の守護者として生息しています。そのふわふわの毛並みは水の精霊の柔らかさを受け継ぎ、鮮やかな色彩は森の生命力と情熱を象徴しています。目には優しさと決意が宿り、敵には水と炎の魔法で立ち向かう力を秘めており、その勇敢さは森を守る戦士の証です。四足歩行で、朝目覚めれば森の中におり、ポケモンのような可愛らしさでありながら、戦いではその力を存分に発揮する生き物です。
	*/
	/*
		完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した3枚の画像を読み込んで、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- %s

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- %s

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- %s

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- %s

		以上の2匹のケモノキャラクターの特徴や詳細と、配合後のケモノの姿から、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持ち、どのような戦闘能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力してください。300文字程度で出力してください。出力する際は、説明文を日本語で、説明文だけを出力してください。
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddImage(images.TestKemonoImageFire)
	if err != nil {
		return
	}
	err = userContent1.AddImage(images.TestKemonoImageAquaFire)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちだけが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\nケモノのキャラクター1匹目の特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクター1匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\nケモノのキャラクター1匹目の画像が生成されるときに使われたプロンプトは以下の通りです。\n- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.\n\nケモノのキャラクター2匹目の特徴は以下の通りです。\n- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。\n\nケモノのキャラクター2匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 炎属性\n\nケモノのキャラクター2匹目の画像が生成されるときに使われたプロンプトは以下の通りです。\n- Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.\n\n以上の2匹のケモノキャラクターの特徴や詳細と、配合後のケモノの姿から、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持ち、どのような戦闘能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力してください。300文字程度で出力してください。出力する際は、説明文を日本語で、説明文だけを出力してください。")
	if err != nil {
		return nil, err
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return nil, err
	}
	err = messages.AddAssistantMessageContent("この新たなケモノキャラクターは、水と炎の属性を併せ持つ森の守護者として生息しています。そのふわふわの毛並みは水の精霊の柔らかさを受け継ぎ、鮮やかな色彩は森の生命力と情熱を象徴しています。目には優しさと決意が宿り、敵には水と炎の魔法で立ち向かう力を秘めており、その勇敢さは森を守る戦士の証です。四足歩行で、朝目覚めれば森の中におり、ポケモンのような可愛らしさでありながら、戦いではその力を存分に発揮する生き物です。")
	if err != nil {
		return nil, err
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した3枚の画像を読み込んで、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\n提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノキャラクターと2匹目のケモノキャラクターの特徴や詳細を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力して欲しいです。\n\nケモノのキャラクター1匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono1.Description)
	promptTexts = append(promptTexts, "\nケモノのキャラクター1匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono1.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, "- %s", *kemono2.Description)
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono2.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\n以上の2匹のケモノキャラクターの特徴や詳細と、配合後のケモノの姿から、配合後のモンスターである3枚目の画像のケモノのキャラクターがこの世界でどのように生息していて、どういった能力を持ち、どのような戦闘能力を持つ生き物なのかを、推測し、考えて、ゲーム内で使用する説明文として出力してください。300文字程度で出力してください。出力する際は、説明文を日本語で、説明文だけを出力してください。")
	promptText := strings.Join(promptTexts, "\n")

	var userContent2 MessageContents
	err = userContent2.AddImage(kemono1.Image)
	if err != nil {
		return
	}
	err = userContent2.AddImage(kemono2.Image)
	if err != nil {
		return
	}
	err = userContent2.AddImage(kemono3.Image)
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

func generateBreedKemonoStatusPrompt(kemono1 domains.Kemono, kemono2 domains.Kemono, kemono3 domains.Kemono) (messages ChatMessages, err error) {
	/*
		あなたには、かわいいマスコットやケモノやマモノたちだけが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。

		提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノのキャラクターと2匹目のケモノのキャラクターの特徴や詳細とステータスパラメーター(最大HPと攻撃力と防御力)を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメーターは以下のフォーマットで表されるので、出力する際も以下のフォーマットに従ってください。
		```
		MaxHP=100
		Attack=10
		Defence=5
		```

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		ケモノのキャラクター1匹目のステータスパラメーターは以下の通りです。
		```
		MaxHP=100
		Attack=10
		Defence=6
		```

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 炎属性

		ケモノのキャラクター2匹目のステータスパラメーターは以下の通りです。
		```
		MaxHP=70
		Attack=18
		Defence=3
		```

		交配後のケモノのキャラクターの特徴は以下の通りです。
		- この新たなケモノキャラクターは、水と炎の属性を併せ持つ森の守護者として生息しています。そのふわふわの毛並みは水の精霊の柔らかさを受け継ぎ、鮮やかな色彩は森の生命力と情熱を象徴しています。目には優しさと決意が宿り、敵には水と炎の魔法で立ち向かう力を秘めており、その勇敢さは森を守る戦士の証です。四足歩行で、朝目覚めれば森の中におり、ポケモンのような可愛らしさでありながら、戦いではその力を存分に発揮する生き物です。

		以上の3匹のケモノキャラクターの特徴や詳細から、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメータを出力する際は、先ほど与えたフォーマットに従ってください。基本的に、配合ではステータスパラメーターは配合前の2匹のステータスパラメーターの平均よりも強くなるが、そこまで大幅に強くなることもないことに留意してください。ただし、配合前の二匹のケモノの相性がとても良いときは、大幅に強くなることが出来ます。
	*/
	/*
		MaxHP=95
		Attack=17
		Defence=5
	*/
	/*
		完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した3枚の画像を読み込んで、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。

		提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。
		ケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノのキャラクターと2匹目のケモノのキャラクターの特徴や詳細とステータスパラメーター(最大HPと攻撃力と防御力)を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメーターは以下のフォーマットで表されるので、出力する際も以下のフォーマットに従ってください。
		```
		MaxHP=100
		Attack=10
		Defence=5
		```

		ケモノのキャラクター1匹目の特徴は以下の通りです。
		- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。

		ケモノのキャラクター1匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 水属性

		ケモノのキャラクター1匹目のステータスパラメーターは以下の通りです。
		```
		MaxHP=100
		Attack=10
		Defence=6
		```

		ケモノのキャラクター2匹目の特徴は以下の通りです。
		- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。

		ケモノのキャラクター2匹目が持つ概念は以下の通りです。
		- とてもかわいい
		- マスコット
		- 四足歩行
		- 目が覚めたら森の中だった
		- ポケモンのようなイメージ
		- 色は鮮やかめ
		- 炎属性

		ケモノのキャラクター2匹目のステータスパラメーターは以下の通りです。
		```
		MaxHP=70
		Attack=18
		Defence=3
		```

		交配後のケモノのキャラクターの特徴は以下の通りです。
		- この新たなケモノキャラクターは、水と炎の属性を併せ持つ森の守護者として生息しています。そのふわふわの毛並みは水の精霊の柔らかさを受け継ぎ、鮮やかな色彩は森の生命力と情熱を象徴しています。目には優しさと決意が宿り、敵には水と炎の魔法で立ち向かう力を秘めており、その勇敢さは森を守る戦士の証です。四足歩行で、朝目覚めれば森の中におり、ポケモンのような可愛らしさでありながら、戦いではその力を存分に発揮する生き物です。

		以上の3匹のケモノキャラクターの特徴や詳細から、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメータを出力する際は、先ほど与えたフォーマットに従ってください。基本的に、配合ではステータスパラメーターは配合前の2匹のステータスパラメーターの平均よりも強くなるが、そこまで大幅に強くなることもないことに留意してください。ただし、配合前の二匹のケモノの相性がとても良いときは、大幅に強くなることが出来ます。
	*/

	var userContent1 MessageContents
	err = userContent1.AddImage(images.TestKemonoImageAqua)
	if err != nil {
		return
	}
	err = userContent1.AddImage(images.TestKemonoImageFire)
	if err != nil {
		return
	}
	err = userContent1.AddImage(images.TestKemonoImageAquaFire)
	if err != nil {
		return
	}
	err = userContent1.AddText("あなたには、かわいいマスコットやケモノやマモノたちだけが生息する世界観のゲームの、ゲーム内システムを担当してもらいます。\n\n提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノのキャラクターと2匹目のケモノのキャラクターの特徴や詳細とステータスパラメーター(最大HPと攻撃力と防御力)を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメーターは以下のフォーマットで表されるので、出力する際も以下のフォーマットに従ってください。\n```\nMaxHP=100\nAttack=10\nDefence=5\n```\n\nケモノのキャラクター1匹目の特徴は以下の通りです。\n- このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。\n\nケモノのキャラクター1匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 水属性\n\nケモノのキャラクター1匹目のステータスパラメーターは以下の通りです。\n```\nMaxHP=100\nAttack=10\nDefence=6\n```\n\nケモノのキャラクター2匹目の特徴は以下の通りです。\n- このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。\n\nケモノのキャラクター2匹目が持つ概念は以下の通りです。\n- とてもかわいい\n- マスコット\n- 四足歩行\n- 目が覚めたら森の中だった\n- ポケモンのようなイメージ\n- 色は鮮やかめ\n- 炎属性\n\nケモノのキャラクター2匹目のステータスパラメーターは以下の通りです。\n```\nMaxHP=70\nAttack=18\nDefence=3\n```\n\n交配後のケモノのキャラクターの特徴は以下の通りです。\n- この新たなケモノキャラクターは、水と炎の属性を併せ持つ森の守護者として生息しています。そのふわふわの毛並みは水の精霊の柔らかさを受け継ぎ、鮮やかな色彩は森の生命力と情熱を象徴しています。目には優しさと決意が宿り、敵には水と炎の魔法で立ち向かう力を秘めており、その勇敢さは森を守る戦士の証です。四足歩行で、朝目覚めれば森の中におり、ポケモンのような可愛らしさでありながら、戦いではその力を存分に発揮する生き物です。\n\n以上の3匹のケモノキャラクターの特徴や詳細から、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメータを出力する際は、先ほど与えたフォーマットに従ってください。基本的に、配合ではステータスパラメーターは配合前の2匹のステータスパラメーターの平均よりも強くなるが、そこまで大幅に強くなることもないことに留意してください。ただし、配合前の二匹のケモノの相性がとても良いときは、大幅に強くなることが出来ます。")
	if err != nil {
		return
	}
	err = messages.AddUserMessageContent(userContent1)
	if err != nil {
		return
	}
	err = messages.AddAssistantMessageContent("MaxHP=95\nAttack=17\nDefence=5")
	if err != nil {
		return
	}

	var promptTexts []string
	promptTexts = append(promptTexts, "完璧です！また別のケモノキャラクターを送るので、まったく同じようにして、提供した3枚の画像を読み込んで、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。\n\n提供した画像3枚は、それぞれケモノのキャラクターの画像です。1匹目のケモノキャラクターと2匹目のケモノキャラクターを配合させ、新しいケモノキャラクターを生成しました。3枚目の画像が、配合されたケモノのキャラクターの画像です。\nケモノのキャラクターの画像を読み込んだうえで、1匹目のケモノのキャラクターと2匹目のケモノのキャラクターの特徴や詳細とステータスパラメーター(最大HPと攻撃力と防御力)を以下に書くので、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメーターは以下のフォーマットで表されるので、出力する際も以下のフォーマットに従ってください。\n```\nMaxHP=100\nAttack=10\nDefence=5\n```\n\nケモノのキャラクター1匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *kemono1.Description))
	promptTexts = append(promptTexts, "\nケモノのキャラクター1匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono1.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nケモノのキャラクター1匹目のステータスパラメーターは以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("```\nMaxHP=%d\nAttack=%d\nDefence=%d\n```", kemono1.MaxHp, kemono1.Attack, kemono1.Defense))
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目の特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *kemono2.Description))
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目が持つ概念は以下の通りです。")
	for _, concept := range kemono2.Concepts.Concepts() {
		promptTexts = append(promptTexts, fmt.Sprintf("- %s", concept))
	}
	promptTexts = append(promptTexts, "\nケモノのキャラクター2匹目のステータスパラメーターは以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("```\nMaxHP=%d\nAttack=%d\nDefence=%d\n```", kemono2.MaxHp, kemono2.Attack, kemono2.Defense))
	promptTexts = append(promptTexts, "\n交配後のケモノのキャラクターの特徴は以下の通りです。")
	promptTexts = append(promptTexts, fmt.Sprintf("- %s", *kemono3.Description))
	promptTexts = append(promptTexts, "\n以上の3匹のケモノキャラクターの特徴や詳細から、配合後のモンスターである3枚目の画像のケモノのキャラクターのステータスパラメーター(最大HPと攻撃力と防御力)を考えて出力してください。ケモノのキャラクターのステータスパラメーターを考える際は、3枚目の画像のキャラクターの見た目と、配合前のケモノのキャラクターの詳細情報、特にステータスパラメータを参考にしてください。ケモノのキャラクターのステータスパラメータを出力する際は、先ほど与えたフォーマットに従ってください。基本的に、配合ではステータスパラメーターは配合前の2匹のステータスパラメーターの平均よりも強くなるが、そこまで大幅に強くなることもないことに留意してください。ただし、配合前の二匹のケモノの相性がとても良いときは、大幅に強くなることが出来ます。")

	var userContent2 MessageContents
	err = userContent2.AddImage(kemono1.Image)
	if err != nil {
		return
	}
	err = userContent2.AddImage(kemono2.Image)
	if err != nil {
		return
	}
	err = userContent2.AddImage(kemono3.Image)
	if err != nil {
		return
	}
	err = userContent2.AddText(strings.Join(promptTexts, "\n"))
	if err != nil {
		return
	}

	err = messages.AddUserMessageContent(userContent2)
	if err != nil {
		return
	}

	return messages, nil
}
