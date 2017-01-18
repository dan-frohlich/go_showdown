package item

import (
	"encoding/json"
	"testing"

	"github.com/ghodss/yaml"
)

func TestItemReadWrite(t *testing.T) {

	pi := &Item{Name_: "pitem", Cost_: 0}
	pij, imerr := json.Marshal(pi)

	if imerr != nil {
		t.Error(imerr)
	}

	pi2 := &Item{}
	iumerror := json.Unmarshal(pij, pi2)

	if iumerror != nil {
		t.Error(iumerror)
	}

	if pi.Name() != pi2.Name() {
		t.Errorf("got %s when expecting %s", pi2.Name(), pi.Name())
	}
}

func TestItemReadYaml(t *testing.T) {
	src_yml := `
id: 0-00-000-0
name: pitem
description: pidescr
cost: 101
`

	pi3 := &Item{}
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

}
