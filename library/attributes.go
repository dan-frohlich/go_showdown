package library

import (
	"showdown/trait"
	"gopkg.in/yaml.v2"
	"showdown/identity"
	"log"
	"showdown/character"
)

func init() {
	loadAttributeLibrary()
	loadSkillLibrary()
	loadEdgeLibrary()
	log.Printf("loaded %d attributes, %d skills and %d edges.\n",
		len(attribute_library), len(skill_library), len(edge_library))
}

type L struct {
	Attributes map[identity.ID]trait.Attribute
	Edges      map[identity.ID]character.Edge
	Skills     map[identity.ID]trait.Skill
}

func GetLibrary() L {
	return L{Edges:edge_library, Attributes:attribute_library, Skills:skill_library}
}

type AttributeLibrary map[identity.ID]trait.Attribute

var attribute_library AttributeLibrary = make(AttributeLibrary)

func loadAttributeLibrary() {
	src := `
showdown.v2009.attribute.agility:
  id: showdown.v2009.attribute.agility
  name: Agility
  short_name: Agi
showdown.v2009.attribute.spirit:
  id: showdown.v2009.attribute.spirit
  name: Spirit
  short_name: Spi
showdown.v2009.attribute.smarts:
  id: showdown.v2009.attribute.smarts
  name: Smarts
  short_name: Sma
showdown.v2009.attribute.strength:
  id: showdown.v2009.attribute.strength
  name: Strength
  short_name: Str
showdown.v2009.attribute.vigor:
  id: showdown.v2009.attribute.vigor
  name: Vigor
  short_name: Vig
`
	if err := yaml.Unmarshal([]byte(src), &attribute_library); err != nil {
		panic(err)
	}
}
