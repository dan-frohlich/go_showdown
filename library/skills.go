package library

import (
	"showdown/identity"
	"showdown/trait"
	"github.com/ghodss/yaml"
)


type SkillLibrary map[identity.ID]trait.Skill

var skill_library SkillLibrary = make(SkillLibrary)

func loadSkillLibrary() {

	src := `
showdown.v2009.skill.arcane:
  id: showdown.v2009.skill.arcane
  name: Arcane
  short_name: Arcan
showdown.v2009.skill.fighting:
  id: showdown.v2009.skill.fighting
  name: Fighting
  short_name: Fight
showdown.v2009.skill.intimidate:
  id: showdown.v2009.skill.intimidate
  name: Intimidate
  short_name: Intim
showdown.v2009.skill.shooting:
  id: showdown.v2009.skill.shooting
  name: Shooting
  short_name: Shoot
showdown.v2009.skill.taunt:
  id: showdown.v2009.skill.taunt
  name: Taunt
  short_name: Taunt
showdown.v2009.skill.throwing:
  id: showdown.v2009.skill.throwing
  name: Throwing
  short_name: Throw
`
	if err := yaml.Unmarshal([]byte(src), &skill_library); err != nil {
		panic( err )
	}
}
