package main

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
	"showdown/item"
	"showdown/character"
	"fmt"
)

func main() {

	//hack to see how items are marshaled in json and yml
	pi := item.Item{Element: item.Element{Name_: "pitem"}, Cost_: 0}
	pij, _ := json.Marshal(pi)
	log.Printf("json pi: %s\n", pij)

	pw := item.Weapon{
		Item:    item.Item{Cost_: 0, Element: item.Element{Name_: "pweap"}},
		Damage_: "2d6"}
	pwj, _ := json.Marshal(pw)
	log.Printf("json pw: %s\n", pwj)

	prw := item.RangedWeapon{
		Weapon: item.Weapon{Damage_: "2d6", Item: item.Item{Cost_: 0, Element: item.Element{Name_: "prngweap"}}},
		Range_: "12/24/48"}
	prwj, _ := json.Marshal(prw)
	log.Printf("json prw: %s\n", prwj)

	piy, _ := yaml.Marshal(pi)
	log.Printf("yml pi:\n%s\n\n", piy)

	pwy, _ := yaml.Marshal(pw)
	log.Printf("yml pw:\n%s\n\n", pwy)

	prwy, _ := yaml.Marshal(prw)
	log.Printf("yml prw:\n%s\n\n", prwy)



	src_yml := `
id: leg.o.lamb.elf.archer
name: Leg O'Lamb
description: Elf archer
notes: snowboarder
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

	a := &character.Character{}
	//log.Printf("unmarshaling a hand weapon...\n")
	a_umerror := yaml.Unmarshal([]byte(src_yml), a)
	if a_umerror != nil {
		log.Println(a_umerror)
	}else{
		string_char, _ := yaml.Marshal(a)
		fmt.Println(string(string_char))
	}

	//l := library.GetLibrary()
	//fmt.Println(l.Attributes)
}


