package character

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"log"
	"testing"
	"showdown/identity"
)

func TestCharacterParse(t *testing.T) {

	src_yml := `
id: showdown.v2009.character.leg.o.lamb.elf.archer
name: Leg O'Lamb
description: Elf archer
notes: snowboarder
edges:
  showdown.v2009.edge.quick:
    id: showdown.v2009.edge.quick
    name: Quick
    description: >-
      The unit cannot have an action card lower than 5
      (discard and redeal until this is true)
hinderances:
powers:
  showdown.v2009.power.quickness:
    id: showdown.v2009.power.quickness
    name: quickness
    description: >-
      Quickness grants a figure incredible celerity and clarity of thought.
      With success, the target has two actions per round instead of the usual one
      (at no multi- action penalty).
      With a raise, the recipient gains the bene ts above, and can also discard
      and redraw any initiative cards lower than 8 each round.
attributes:
  showdown.v2009.attribute.agility:
    id: showdown.v2009.attribute.agility
    name: Agility
    short_name: Agi
    level: 12
  showdown.v2009.attribute.spirit:
    id: showdown.v2009.attribute.spirit
    name: Spirit
    short_name: Spi
    level: 8
  showdown.v2009.attribute.smarts:
    id: showdown.v2009.attribute.smarts
    name: Smarts
    short_name: Sma
    level: 8
  showdown.v2009.attribute.strength:
    id: showdown.v2009.attribute.strength
    name: Strength
    short_name: Str
    level: 6
  showdown.v2009.attribute.vigor:
    id: showdown.v2009.attribute.vigor
    name: Vigor
    short_name: Vig
    level: 8
skills:
  showdown.v2009.skill.fighting:
    id: showdown.v2009.skill.fighting
    name: Fighting
    short_name: Fight
    level: 10
  showdown.v2009.skill.shooting:
    id: showdown.v2009.skill.shooting
    name: Shooting
    short_name: Shoot
    level: 12
armor:
  showdown.v2009.armor.leather:
    id: showdown.v2009.armor.leather
    name: Leather Armor
    short_name: Leather
    armor: 1
    cost: 1
ranged_weapons:
  showdown.v2009.ranged.english-long-bow:
    id: showdown.v2009.ranged.english-long-bow
    name: English Long Bow
    short_name: Lng Bw
    damage: 2d6
    range: 15/30/60
    rof: 1
    cost: 7
    count: 1
hand_weapons:
  showdown.v2009.hand.knife:
    id: showdown.v2009.hand.knife
    name: Knife
    short_name: Knife
    damage: Str+d4
    cost: 3
    count: 2
`

	a := &Character{}
	//log.Printf("unmarshaling a hand weapon...\n")
	a_umerror := yaml.Unmarshal([]byte(src_yml), a)
	if a_umerror != nil {
		t.Error(a_umerror)
	}

	log.Printf("unmarshaled a character as %v\n", a)

	format := "expected_s %v for %bytes but found %v"

	expected_s := "Leg O'Lamb"
	actual_s := a.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "char.name", actual_s)
	}
	var knife_id identity.ID = "showdown.v2009.hand.knife"
	actual := a.HandWeapons_[knife_id].Count_
	expected := 2
	if expected != actual {
		t.Errorf(format, expected, "char.knife.count", actual)
	}

	actual = a.HandWeapons_[knife_id].Damage().Max()
	expected = 4
	if expected != actual {
		t.Errorf(format, expected, "char.knife.max.damage", actual)
	}

	fmt.Println("remarshaling char...")
	bytes, _ := json.Marshal(a)
	fmt.Println(string(bytes))

	var shooting_id identity.ID = "showdown.v2009.skill.shooting"
	shooting_skill := a.Skills_[shooting_id]
	actual_id := shooting_skill.ID()
	expected_id := shooting_id
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "shooting id", actual_id)
	}

	actual = shooting_skill.Level()
	expected = 12
	if expected != actual {
		t.Errorf(format, expected, "shooting level", actual)
	}
}
