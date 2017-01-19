package character

import (
	"showdown/item"
	"showdown/trait"
	"showdown/identity"
)


type Character struct {
	item.Element
	Attributes_    map[identity.ID]*trait.Attribute   `json:"attributes" yaml:"attributes"`
	Skills_        map[identity.ID]*trait.Skill       `json:"skills" yaml:"skills"`
	Edges_         map[identity.ID]*Edge              `json:"edges" yaml:"edges"`
	Hindrances_    map[identity.ID]*Hindrance         `json:"hinderances" yaml:"hinderances"`
	Powers_        map[identity.ID]*Power             `json:"powers" yaml:"powers"`
	Armor_         map[identity.ID]*item.Armor        `json:"armor" yaml:"armor"`
	RangedWeapons_ map[identity.ID]*item.RangedWeapon `json:"ranged_weapons" yaml:"ranged_weapons"`
	HandWeapons_   map[identity.ID]*item.HandWeapon   `json:"hand_weapons" yaml:"hand_weapons"`
}
