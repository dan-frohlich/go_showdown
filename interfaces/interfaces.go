package interfaces

import (
	"showdown/item"
)

type Identifiable interface {
	ID() string
}

type Named interface {
	Name() string
	ShortName() string
	Description() string
}

type Noted interface {
	Notes() string
}

type Costly interface {
	Cost() int
	SpecialCost() int
}

type Leveled interface {
	Level() int
}

type Damaging interface {
	Damage() string
	ArmorPiercing() int
}

type Ranged interface {
	Range() string
	RateOfFire() int
}

type Hand interface {
	Parry() int
	Reach() int
}

type Armored interface {
	Armor() int
	Parry() int
}

type HandWeaponCalculator interface {
	CalculateHandWeaponCost(w item.HandWeapon) int
}

type RangedWeaponCalculator interface {
	CalculateRangedWeaponCost(w item.RangedWeapon) int
}

type ArmorCalculator interface {
	CalculateArmorCost(a item.Armor) int
}
