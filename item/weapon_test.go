package item

import (
	"testing"
	"encoding/json"
	"github.com/ghodss/yaml"
)

func TestWeaponReadWrite(t *testing.T) {

	pi := &Weapon{
		Item:   Item{Name_:  "pitem", Cost_:0},
		Damage_:Damage{dice_count:1, dice_sides:6, base_stat:"Str"} }
	pij, imerr := json.Marshal(pi)

	if imerr != nil {
		t.Error(imerr)
	}

	pi2 := &Weapon{}
	iumerror := json.Unmarshal(pij, pi2)

	if iumerror != nil {
		t.Error(iumerror)
	}

	if pi.Name_ != pi2.Name_ {
		t.Errorf("got %s when expecting %s", pi2.Name_, pi.Name_)
	}

}

func TestWeaponReadYaml(t *testing.T) {
	src_yml := `
id: 0-00-000-0
name: pitem
description: pidescr
cost: 101
damage: 2d6
AP: 111
`

	pi3 := &Weapon{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	format := "expected %v for %s but found %v"
	if 101 != pi3.Cost() {
		t.Errorf(format, 101, "cost", pi3.Cost())
	}

	if "0-00-000-0" != pi3.ID() {
		t.Errorf(format, "0-00-000-0", "id", pi3.ID())
	}

	if "pitem" != pi3.Name() {
		t.Errorf(format, "pitem", "name", pi3.Name())
	}

	if "pidescr" != pi3.Description() {
		t.Errorf(format, "pidescr", "description", pi3.Description())
	}

	if "2d6" != pi3.Damage_.String() {
		t.Errorf(format, "2d6", "damage", pi3.Damage_.String())
	}

	if 111 != pi3.ArmorPiercing() {
		t.Errorf(format, 111, "AP", pi3.ArmorPiercing())
	}

	if 12 != pi3.Damage_.Max() {
		t.Errorf(format, 12, "max damage", pi3.Damage_.Max())
	}
}
