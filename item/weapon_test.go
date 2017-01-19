package item

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
	"testing"
	"showdown/identity"
)

func TestWeaponReadWrite(t *testing.T) {

	pi := &Weapon{
		Item:    Item{Element: Element{Name_: "pitem"}, Cost_: 0},
		Damage_: "Str+d6"}
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

	var expected_id identity.ID = "0-00-000-0"
	actual_id := pi3.ID()
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "id", actual_id)
	}

	expected_s := "pitem"
	actual_s := pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "pidescr"
	actual_s = pi3.Description()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "description", actual_s)
	}

	expected_s = "2d6"
	actual_s = pi3.Damage_
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "damage", actual_s)
	}

	expected = 111
	actual = pi3.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}

	expected = 12
	actual = pi3.Damage().Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}
}

func TestRoundTripWeapon(t *testing.T) {

	src_yml := `
id: 0-00-000-0
name: sword
description: atypical sword
notes: a note
special_cost: 6
cost: 2
AP: 3
damage: Str+d6+1
`
	hw := &Weapon{}
	//log.Printf("unmarshaling a hand weapon...\n")
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), hw)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}
	log.Printf("unmarshaled a hand weapon as %v\n", hw)

	json_bytes, _ := json.Marshal(hw)

	log.Println(string(json_bytes))

	hw2 := &Weapon{}
	json.Unmarshal(json_bytes, hw2)

	format := "expected %v for %s but found %v"

	expected := hw.Cost()
	actual := hw2.Cost()
	if expected != actual {
		t.Errorf(format, expected, "cost", actual)
	}

	expected_s := hw.Damage().String()
	actual_s := hw2.Damage().String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "max damage", actual_s)
	}

	expected = hw.Damage().Max()
	actual = hw2.Damage().Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}

	expected = hw.ArmorPiercing()
	actual = hw2.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}
}
