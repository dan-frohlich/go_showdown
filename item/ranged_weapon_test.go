package item

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"testing"
)

func TestRangedWeaponReadYaml(t *testing.T) {
	//t.Skip("skipping test;")

	//log.Printf("running TestRangedWeaponReadYaml")
	src_yml := `
id: 0-00-000-0
name: pistol
description: atypical pistol
notes: a note
special_cost: 6
cost: 1
AP: 2
damage: 2d6
range: 12/24/48
RoF: 3
`

	pi3 := &RangedWeapon{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	format := "expected %v for %s but found %v"

	expected := 1
	actual := pi3.Cost()
	if expected != actual {
		t.Errorf(format, expected, "cost", actual)
	}

	var expected_id ID = "0-00-000-0"
	actual_id := pi3.ID()
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "id", actual_id)
	}

	expected_s := "pistol"
	actual_s := pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "atypical pistol"
	actual_s = pi3.Description()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "description", actual_s)
	}

	expected_s = "a note"
	actual_s = pi3.Notes()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "notes", actual_s)
	}

	expected = 6
	actual = pi3.SpecialCost()
	if expected != actual {
		t.Errorf(format, expected, "special_cost", actual)
	}

	expected_s = "2d6"
	actual_s = pi3.Damage().String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "damage", actual_s)
	}

	expected = 2
	actual = pi3.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}

	expected = 12
	actual = pi3.Damage().Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}

	expected = 12
	actual = pi3.Range().Short()
	if expected != actual {
		t.Errorf(format, expected, "short range", actual)
	}

	expected_s = "12/24/48"
	actual_s = pi3.Range().String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "range fields", actual_s)
	}

	expected = 3
	actual = pi3.RateOfFire()
	if expected != actual {
		t.Errorf(format, expected, "RoF", actual)
	}

}

func TestRoundTripRangedWeapon(t *testing.T) {

	src_yml := `
id: 0-00-000-0
name: sword
description: atypical sword
notes: a note
special_cost: 6
cost: 2
AP: 3
damage: Str+d6+1
range: 12/24/48
rof: 1
`
	hw := &RangedWeapon{}
	//log.Printf("unmarshaling a hand weapon...\n")
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), hw)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}
	//log.Printf("unmarshaled a hand weapon as %v\n", hw)

	json_str, _ := json.Marshal(hw)
	hw2 := &RangedWeapon{}
	json.Unmarshal(json_str, hw2)

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

	expected = hw.Range().Short()
	actual = hw2.Range().Short()
	if expected != actual {
		t.Errorf(format, expected, "range short", actual)
	}

	expected = hw.RateOfFire()
	actual = hw2.RateOfFire()
	if expected != actual {
		t.Errorf(format, expected, "RoF", actual)
	}
}
