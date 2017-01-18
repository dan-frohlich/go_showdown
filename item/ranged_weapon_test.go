package item

import (
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

	expected_s := "0-00-000-0"
	actual_s := pi3.ID()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "id", actual_s)
	}

	expected_s = "pistol"
	actual_s = pi3.Name()
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
	actual_s = pi3.Damage_.String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "damage", actual_s)
	}

	expected = 2
	actual = pi3.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}

	expected = 12
	actual = pi3.Damage_.Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}

	expected = 12
	actual = pi3.Range().Short()
	if expected != actual {
		t.Errorf(format, expected, "short range", actual)
	}

	expected_s = "12/24/48"
	actual_s = pi3.Range_.String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "range fields", actual_s)
	}

	expected = 3
	actual = pi3.RateOfFire()
	if expected != actual {
		t.Errorf(format, expected, "RoF", actual)
	}

}
