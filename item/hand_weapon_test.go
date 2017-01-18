package item

import (
	"github.com/ghodss/yaml"
	"testing"
)

func TestHanmdWeaponReadYaml(t *testing.T) {
	//t.Skip("skipping test;")

	//log.Printf("running TestRangedWeaponReadYaml")
	src_yml := `
id: 0-00-000-0
name: pistol
description: atypical pistol
notes: a note
special_cost: 6
cost: 2
AP: 3
damage: 2d6
parry: 4
reach: 5
`

	pi3 := &HandWeapon{}
	//log.Printf("unmarshaling a hand weapon...\n")
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}
	//log.Printf("unmarshaled a hand weapon as %v\n", pi3)

	format := "expected %v for %s but found %v"

	expected := 2
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

	expected_s = "2d6"
	actual_s = pi3.Damage_.String()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "damage", actual_s)
	}

	expected_s = "a note"
	actual_s = pi3.Notes()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "notes", actual_s)
	}

	expected = 3
	actual = pi3.ArmorPiercing()
	if expected != actual {
		t.Errorf(format, expected, "AP", actual)
	}

	expected = 12
	actual = pi3.Damage_.Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}

	expected =4
	actual = pi3.Parry()
	if expected != actual {
		t.Errorf(format, expected, "parry", actual)
	}

	expected = 5
	actual = pi3.Reach()
	if expected != actual {
		t.Errorf(format, expected, "reach", actual)
	}

	expected = 6
	actual = pi3.SpecialCost()
	if expected != actual {
		t.Errorf(format, expected, "special_cost", actual)
	}

}
