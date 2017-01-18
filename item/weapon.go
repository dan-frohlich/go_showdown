package item

import "log"

type Weapon struct {
	Item
	Damage_ string `json:"damage" yaml:"damage"`
	AP_     int    `json:"AP" yaml:"AP"`
	_damage Damage
}

//Damaging
func (w *Weapon) Damage() Damage {
	if w._damage.initialized {
		return w._damage
	}
	var err error
	w._damage, err = ParseDamage(w.Damage_)
	if err != nil {
		log.Println(err)
	}
	return w._damage
}

func (w Weapon) ArmorPiercing() int { return w.AP_ }
