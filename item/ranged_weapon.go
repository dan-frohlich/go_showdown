package item

import (
	"log"
)

type RangedWeapon struct {
	Weapon
	Range_ string `json:"range" yaml:"range"`
	RoF_   int    `json:"RoF" yaml:"RoF"`
	_range Range
}

//Ranged
func (w *RangedWeapon) Range() Range {
	if w._range.short != 0 {
		return w._range
	}

	r, err := ParseRangeBand(w.Range_)
	if err != nil {
		log.Println(err)
	}
	w._range = r
	return w._range
}

func (w *RangedWeapon) RateOfFire() int { return w.RoF_ }
