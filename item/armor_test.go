package item

import (
	"testing"
	"github.com/ghodss/yaml"
)

func TestArmorParse(t *testing.T) {

	src_yml := `
id: 0-00-000-0
name: medium shield
description: descr
notes: a note
special_cost: 1
cost: 5
parry: 1
armor: 2
`

	a := &Armor{}
	//log.Printf("unmarshaling a hand weapon...\n")
	a_umerror := yaml.Unmarshal([]byte(src_yml), a)
	if a_umerror != nil {
		t.Error(a_umerror)
	}
	//log.Printf("unmarshaled a hand weapon as %v\n", a)

	format := "expected %v for %s but found %v"

	expected := 5
	actual := a.Cost()
	if expected != actual {
		t.Errorf(format, expected, "med.shield.cost", actual)
	}

	expected = 1
	actual = a.SpecialCost()
	if expected != actual {
		t.Errorf(format, expected, "med.shield.special.cost", actual)
	}

	expected = 1
	actual = a.Parry()
	if expected != actual {
		t.Errorf(format, expected, "med.shield.parry", actual)
	}

	expected = 2
	actual = a.Armor()
	if expected != actual {
		t.Errorf(format, expected, "med.shield.armor", actual)
	}
}
