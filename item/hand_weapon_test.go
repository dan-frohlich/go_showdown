package item

import (
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
	"testing"
	"showdown/identity"
)

func TestHandWeaponReadYaml(t *testing.T) {
	//t.Skip("skipping test;")

	//log.Printf("running TestRangedWeaponReadYaml")
	src_yml := `
id: 0-00-000-0
name: sword
description: atypical sword
notes: a note
special_cost: 6
cost: 2
AP: 3
damage: Str+d6+1
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

	var expected_id identity.ID = "0-00-000-0"
	actual_id := pi3.ID()
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "id", actual_id)
	}

	expected_s := "sword"
	actual_s := pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "atypical sword"
	actual_s = pi3.Description()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "description", actual_s)
	}

	expected_s = "Str+d6+1"
	actual_s = pi3.Damage().String()
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

	expected = 7
	actual = pi3.Damage().Max()
	if expected != actual {
		t.Errorf(format, expected, "max damage", actual)
	}

	expected = 4
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

func TestRoundTripHandWeapon(t *testing.T) {

	src_yml := `
id: 0-00-000-0
name: sword
description: atypical sword
notes: a note
special_cost: 6
cost: 2
AP: 3
damage: Str+d6+1
parry: 4
reach: 5
`
	hw := &HandWeapon{}
	//log.Printf("unmarshaling a hand weapon...\n")
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), hw)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}
	log.Printf("unmarshaled a hand weapon as %v\n", hw)

	json_bytes, _ := json.Marshal(hw)

	log.Println(string(json_bytes))

	hw2 := &HandWeapon{}
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

	expected = hw.Parry()
	actual = hw2.Parry()
	if expected != actual {
		t.Errorf(format, expected, "parry", actual)
	}

	expected = hw.Reach()
	actual = hw2.Reach()
	if expected != actual {
		t.Errorf(format, expected, "reach", actual)
	}
}
