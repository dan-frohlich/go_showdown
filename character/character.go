package character

import (
	"showdown/item"
	"showdown/trait"
)

type Character struct {
	item.Element
	Attributes_    map[item.ID]*trait.Attribute   `json:"attributes" yaml:"attributes"`
	Skills_        map[item.ID]*trait.Skill       `json:"skills" yaml:"skills"`
	Armor_         map[item.ID]*item.Armor        `json:"armor" yaml:"armor"`
	RangedWeapons_ map[item.ID]*item.RangedWeapon `json:"ranged_weapons" yaml:"ranged_weapons"`
	HandWeapons_   map[item.ID]*item.HandWeapon   `json:"hand_weapons" yaml:"hand_weapons"`
}
