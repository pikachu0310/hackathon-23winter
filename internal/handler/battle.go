package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pikachu0310/hackathon-23winter/internal/repository"
	"github.com/pikachu0310/hackathon-23winter/internal/repository/api"
	"net/http"
	"strconv"
	"strings"
)

type (
	/*
		CREATE TABLE IF NOT EXISTS battle (
		    id CHAR(36) NOT NULL,
		    my_kemono_id CHAR(36) NOT NULL,
		    enemy_kemono_id CHAR(36) NOT NULL,
		    text TEXT DEFAULT '',
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	*/

	PostBattleRequest struct {
		MyKemonoID    uuid.UUID `json:"my_kemono_id"`
		EnemyKemonoID uuid.UUID `json:"enemy_kemono_id"`
	}

	PostBattleResponse struct {
		BattleId uuid.UUID `json:"battle_id"`
	}

	PostBattleDamageRequest struct {
		Damage int `json:"damage"`
	}

	PostBattleDamageResponse struct {
		Text string `json:"text"`
	}
)

// GET /api/v1/battles
func (h *Handler) GetBattles(c echo.Context) error {
	battles, err := h.repo.GetBattles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, battles)
}

// GET /api/v1/battles/:battle_id
func (h *Handler) GetBattle(c echo.Context) error {
	battleId, err := uuid.Parse(c.Param("battle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid battle id").SetInternal(err)
	}

	battle, err := h.repo.GetBattle(battleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.JSON(http.StatusOK, battle)
}

// POST /api/v1/battles
func (h *Handler) PostBattle(c echo.Context) error {
	req := new(PostBattleRequest)
	//if err := c.Bind(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}
	myKemonoID, err := uuid.Parse(c.FormValue("my_kemono_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid my kemono id").SetInternal(err)
	}
	enemyKemonoID, err := uuid.Parse(c.FormValue("enemy_kemono_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid enemy kemono id").SetInternal(err)
	}
	req.MyKemonoID = myKemonoID
	req.EnemyKemonoID = enemyKemonoID
	if req.MyKemonoID == uuid.Nil || req.EnemyKemonoID == uuid.Nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	//if err := c.Validate(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}

	myKemono, err := h.repo.GetKemono(c.Request().Context(), req.MyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	enemyKemono, err := h.repo.GetKemono(c.Request().Context(), req.EnemyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	battleId, err := h.repo.CreateBattle(req.MyKemonoID, req.EnemyKemonoID, DefaultBattleText(*myKemono, *enemyKemono))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := PostBattleResponse{
		BattleId: battleId,
	}

	return c.JSON(http.StatusOK, res)
}

// POST /api/v1/battles/:battle_id
func (h *Handler) PostBattleDamage(c echo.Context) error {
	battleId, err := uuid.Parse(c.Param("battle_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid battle id").SetInternal(err)
	}

	req := new(PostBattleDamageRequest)
	//if err := c.Bind(req); err != nil {
	//	return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	//}
	damage, err := strconv.Atoi(c.FormValue("damage"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid damage").SetInternal(err)
	}
	req.Damage = damage

	battle, err := h.repo.GetBattle(battleId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	myKemono, err := h.repo.GetKemono(c.Request().Context(), *battle.MyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}
	enemyKemono, err := h.repo.GetKemono(c.Request().Context(), *battle.EnemyKemonoID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	prompt := GenerateBattlePromptText(*myKemono, *enemyKemono, req.Damage, *battle.Text)
	text, err := api.GetGeneratedText(fmt.Sprintf(prompt))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	err = h.repo.UpdateBattleText(battleId, *battle.Text+"\n"+*text)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	res := PostBattleDamageResponse{
		Text: *text,
	}

	return c.JSON(http.StatusOK, res)
}

func DefaultBattleText(attacker repository.Kemono, defender repository.Kemono) string {
	return fmt.Sprintf("%sは", *attacker.Name)
}

func GenerateBattlePromptText(attacker repository.Kemono, defender repository.Kemono, damage int, battleText string) string {
	var prompt []string
	/*
		この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。
		いま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。

		以下にケモノ「アクアフローラ」の特徴を載せます。
		名前:「アクアフローラ」
		能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」
		そのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」
		最大体力:「100」
		現在の体力:「100」
		攻撃力:「10」
		防御力:「10」

		以下にケモノ「ヒバナ」の特徴を載せます。
		名前:「ヒバナ」
		能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」
		そのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」
		最大体力:「20」
		現在の体力:「20」
		攻撃力:「5」
		防御力:「3」

		これからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「8」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「8」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。

		---

		アクアフローラは穏やかな目でヒバナを見つめ、水の魔法を準備します。「森を守るのは私の使命。私の水は、君の炎を鎮めることができる。」

		ヒバナ、そのふわふわの耳を揺らしながら、炎を纏い反撃の構えをとります。「森の守り神として、私はこの地を護る。あなたの水など、私の炎には及ばない！」

		戦いが始まると、アクアフローラは軽やかに跳ねながら、水の魔法を発動。鮮やかな水流がヒバナに向かって躍動します。

		ヒバナは身軽に避けようとしますが、アクアフローラの水流は素早く、ヒバナの体を捉えて「8」のダメージを与えます。水の衝撃にヒバナは少し後退し、驚きの表情を浮かべます。

		「こんなにも強い水の力…だが、私はまだ負けない！」ヒバナは決意を新たにし、再び炎を燃やし始めます。アクアフローラもそれに応えるように、次の一撃を準備し、バトルは更に激しく続いていきます。

		------

		この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。
		いま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。

		以下にケモノ「アクアフローラ」の特徴を載せます。
		名前:「アクアフローラ」
		能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」
		そのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」
		最大体力:「100」
		現在の体力:「100」
		攻撃力:「10」
		防御力:「10」

		以下にケモノ「ヒバナ」の特徴を載せます。
		名前:「ヒバナ」
		能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」
		そのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」
		最大体力:「20」
		現在の体力:「20」
		攻撃力:「5」
		防御力:「3」

		これからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「12」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「12」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。

		---

		アクアフローラは、ヒバナに向かって穏やかながらも力強い声で語りかけます。「私の水は、森の平和を守るためのもの。君の炎もまた、森を護る力。しかし、今は争いの時だ。」

		ヒバナは激しく炎を揺らし、アクアフローラに対して挑戦的に応じます。「私の炎は、森を守るために燃える！君の水がそれを消すことはできない！」

		その瞬間、アクアフローラは鮮やかな水の魔法を解き放ち、ヒバナに向けて大波を発生させます。水の壁がヒバナに迫る中、ヒバナは必死に抵抗しようとしますが、アクアフローラの水流は圧倒的で、ヒバナは大きなダメージを受けてしまいます。

		ヒバナは「12」のダメージを受け、一時的に立ちすくんでしまいますが、すぐに元気を取り戻し、アクアフローラに対して強い決意を見せつけます。「このダメージ…しかし、私はまだ戦える！私の炎は、こんなことでは消えない！」

		アクアフローラもまた、ヒバナの強い意志を感じ取り、次の攻撃を準備します。「君の強さはわかっている。だが、私も負けられない。この戦い、私たちは全力を尽くそう。」

		------

	*/
	/*
		この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。
		いま、「%s」という名前のケモノと「%s」という名前のケモノがバトルをしています。

		以下にケモノ「%s」の特徴を載せます。
		名前:「%s」
		能力や特徴:「%s」
		そのケモノを表す要素:「%s」
		最大体力:「%d」
		現在の体力:「%d」
		攻撃力:「%d」
		防御力:「%d」

		以下にケモノ「%s」の特徴を載せます。
		名前:「%s」
		能力や特徴:「%s」
		そのケモノを表す要素:「%s」
		最大体力:「%d」
		現在の体力:「%d」
		攻撃力:「%d」
		防御力:「%d」

		これからこの物語では、ケモノ「%s」は、ケモノ「%s」に「%d」のダメージを与えるので、その戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。

		---

		%sは、
	*/
	prompt = append(prompt, fmt.Sprintf("この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。\nいま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。\n\n以下にケモノ「アクアフローラ」の特徴を載せます。\n名前:「アクアフローラ」\n能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「100」\n現在の体力:「100」\n攻撃力:「10」\n防御力:「10」\n\n以下にケモノ「ヒバナ」の特徴を載せます。\n名前:「ヒバナ」\n能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「20」\n現在の体力:「20」\n攻撃力:「5」\n防御力:「3」\n\nこれからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「8」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「8」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。\n\n---\n\nアクアフローラは穏やかな目でヒバナを見つめ、水の魔法を準備します。「森を守るのは私の使命。私の水は、君の炎を鎮めることができる。」\n\nヒバナ、そのふわふわの耳を揺らしながら、炎を纏い反撃の構えをとります。「森の守り神として、私はこの地を護る。あなたの水など、私の炎には及ばない！」\n\n戦いが始まると、アクアフローラは軽やかに跳ねながら、水の魔法を発動。鮮やかな水流がヒバナに向かって躍動します。\n\nヒバナは身軽に避けようとしますが、アクアフローラの水流は素早く、ヒバナの体を捉えて「8」のダメージを与えます。水の衝撃にヒバナは少し後退し、驚きの表情を浮かべます。\n\n「こんなにも強い水の力…だが、私はまだ負けない！」ヒバナは決意を新たにし、再び炎を燃やし始めます。アクアフローラもそれに応えるように、次の一撃を準備し、バトルは更に激しく続いていきます。\n\n------\n\nこの物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。\nいま、「アクアフローラ」という名前のケモノと「ヒバナ」という名前のケモノがバトルをしています。\n\n以下にケモノ「アクアフローラ」の特徴を載せます。\n名前:「アクアフローラ」\n能力や特徴:「このキャラクターは、やさしい目とふわふわの尾を持つ水属性の森の精霊です。生まれながらにして森を潤し、清らかな水を操る能力を持つ。その鮮やかな色合いは森の生命力を象徴し、可愛らしい外見にもかかわらず、敵には強力な水の魔法で立ち向かう勇敢さを秘めている。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, similar to a Pokemon style. The design should be vividly colored and embody the water element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「100」\n現在の体力:「100」\n攻撃力:「10」\n防御力:「10」\n\n以下にケモノ「ヒバナ」の特徴を載せます。\n名前:「ヒバナ」\n能力や特徴:「このキャラクターは、炎属性を持つ森の守り神であり、その生き生きとしたオレンジと赤の色合いが情熱と元気を象徴しています。ふわふわの耳と尾はその愛らしさを際立たせる一方で、目には決意と勇気が宿っており、戦いのときには強い炎を操る力を秘めています。このキャラクターは森を通り抜ける冒険者にとって、時には可愛らしいガイドとなり、時には炎の魔法で道を阻む挑戦者となります。その愛くるしい外見に騙されてはならず、彼の炎の力は敵を一瞬にして灰に変えるほど強力です。」\nそのケモノを表す要素:「Create a single, cute, four-legged kemomimi (animal-eared) character for a game. This character should have the appearance of waking up in a forest, inspired by Pokemon-style creatures. The design should be brightly colored and embody the fire element, ensuring it fits as a potential enemy in the game without feeling out of place. The character should be immediately usable in a game, designed without any specific color palettes or markers used for design purposes.」\n最大体力:「20」\n現在の体力:「20」\n攻撃力:「5」\n防御力:「3」\n\nこれからこの物語では、これから「アクアフローラ」が、「ヒバナ」に「12」のダメージを与えるので、「アクアフローラ」が「ヒバナ」に「12」のダメージを与える瞬間の戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。\n\n---\n\nアクアフローラは、ヒバナに向かって穏やかながらも力強い声で語りかけます。「私の水は、森の平和を守るためのもの。君の炎もまた、森を護る力。しかし、今は争いの時だ。」\n\nヒバナは激しく炎を揺らし、アクアフローラに対して挑戦的に応じます。「私の炎は、森を守るために燃える！君の水がそれを消すことはできない！」\n\nその瞬間、アクアフローラは鮮やかな水の魔法を解き放ち、ヒバナに向けて大波を発生させます。水の壁がヒバナに迫る中、ヒバナは必死に抵抗しようとしますが、アクアフローラの水流は圧倒的で、ヒバナは大きなダメージを受けてしまいます。\n\nヒバナは「12」のダメージを受け、一時的に立ちすくんでしまいますが、すぐに元気を取り戻し、アクアフローラに対して強い決意を見せつけます。「このダメージ…しかし、私はまだ戦える！私の炎は、こんなことでは消えない！」\n\nアクアフローラもまた、ヒバナの強い意志を感じ取り、次の攻撃を準備します。「君の強さはわかっている。だが、私も負けられない。この戦い、私たちは全力を尽くそう。」\n\n------\n"))
	prompt = append(prompt, fmt.Sprintf("この物語は、ゲームの世界でケモノ達がバトルを繰り広げる世界の中で、二匹のケモノがバトルをしているところを描いた物語です。"))
	prompt = append(prompt, fmt.Sprintf("いま、「%s」という名前のケモノと「%s」という名前のケモノがバトルをしています。", *attacker.Name, *defender.Name))
	prompt = append(prompt, fmt.Sprintf(""))
	prompt = append(prompt, fmt.Sprintf("以下にケモノ「%s」の特徴を載せます。\n名前:「%s」\n能力や特徴:「%s」\nそのケモノを表す要素:「%s」\n最大体力:「%d」\n現在の体力:「%d」\n攻撃力:「%d」\n防御力:「%d」", *attacker.Name, *attacker.Name, *attacker.Description, *attacker.Prompt, *attacker.MaxHp, *attacker.Hp, *attacker.Attack, *attacker.Defense))
	prompt = append(prompt, fmt.Sprintf(""))
	prompt = append(prompt, fmt.Sprintf("以下にケモノ「%s」の特徴を載せます。\n名前:「%s」\n能力や特徴:「%s」\nそのケモノを表す要素:「%s」\n最大体力:「%d」\n現在の体力:「%d」\n攻撃力:「%d」\n防御力:「%d」", *defender.Name, *defender.Name, *defender.Description, *defender.Prompt, *defender.MaxHp, *defender.Hp, *defender.Attack, *defender.Defense))
	prompt = append(prompt, fmt.Sprintf(""))
	prompt = append(prompt, fmt.Sprintf("これからこの物語では、ケモノ「%s」は、ケモノ「%s」に「%d」のダメージを与えるので、その戦闘シーンをバトル小説のようにお互いが会話をしながらダメージを与える/受ける描写を書いていきます。", *attacker.Name, *defender.Name, damage))
	prompt = append(prompt, fmt.Sprintf(""))
	prompt = append(prompt, fmt.Sprintf("---"))
	prompt = append(prompt, fmt.Sprintf(""))
	prompt = append(prompt, fmt.Sprintf("%s", battleText))
	return strings.Join(prompt, "\n")
}

// GET /api/v1/reset/battles
func (h *Handler) ResetBattles(c echo.Context) error {
	err := h.repo.ResetBattles()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error()).SetInternal(err)
	}

	return c.NoContent(http.StatusOK)
}
