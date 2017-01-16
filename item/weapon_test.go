package item

import (
	"testing"
	"encoding/json"
	"github.com/ghodss/yaml"
)

func TestWeaponReadWrite(t *testing.T) {

	pi := &Weapon{
		Item:   Item{Name_:  "pitem", Cost_:0},
		Damage_:Damage{dice_count:1, dice_sides:6, base_stat:"Str"} }
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
`

	pi3 := &Weapon{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	if pi3.Cost_ != 101 ||
			pi3.ID_ != "0-00-000-0" ||
			pi3.Name_ != "pitem" ||
			pi3.Description_ != "pidescr" ||
			pi3.Damage().String() != "2d6" {
		t.Errorf("got %s from %s", pi3, src_yml)
	}
}

func TestRangedWeaponReadYaml(t *testing.T) {
	src_yml := `
id: 0-00-000-0
name: pistol
description: typical pistol
cost: 2
damage: 2d6
range: "12/24/48"
`

	pi3 := &RangedWeapon{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	if pi3.Cost_ != 101 ||
			pi3.ID_ != "0-00-000-0" ||
			pi3.Name_ != "pitem" ||
			pi3.Description_ != "pidescr" ||
			pi3.Damage().String() != "2d6" {
		t.Errorf("got %s from %s", pi3, src_yml)
	}
}
