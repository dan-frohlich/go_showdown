package library

import (
	"showdown/identity"
	"showdown/character"
	"github.com/ghodss/yaml"
)

type EdgeLibrary map[identity.ID]character.Edge

var edge_library EdgeLibrary = make(EdgeLibrary)

func loadEdgeLibrary() {
	src := `
showdown.2009.edge.Ace:
  id:showdown.2009.edge.Ace
  name:Ace
showdown.2009.edge.Acrobat:
  id:showdown.2009.edge.Acrobat
  name:Acrobat
`

	if err := yaml.Unmarshal([]byte(src), &edge_library); err != nil {
		panic(err)
	}
}
