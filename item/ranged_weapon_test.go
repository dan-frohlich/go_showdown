package item

import (
	"testing"
	"github.com/ghodss/yaml"
)

func TestRangedWeaponReadYaml(t *testing.T) {
	//t.Skip("skipping test;")

	//log.Printf("running TestRangedWeaponReadYaml")
	src_yml := `
id: 0-00-000-0
name: pistol
description: atypical pistol
cost: 2
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
	if 2 != pi3.Cost() {
		t.Errorf(format, 2, "cost", pi3.Cost())
	}

	if "0-00-000-0" != pi3.ID() {
		t.Errorf(format, "0-00-000-0", "id", pi3.ID())
	}

	if "pistol" != pi3.Name() {
		t.Errorf(format, "pistol", "name", pi3.Name())
	}

	if "atypical pistol" != pi3.Description() {
		t.Errorf(format, "atypical pistol", "description", pi3.Description())
	}

	if "2d6" != pi3.Damage_.String() {
		t.Errorf(format, "2d6", "damage", pi3.Damage_.String())
	}

	if 2 != pi3.ArmorPiercing() {
		t.Errorf(format, 2, "AP", pi3.ArmorPiercing())
	}

	if 12 != pi3.Damage_.Max() {
		t.Errorf(format, 12, "max damage", pi3.Damage_.Max())
	}

	if 12 != pi3.Range().Short() {
		t.Errorf(format, 12, "short range", pi3.Range().Short())
	}
	if "12/24/48" != pi3.Range_.String() {
		t.Errorf(format, "2d6", "range fields", pi3.Range_.String())
	}
	if 3 != pi3.RateOfFire() {
		t.Errorf(format, 3, "RoF", pi3.RateOfFire())
	}

}
