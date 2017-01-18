package main

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
	"showdown/item"
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
}
