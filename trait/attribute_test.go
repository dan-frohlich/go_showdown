package trait

import (
	"github.com/ghodss/yaml"
	"log"
	"showdown/item"
	"testing"
)

func TestAttributeParse(t *testing.T) {
	str_src_yml := `
id: core.deluxe.attribute.strength
name: Strength
short_name: Str
description: Might
notes: physical strnegth
level: 1
`

	pi3 := &Attribute{}
	pi3_umerror := yaml.Unmarshal([]byte(str_src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	log.Printf("unmarshaled %v\n", pi3)

	format := "expected %v for %s but found %v"

	var expected_id item.ID = "core.deluxe.attribute.strength"
	actual_id := pi3.ID()
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "id", actual_id)
	}

	expected_s := "Strength"
	actual_s := pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "physical strnegth"
	actual_s = pi3.Notes()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "notes", actual_s)
	}

	expected_s = "Str"
	actual_s = pi3.ShortName()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "short_name", actual_s)
	}

	expected_s = "Might"
	actual_s = pi3.Description()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "description", actual_s)
	}

	expected := 1
	actual := pi3.Level()
	if expected != actual {
		t.Errorf(format, expected, "level", actual)
	}

}
