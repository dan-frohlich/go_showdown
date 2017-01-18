package character

import (
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"log"
	"showdown/item"
	"testing"
)

func TestCharacterParse(t *testing.T) {

	src_yml := `
id: leg.o.lamb.elf.archer
name: Leg O'Lamb
description: Elf archer
notes: snowboarder
attributes:
  showdown.v2009.arrtibute.agility:
    id: showdown.v2009.arrtibute.agility
    name: Agility
    short_name: Agi
    level: 12
  showdown.v2009.arrtibute.spirit:
    id: showdown.v2009.arrtibute.spirit
    name: Spirit
    short_name: Spi
    level: 8
  showdown.v2009.arrtibute.smarts:
    id: showdown.v2009.arrtibute.smarts
    name: Smarts
    short_name: Sma
    level: 8
  showdown.v2009.arrtibute.strength:
    id: showdown.v2009.arrtibute.strength
    name: Strength
    short_name: Str
    level: 6
  showdown.v2009.arrtibute.vigor:
    id: showdown.v2009.arrtibute.vigor
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
	var knife_id item.ID = "showdown.v2009.hand.knife"
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

	var shooting_id item.ID = "showdown.v2009.skill.shooting"
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
