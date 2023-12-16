package domains

import (
	"math/rand"
	"strings"
)

type (
	ConceptsText string
	Concepts     []string
)

func (c ConceptsText) Concepts() Concepts {
	// "a,b,c" -> ["a", "b", "c"]
	return strings.Split(c.String(), ",")
}

func (c ConceptsText) String() string {
	return string(c)
}

func (c Concepts) String() string {
	// ["a", "b", "c"] -> "a,b,c"
	return strings.Join(c, ",")
}

func (c Concepts) Text() *ConceptsText {
	// ["a", "b", "c"] -> "a,b,c"
	conceptsText := ConceptsText(c.String())
	return &conceptsText
}

func (c Concepts) SelectConceptByRandom() string {
	if len(c) == 0 {
		return ""
	}
	return c[rand.Intn(len(c))]
}

func (c *Concepts) Add(concept string) {
	*c = append(*c, concept)
}
