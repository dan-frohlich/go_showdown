package item

import (
	"math"
)

type WeaponCalculator interface {
	CalculateWeaponCost(w Weapon) int
}

type RangedWeaponCalculator interface {
	CalculateRangedWeaponCost(w RangedWeapon) int
}

type Showdown2009Calculator struct {
}

func (c Showdown2009Calculator) CalculateWeaponCost(w Weapon) int {
	// extracted from troop builder spreadsheet...
	//B Damage
	//C Dam Bonus
	//D AP
	//E Reach
	//F Parry
	//H Special
	//(((B11/2)−1)×3)+(3×C11)+(2×D11)+(E11×2)+(F11×3)+H11

	damage := w.Damage()
	B := damage.dice_count * damage.dice_sides
	C := damage.damage_bonus
	D := w.ArmorPiercing()
	E := 0 //todo support reach !?!
	F := 0 //todo support parry
	H := 0 //todo support special
	return ((B/2)-1)*3 + 3*C + 2*D + E*2 + F*3 + H

	//Cost	Special
	//-3	Requirement (Lance)
	//-5	Malfunction
	//3	Ignores Shield Bonus
}

func roundToInt(f float64) int {
	return int(math.Floor(f + .5))
}

func (c Showdown2009Calculator) CalculateRangedWeaponCost(w RangedWeapon) int {
	// extracted from troop builder spreadsheet...
	//B short range
	//C max damage
	//D ROF
	//E AP
	//F Special
	//G One Shot?
	//((2+(B3/6)+(C3−10)+((D3−1)×10)+(E3×3)+F3))×(IF(G3=1,0.16,1))
	B := w.Range().Short()
	C := w.Damage().Max()
	D := w.RateOfFire()
	E := w.ArmorPiercing()
	F := 0 //special?!?!
	G := 0 //only 1 shot?!?
	oneShotFactor := 1.0
	if G == 1 {
		oneShotFactor = 0.16
	}
	baseCost := (2 + (B / 6)) + (C - 10) + ((D - 1) * 10) + (E * 3) + F
	return roundToInt(float64(baseCost) * oneShotFactor)

	//Cost	Special Modifier
	//-5	May Not Move
	//20	Ignores Armor
	//-5	Snapfire
	//5	Small Burst Template
	//10	Medium Burst Template
	//15	Large Burst Template
	//15	Cone Template
	//20	Heavy Weapon
	//-10	Fixed Weapon
}
