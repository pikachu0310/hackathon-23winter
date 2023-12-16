package domains

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

type (
	/*
		CREATE TABLE IF NOT EXISTS kemono (
		    id CHAR(36) NOT NULL,
		    image MEDIUMBLOB NOT NULL,
		    prompt TEXT DEFAULT '',
		    concepts TEXT DEFAULT '',
		    name TEXT DEFAULT '',
		    description TEXT DEFAULT '',
		    kind INT DEFAULT -1,
		    color INT DEFAULT -1,
		    is_player BOOLEAN NOT NULL DEFAULT FALSE DEFAULT FALSE,
		    is_for_battle BOOLEAN NOT NULL DEFAULT FALSE DEFAULT FALSE,
		    is_owned BOOLEAN NOT NULL DEFAULT FALSE DEFAULT FALSE,
		    owner_id CHAR(36) DEFAULT '',
		    is_in_field BOOLEAN NOT NULL DEFAULT TRUE DEFAULT TRUE,
		    is_boss BOOLEAN NOT NULL DEFAULT FALSE,
		    field INT DEFAULT -1,
		    x INT DEFAULT -1,
		    y INT DEFAULT -1,
		    has_parent BOOLEAN NOT NULL DEFAULT FALSE,
		    parent1_id CHAR(36) DEFAULT '',
		    parent2_id CHAR(36) DEFAULT '',
		    has_child BOOLEAN NOT NULL DEFAULT FALSE,
		    child_id CHAR(36) DEFAULT '',
		    max_hp INT DEFAULT -1,
		    hp INT DEFAULT -1,
		    attack INT DEFAULT -1,
		    defense INT DEFAULT -1,
		    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		    PRIMARY KEY (id)
		);
	*/

	Kemono struct {
		ID          *uuid.UUID    `db:"id"`
		Image       []byte        `db:"image"`
		Prompt      *string       `db:"prompt"`
		Concepts    *ConceptsText `db:"concepts"`
		Name        *string       `db:"name"`
		Description *string       `db:"description"`
		Kind        *int          `db:"kind"`
		Color       *int          `db:"color"`
		IsPlayer    *bool         `db:"is_player"`
		IsForBattle *bool         `db:"is_for_battle"`
		IsOwned     *bool         `db:"is_owned"`
		OwnerID     *uuid.UUID    `db:"owner_id"`
		IsInField   *bool         `db:"is_in_field"`
		IsBoss      *bool         `db:"is_boss"`
		Field       *int          `db:"field"`
		X           *int          `db:"x"`
		Y           *int          `db:"y"`
		HasParent   *bool         `db:"has_parent"`
		Parent1ID   *uuid.UUID    `db:"parent1_id"`
		Parent2ID   *uuid.UUID    `db:"parent2_id"`
		HasChild    *bool         `db:"has_child"`
		ChildID     *uuid.UUID    `db:"child_id"`
		MaxHp       *int          `db:"max_hp"`
		Hp          *int          `db:"hp"`
		Attack      *int          `db:"attack"`
		Defense     *int          `db:"defense"`
		CreatedAt   *string       `db:"created_at"`
	}

	KemonoParams struct {
		ID          uuid.UUID
		Image       []byte
		Prompt      string
		Concepts    ConceptsText
		Name        string
		Description string
		Kind        int
		Color       int
		IsPlayer    bool
		IsForBattle bool
		IsOwned     bool
		OwnerID     uuid.UUID
		IsInField   bool
		IsBoss      bool
		Field       int
		X           int
		Y           int
		HasParent   bool
		Parent1ID   uuid.UUID
		Parent2ID   uuid.UUID
		HasChild    bool
		ChildID     uuid.UUID
		MaxHp       int
		Hp          int
		Attack      int
		Defense     int
		CreatedAt   string
	}
)

func (kemonoParams *KemonoParams) ToKemono() *Kemono {
	return &Kemono{
		ID:          &kemonoParams.ID,
		Image:       kemonoParams.Image,
		Prompt:      &kemonoParams.Prompt,
		Concepts:    &kemonoParams.Concepts,
		Name:        &kemonoParams.Name,
		Description: &kemonoParams.Description,
		Kind:        &kemonoParams.Kind,
		Color:       &kemonoParams.Color,
		IsPlayer:    &kemonoParams.IsPlayer,
		IsForBattle: &kemonoParams.IsForBattle,
		IsOwned:     &kemonoParams.IsOwned,
		OwnerID:     &kemonoParams.OwnerID,
		IsInField:   &kemonoParams.IsInField,
		IsBoss:      &kemonoParams.IsBoss,
		Field:       &kemonoParams.Field,
		X:           &kemonoParams.X,
		Y:           &kemonoParams.Y,
		HasParent:   &kemonoParams.HasParent,
		Parent1ID:   &kemonoParams.Parent1ID,
		Parent2ID:   &kemonoParams.Parent2ID,
		HasChild:    &kemonoParams.HasChild,
		ChildID:     &kemonoParams.ChildID,
		MaxHp:       &kemonoParams.MaxHp,
		Hp:          &kemonoParams.Hp,
		Attack:      &kemonoParams.Attack,
		Defense:     &kemonoParams.Defense,
		CreatedAt:   &kemonoParams.CreatedAt,
	}
}

type KemonoStatus struct {
	MaxHP   *int
	Attack  *int
	Defence *int
}

func ParseKemonoStatus(s *string) (*KemonoStatus, error) {
	// 改行文字で分割して各行を処理
	lines := strings.Split(*s, "\n")

	// status構造体の初期化
	st := KemonoStatus{}

	// 各行を繰り返し処理
	for _, line := range lines {
		// "="で分割してキーと値を取得
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue // "="が含まれていない行は無視
		}

		// キーに応じて構造体に値を設定
		key := strings.TrimSpace(parts[0])
		value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue // 数値に変換できない行は無視
		}

		switch key {
		case "MaxHP":
			st.MaxHP = &value
		case "Attack":
			st.Attack = &value
		case "Defence":
			st.Defence = &value
		}
	}

	// 全てのキーがセットされたかチェック
	if st.MaxHP == nil || st.Attack == nil || st.Defence == nil {
		return nil, fmt.Errorf("missing status values")
	}

	return &st, nil
}

type KemonoCharacterChip struct {
	Kind  *int
	Color *int
}

func ParseKemonoCharacterChip(s *string) (*KemonoCharacterChip, error) {
	// 改行文字で分割して各行を処理
	lines := strings.Split(*s, "\n")

	// KemonoCharacter構造体の初期化
	kc := KemonoCharacterChip{}

	// 各行を繰り返し処理
	for _, line := range lines {
		// "="で分割してキーと値を取得
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue // "="が含まれていない行は無視
		}

		// キーに応じて構造体に値を設定
		key := strings.TrimSpace(parts[0])
		value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue // 数値に変換できない行は無視
		}

		switch key {
		case "動物":
			kc.Kind = &value
		case "色":
			kc.Color = &value
		}
	}

	// 全てのキーがセットされたかチェック
	if kc.Kind == nil || kc.Color == nil {
		return nil, fmt.Errorf("missing character attributes")
	}

	return &kc, nil
}

func ParseKemonoConcepts(s *string) (*Concepts, error) {
	// 求めるパターンに合う正規表現を定義します
	re := regexp.MustCompile(`- ([^\n]+)`)

	// 正規表現を使ってマッチする部分を全て見つけます
	matches := re.FindAllStringSubmatch(*s, -1)

	var concepts Concepts
	for _, match := range matches {
		if len(match) > 1 {
			// キャプチャしたグループ（マッチしたサブストリング）を追加します
			concepts = append(concepts, match[1])
		}
	}

	return &concepts, nil
}
