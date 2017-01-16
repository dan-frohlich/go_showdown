package item

import (
	"testing"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

func TestItemReadWrite(t *testing.T) {

	pi := &Item{Name_:"pitem", Cost_:0}
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

	if pi3.Cost() != 101 ||
			pi3.ID_ != "0-00-000-0" ||
			pi3.Name() != "pitem" ||
			pi3.Description() != "pidescr" {
		t.Errorf("got %s from %s", pi3, src_yml)
	}
}
