package item

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"testing"
)

func TestWeaponReadWrite(t *testing.T) {

	pi := &Weapon{
		Item:    Item{Name_: "pitem", Cost_: 0},
		Damage_: Damage{dice_count: 1, dice_sides: 6, base_stat: "Str"}}
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

	expected := 101
	actual := pi3.Cost()
	if expected != actual {
		t.Errorf(format, expected, "cost", actual)
	}

	expected_s := "0-00-000-0"
	actual_s := pi3.ID()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "id", actual_s)
	}

	expected_s = "pitem"
	actual_s = pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "pidescr"
	actual_s = pi3.Description()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "description", actual_s)
	}

	expected_s = "2d6"
	actual_s = pi3.Damage_.String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "damage", actual_s)
	}

	expected = 111
	actual = pi3.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}

	expected = 12
	actual = pi3.Damage_.Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}
}
