package item

import (
	"testing"
)

func TestDamageParse(t *testing.T) {

	format := "expected %v for %s but found %v"

	twoD6 := "2d6"

	damage, err := ParseDamage(twoD6)
	if err != nil {
		t.Errorf("parsing %s, error: %v", twoD6, err)
	}

	actual := damage.damage_bonus
	expected := 0
	if actual != expected {
		t.Errorf(format, expected, "damage bonus", actual)
	}

	actual = damage.dice_count
	expected = 2
	if actual != expected {
		t.Errorf(format, expected, "dice count", actual)
	}

	actual = damage.dice_sides
	expected = 6
	if actual != expected {
		t.Errorf(format, expected, "dice sides", actual)
	}

	actual_s := damage.base_stat
	expected_s := ""
	if actual_s != expected_s {
		t.Errorf(format, expected_s, "base stat", actual_s)
	}

	actual = damage.Max()
	expected = 12
	if actual != expected {
		t.Errorf(format, expected, "damage max", actual)
	}

	sword := "Str+d8+1"

	damage, err = ParseDamage(sword)
	if err != nil {
		t.Errorf("parsing %s, error: %v", sword, err)
	}

	actual = damage.damage_bonus
	expected = 1
	if actual != expected {
		t.Errorf(format, expected, "damage bonus", actual)
	}

	actual = damage.dice_count
	expected = 1
	if actual != expected {
		t.Errorf(format, expected, "dice count", actual)
	}

	actual = damage.dice_sides
	expected = 8
	if actual != expected {
		t.Errorf(format, expected, "dice sides", actual)
	}

	actual_s = damage.base_stat
	expected_s = "Str"
	if actual_s != expected_s {
		t.Errorf(format, expected_s, "base stat", actual_s)
	}

	actual = damage.Max()
	expected = 9
	if actual != expected {
		t.Errorf(format, expected, "damage max", actual)
	}

	whazit := "Str+2d4-1"

	damage, err = ParseDamage(whazit)
	if err != nil {
		t.Errorf("parsing %s, error: %v", whazit, err)
	}

	actual = damage.damage_bonus
	expected = -1
	if actual != expected {
		t.Errorf(format, expected, "damage bonus", actual)
	}

	actual = damage.dice_count
	expected = 2
	if actual != expected {
		t.Errorf(format, expected, "dice count", actual)
	}

	actual = damage.dice_sides
	expected = 4
	if actual != expected {
		t.Errorf(format, expected, "dice sides", actual)
	}

	actual_s = damage.base_stat
	expected_s = "Str"
	if actual_s != expected_s {
		t.Errorf(format, expected_s, "base stat", actual_s)
	}

	actual = damage.Max()
	expected = 7
	if actual != expected {
		t.Errorf(format, expected, "damage max", actual)
	}
}

func TestDamageStrings(t *testing.T) {
	d := Damage{dice_count: 2, dice_sides: 6, damage_bonus: 1}
	ds := d.String()
	if ds != "2d6+1" {
		t.Errorf("expected '2d6+1' but got %s", ds)
	}

	d2 := Damage{base_stat: "Str", dice_count: 1, dice_sides: 8}
	d2s := d2.String()
	if d2s != "Str+d8" {
		t.Errorf("expected 'Str+d8' but got %s", d2s)
	}
}

func TestDamageMax(t *testing.T) {
	d := Damage{dice_count: 2, dice_sides: 6, damage_bonus: 1}
	dm := d.Max()
	if dm != 13 {
		t.Errorf("expected '13' but got %s", dm)
	}

	d2 := Damage{base_stat: "Str", dice_count: 1, dice_sides: 8}
	d2m := d2.Max()
	if d2m != 8 {
		t.Errorf("expected '8' but got %s", d2m)
	}
}

func TestRoundTripDamage(t *testing.T) {

	format := "expected %v for %s but found %v"

	expected := "2d6"
	damage, _ := ParseDamage(expected)
	actual := damage.String()
	if expected != actual {
		t.Errorf(format, expected, "round.trip.2d6", actual)
	}

	expected = "2d6+1"
	damage, _ = ParseDamage(expected)
	actual = damage.String()
	if expected != actual {
		t.Errorf(format, expected, "round.trip.2d6+1", actual)
	}

	expected = "Str+d8"
	damage, _ = ParseDamage(expected)
	actual = damage.String()
	if expected != actual {
		t.Errorf(format, expected, "round.trip.Str+d8", actual)
	}
}
