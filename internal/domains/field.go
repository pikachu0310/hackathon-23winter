package domains

import (
	"math/rand"
)

type FieldType int

const (
	Forest FieldType = iota
	Grassland
)

func FieldIdToFieldType(fieldId int) FieldType {
	switch fieldId {
	case 1:
		r := rand.Intn(2)
		if r == 0 {
			return Forest
		} else {
			return Grassland
		}
	}
	return Grassland
}

func FieldTypeToConcepts1(fieldType FieldType) Concepts {
	switch fieldType {
	case Forest:
		return Concepts{
			"平原に生息",
			"お花畑に生息",
		}
	case Grassland:
		return Concepts{
			"森に生息",
			"森林に生息",
		}
	}
	return nil
}

func FieldTypeToConcepts2(fieldType FieldType) Concepts {
	switch fieldType {
	case Forest:
		return Concepts{
			"自然と調和する",
			"花の精霊",
			"緑の中を跳ねる",
			"植物と会話する",
			"太陽の恵みを受ける",
			"ふわもこ質感",
			"ふんわり毛皮",
			"パステルカラー",
			"蜜愛好家",
			"四足歩行",
		}
	case Grassland:
		return Concepts{
			"自然と調和する",
			"花の精霊",
			"緑の中を跳ねる",
			"植物と会話する",
			"太陽の恵みを受ける",
			"ふわもこ質感",
			"ふんわり毛皮",
			"パステルカラー",
			"蜜愛好家",
			"四足歩行",
		}
	}
	return nil
}

func FieldTypeToPosition(fieldType FieldType) (x int, y int) {
	switch fieldType {
	case Forest:
		return 100 + rand.Intn(40), 0 + rand.Intn(15)
	case Grassland:
		return 100 + rand.Intn(40), 15 + rand.Intn(15)
	default:
		return 100 + rand.Intn(40), 0 + rand.Intn(30)
	}
}
