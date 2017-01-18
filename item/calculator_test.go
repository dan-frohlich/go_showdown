package item

import "testing"

func TestHandWeaponCalculations(t *testing.T) {

	c := Showdown2009Calculator{}
	format := "expected %v for %s but found %v"

	//test data from ShowdownTroopBuilder.xsl ...
	//https://s3-us-west-2.amazonaws.com/peg-freebies/ShowdownTroopBuilder.xls
	dagger := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Dagger"}},
			Damage_: "1d4",
			AP_:     0}}
	actual := c.CalculateHandWeaponCost(dagger)
	expected := 3
	if expected != actual {
		t.Errorf(format, expected, "dagger.cost", actual)
	}

	gs := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Great Sword"}},
			Damage_: "1d10",
			AP_:     0}}
	actual = c.CalculateHandWeaponCost(gs)
	expected = 12
	if expected != actual {
		t.Errorf(format, expected, "Great.Sword.cost", actual)
	}

	ls := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Laser Sword"}},
			Damage_: "d6+8",
			AP_:     12}}
	actual = c.CalculateHandWeaponCost(ls)
	expected = 54
	if expected != actual {
		t.Errorf(format, expected, "Laser.Sword.cost", actual)
	}

	ga := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Great Axe"}},
			Damage_: "1d10",
			AP_:     1},
		Parry_: -1}

	expected = 11
	actual = c.CalculateHandWeaponCost(ga)
	if expected != actual {
		t.Errorf(format, expected, "Great Axe.cost", actual)
	}

	flail := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Flail"}, SpecialCost_: 3},
			Damage_: "1d6"}}

	expected = 9
	actual = c.CalculateHandWeaponCost(flail)
	if expected != actual {
		t.Errorf(format, expected, "flail.cost", actual)
	}

	lance := HandWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Lance"}, SpecialCost_: -3},
			Damage_: "1d8",
			AP_:     2},
		Reach_: 2}

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
			Item:    Item{Element: Element{Name_: "Bazooka"}, SpecialCost_: 10},
			Damage_: "4d8",
			AP_:     9},
		_range: Range{short: 96},
		RoF_:   1}
	actual := c.CalculateRangedWeaponCost(rw)
	expected := 78 //77
	if expected != actual {
		t.Errorf(format, expected, "Bazooka.cost", actual)
	}

	rw = RangedWeapon{
		Weapon: Weapon{
			Item:    Item{Element: Element{Name_: "Kentucky Rifle (.45)"}},
			Damage_: "2d8",
			AP_:     2},
		_range: Range{short: 15},
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
		Item:   Item{Element: Element{Name_: "Shield Large"}, SpecialCost_: 1},
		Armor_: 2,
		Parry_: 2}

	actual := c.CalculateArmorCost(a)
	expected := 8
	if expected != actual {
		t.Errorf(format, expected, "large shield.cost", actual)
	}
}
