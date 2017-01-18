package main

import (
	"showdown/item"
	"log"
	"encoding/json"
	"github.com/ghodss/yaml"
)

func main() {

	//hack to see how items are marshaled in json and yml
	pi := item.Item{Name_:"pitem", Cost_:0}
	pij, _ := json.Marshal(pi)
	log.Printf("json pi: %s\n", pij)

	d, _ := item.ParseDamage("2d6")
	pw := item.Weapon{
		Item:item.Item{Cost_:0, Name_:"pweap"},
		Damage_:d}
	pwj, _ := json.Marshal(pw)
	log.Printf("json pw: %s\n", pwj)

	r12, _ := item.ParseRangeBand("12/24/48")

	prw := item.RangedWeapon{
		Weapon:item.Weapon{Damage_:d,
			Item:item.Item{Cost_:0, Name_:"prngweap"},
		},
		Range_:r12}
	prwj, _ := json.Marshal(prw)
	log.Printf("json prw: %s\n", prwj)

	piy, _ := yaml.Marshal(pi)
	log.Printf("yml pi:\n%s\n\n", piy)

	pwy, _ := yaml.Marshal(pw)
	log.Printf("yml pw:\n%s\n\n", pwy)

	prwy, _ := yaml.Marshal(prw)
	log.Printf("yml prw:\n%s\n\n", prwy)
}