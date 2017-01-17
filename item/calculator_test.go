package item

import "testing"


func TestHandWeapons(t *testing.T) {

	c := Showdown2009Calculator{}
	format := "expected %v for %s but found %v"

	//test data from ShowdownTroopBuilder.xsl ...
	//https://s3-us-west-2.amazonaws.com/peg-freebies/ShowdownTroopBuilder.xls
	dagger := Weapon{Item:Item{Name_:"Dagger"}, Damage_:Damage{dice_sides:4, dice_count:1}, AP_:0}
	dagger_cost := c.CalculateWeaponCost(dagger)
	if 3 != dagger_cost {
		t.Errorf(format, 3, "dagger.cost", dagger_cost)
	}

	gs := Weapon{Item:Item{Name_:"Great Sword"}, Damage_:Damage{dice_sides:10, dice_count:1}, AP_:0}
	gs_cost := c.CalculateWeaponCost(gs)
	if 12 != gs_cost {
		t.Errorf(format, 12, "Great.Sword.cost", gs_cost)
	}

	ls := Weapon{Item:Item{Name_:"Laser Sword"}, Damage_:Damage{dice_sides:6, damage_bonus: 8, dice_count:1}, AP_:12}
	ls_cost := c.CalculateWeaponCost(ls)
	if 54 != ls_cost {
		t.Errorf(format, 54, "Laser.Sword.cost", ls_cost)
	}
}
