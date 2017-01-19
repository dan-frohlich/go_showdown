package trait

import (
	"github.com/ghodss/yaml"
	"log"
	"testing"
	"showdown/identity"
)

func TestSkillParse(t *testing.T) {
	src_yml := `
id: core.deluxe.skill.climb
linked_attribute: core.deluxe.attribute.strength
name: Climb
short_name: Clm
description: climbing
notes: get on up now
level: 1
`

	pi3 := &Attribute{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	log.Printf("unmarshaled %v\n", pi3)

	format := "expected %v for %s but found %v"

	var expected_id identity.ID = "core.deluxe.skill.climb"
	actual_id := pi3.ID()
	if expected_id != actual_id {
		t.Errorf(format, expected_id, "id", actual_id)
	}

	expected_s := "Climb"
	actual_s := pi3.Name()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "name", actual_s)
	}

	expected_s = "get on up now"
	actual_s = pi3.Notes()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "notes", actual_s)
	}

	expected_s = "Clm"
	actual_s = pi3.ShortName()
	if expected_s != actual_s {
		t.Errorf(format, expected_s, "short_name", actual_s)
	}

	expected_s = "climbing"
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
