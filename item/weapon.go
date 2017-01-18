package item

import (
	"encoding/json"
	"log"
)

type Weapon struct {
	Item
	Damage_ Damage `json:"damage" yaml:"damage"`
	AP_     int    `json:"AP" yaml:"AP"`
}

//Damaging
func (w *Weapon) Damage() Damage { return w.Damage_ }

func (w *Weapon) ArmorPiercing() int { return w.AP_ }

func (u *Weapon) MarshalJSON() ([]byte, error) {
	type Alias Weapon

	//log.Printf("marshaling %v\n", u)
	b, e := json.Marshal(&struct {
		Damage_ string `json:"damage"`
		*Alias
	}{
		Damage_: u.Damage_.String(),
		Alias:   (*Alias)(u),
	})

	//log.Printf("marshaled %v, %v\n", string(b), e)
	return b, e
}

func (u *Weapon) UnmarshalJSON(data []byte) error {
	//defer log.Printf("unmarshaled Weapon as %v\n", u)

	//log.Printf("unmarshaling Weapon %v\n", string(data))
	type Alias Weapon
	aux := &struct {
		Damage_ string `json:"damage"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		log.Printf("failed to unmarshal into aux struxt %v\n", aux)
		return err
	}
	//log.Printf("unmarshaled aux struxt for Weapon %v - %v\n", aux.Damage_, aux.Alias)

	d, err := ParseDamage(aux.Damage_)
	if err != nil {
		log.Printf("failed to parse aux damage %v\n", aux.Damage_)
		return err
	}
	//log.Printf("parsed aux damage %s\n", d)
	u.Damage_ = d
	return nil
}
