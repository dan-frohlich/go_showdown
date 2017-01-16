package item

import (
	"testing"
)

func TestDamageParse(t *testing.T) {

	twoD6 := "2d6"

	damage, err := ParseDamage(twoD6)
	if err != nil {
		t.Errorf("parsing %s, error: %v", twoD6, err)
	}
	if damage.damage_bonus != 0 {
		t.Errorf("parsing %s, parsed bonus damage: %s", twoD6, damage.damage_bonus)
	}
	if damage.dice_count != 2 {
		t.Errorf("parsing %s, parsed dice count: %s", twoD6, damage.dice_count)
	}
	if damage.dice_sides != 6 {
		t.Errorf("parsing %s, parsed dice sides: %s", twoD6, damage.dice_sides)
	}
	if damage.base_stat != "" {
		t.Errorf("parsing %s, parsed base stat: %s", twoD6, damage.base_stat)
	}
}

func TestDamageStrings(t *testing.T) {
	d := Damage{dice_count:2, dice_sides:6, damage_bonus:1}
	ds := d.String()
	if ds != "2d6+1" {
		t.Errorf("expected '2d6+1' but got %s", ds)
	}

	d2 := Damage{base_stat:"Str", dice_count:1, dice_sides:8}
	d2s := d2.String()
	if d2s != "Str+d8" {
		t.Errorf("expected 'Str+d8' but got %s", d2s)
	}
}

func TestDamageMax(t *testing.T) {
	d := Damage{dice_count:2, dice_sides:6, damage_bonus:1}
	dm := d.Max()
	if dm != 13 {
		t.Errorf("expected '13' but got %s", dm)
	}

	d2 := Damage{base_stat:"Str", dice_count:1, dice_sides:8}
	d2m := d2.Max()
	if d2m != 8 {
		t.Errorf("expected '8' but got %s", d2m)
	}
}