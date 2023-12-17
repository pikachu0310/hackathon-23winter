package domains

import (
	"math/rand"
)

type FieldType int

const (
	Forest FieldType = iota
	Grassland
	Water
	Snow
	Ice
	HotPlace
	Volcano
	Lava
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
	case 2:
		r := rand.Intn(3)
		if r == 0 {
			return Water
		}
		if r == 1 {
			return Snow
		}
		return Ice
	case 3:
		r := rand.Intn(3)
		if r == 0 {
			return HotPlace
		}
		if r == 1 {
			return Volcano
		}
		return Lava
	}
	return Grassland
}

func (filed FieldType) String() string {
	switch filed {
	case Forest:
		return "森"
	case Grassland:
		return "草原"
	case Water:
		return "水"
	case Snow:
		return "雪"
	case Ice:
		return "氷"
	case HotPlace:
		return "熱帯"
	case Volcano:
		return "火山"
	case Lava:
		return "溶岩"
	}
	return ""
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
	case Water:
		return Concepts{
			"水辺に生息",
			"湖に生息",
		}
	case Snow:
		return Concepts{
			"雪原に生息",
		}
	case Ice:
		return Concepts{
			"氷原に生息",
		}
	case HotPlace:
		return Concepts{
			"熱帯に生息",
		}
	case Volcano:
		return Concepts{
			"火山に生息",
		}
	case Lava:
		return Concepts{
			"溶岩に生息",
		}
	}
	return nil
}

func FieldTypeToConcepts2(fieldType FieldType) Concepts {
	switch fieldType {
	case Forest:
		return Concepts{
			"四足歩行",
			"活発に動き回る",
			"マスコット",
			"とてもかわいい",
			"愛くるしい",
		}
	case Grassland:
		return Concepts{
			"四足歩行",
			"活発に動き回る",
			"二足歩行",
			"マスコット",
			"とてもかわいい",
			"愛くるしい",
		}
	case Water:
		return Concepts{
			"魚",
			"水棲",
			"水中生物",
			"水を操る",
			"マスコット",
			"とてもかわいい",
		}
	case Snow:
		return Concepts{
			"厚い毛皮",
			"雪を操る",
			"暖かい見た目",
			"寒さに強い",
		}
	case Ice:
		return Concepts{
			"寒さに凄く強い",
			"暖かい見た目",
			"氷を操る",
			"ドラゴン",
			"強い",
		}
	case HotPlace:
		return Concepts{
			"暑さに強い",
			"熱風を扱う",
			"マスコット",
		}
	case Volcano:
		return Concepts{
			"炎を操る",
			"ドラゴン",
			"かっこよい",
		}
	case Lava:
		return Concepts{
			"溶岩を操る",
			"ドラゴン",
			"強い",
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
	case Water:
		return 200 + rand.Intn(23), 0 + rand.Intn(13)
	case Snow:
		return 200 + rand.Intn(23), 13 + rand.Intn(17)
	case Ice:
		return 223 + rand.Intn(17), 13 + rand.Intn(17)
	case HotPlace:
		return 300 + rand.Intn(30), 0 + rand.Intn(14)
	case Volcano:
		return 300 + rand.Intn(30), 14 + rand.Intn(16)
	case Lava:
		return 318 + rand.Intn(12), 21 + rand.Intn(9)
	default:
		return 100 + rand.Intn(40), 0 + rand.Intn(15)
	}
}
