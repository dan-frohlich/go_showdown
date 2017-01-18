package item

import "testing"

func TestHandWeaponCalculations(t *testing.T) {

	c := Showdown2009Calculator{}
	format := "expected %v for %s but found %v"

	//test data from ShowdownTroopBuilder.xsl ...
	//https://s3-us-west-2.amazonaws.com/peg-freebies/ShowdownTroopBuilder.xls
	dagger := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Dagger"},
			Damage_: Damage{dice_sides: 4, dice_count: 1},
			AP_:     0}}
	actual := c.CalculateHandWeaponCost(dagger)
	expected := 3
	if expected != actual {
		t.Errorf(format, expected, "dagger.cost", actual)
	}

	gs := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Great Sword"},
			Damage_: Damage{dice_sides: 10, dice_count: 1},
			AP_:     0}}
	actual = c.CalculateHandWeaponCost(gs)
	expected = 12
	if expected != actual {
		t.Errorf(format, expected, "Great.Sword.cost", actual)
	}

	ls := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Laser Sword"},
			Damage_: Damage{dice_sides: 6, damage_bonus: 8, dice_count: 1},
			AP_:     12}}
	actual = c.CalculateHandWeaponCost(ls)
	expected = 54
	if expected != actual {
		t.Errorf(format, expected, "Laser.Sword.cost", actual)
	}

	ga := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Great Axe"},
			Damage_: Damage{dice_sides: 10, dice_count: 1},
			AP_:     1},
		Parry_: -1 }

	expected = 11
	actual = c.CalculateHandWeaponCost(ga)
	if expected != actual {
		t.Errorf(format, expected, "Great Axe.cost", actual)
	}

	flail := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Flail", SpecialCost_: 3 },
			Damage_: Damage{dice_sides: 6, dice_count: 1} }}

	expected = 9
	actual = c.CalculateHandWeaponCost(flail)
	if expected != actual {
		t.Errorf(format, expected, "flail.cost", actual)
	}

	lance := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Lance", SpecialCost_: -3 },
			Damage_: Damage{dice_sides: 8, dice_count: 1},
			AP_:     2 },
		Reach_:2}

	expected = 14
	actual = c.CalculateHandWeaponCost(lance)
	if expected != actual {
		t.Errorf(format, expected, "lance.cost", actual)
	}
}

func TestRangedWeaponCalculations(t *testing.T) {

	c := Showdown2009Calculator{}
	format := "expected %v for %s but found %v"

	rw := RangedWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Bazooka", SpecialCost_:10},
			Damage_: Damage{dice_sides: 8, dice_count: 4},
			AP_:     9 },
		Range_: Range{short:96},
		RoF_:   1}
	actual := c.CalculateRangedWeaponCost(rw)
	expected := 78 //77
	if expected != actual {
		t.Errorf(format, expected, "Bazooka.cost", actual)
	}

	rw = RangedWeapon{
		Weapon: Weapon{
			Item:    Item{Name_: "Kentucky Rifle (.45)" },
			Damage_: Damage{dice_sides: 8, dice_count: 2},
			AP_:     2 },
		Range_: Range{short:15},
		RoF_:   1}
	actual = c.CalculateRangedWeaponCost(rw)
	expected = 17
	if expected != actual {
		t.Errorf(format, expected, "Kentucky Rifle.cost", actual)
	}

}

func TestArmorCalculations(t *testing.T) {

	c := Showdown2009Calculator{}
	format := "expected %v for %s but found %v"

	a := Armor{
		Item:   Item{Name_:"Shield Large", SpecialCost_:1},
		Armor_: 2,
		Parry_: 2 }

	actual := c.CalculateArmorCost(a)
	expected := 8
	if expected != actual {
		t.Errorf(format, expected, "large shield.cost", actual)
	}
}
