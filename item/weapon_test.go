package item

import (
	"testing"
	"encoding/json"
	"github.com/ghodss/yaml"
	"log"
	"fmt"
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
	//t.Skip("skipping test;")

	log.Printf("running TestRangedWeaponReadYaml")
	src_yml := `
id: 0-00-000-0
name: pistol
description: typical pistol
cost: 2
damage: 2d6
range: 12/24/48
`

	pi3 := &RangedWeapon{}
	pi3_umerror := yaml.Unmarshal([]byte(src_yml), pi3)
	if pi3_umerror != nil {
		t.Error(pi3_umerror)
	}

	field_assertion("cost", 2, pi3.Cost_)
	field_assertion("id", "0-00-000-0", pi3.ID_)
	field_assertion("name", "pistol", pi3.Name_)
	field_assertion("description", "typical pistol", pi3.Description_)
	field_assertion("damage", "2d6", pi3.Damage_.String())
	field_assertion("max damage", 12, pi3.Damage_.Max())
	field_assertion("range fields", "12/24/48", pi3.Range_.String())
}

func field_assertion(field string, expected interface{}, actual interface{} ){
	if expected != actual {
		fmt.Errorf("expected %s of %v but found %v", field, expected, actual)
	}
}
