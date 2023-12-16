package domains

import "strings"

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
