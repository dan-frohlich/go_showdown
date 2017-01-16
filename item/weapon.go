package item

import (
	"log"
	"encoding/json"
)

type Weapon struct {
	Item
	Damage_ Damage `json:"damage" yaml:"damage"`
}

//Damaging
func (w *Weapon) Damage() Damage { return w.Damage_ }

type RangedWeapon struct {
	Weapon
	RangeBand_ RangeBand `json:"range_band" yaml:"range_band"`
}

func (u *Weapon) MarshalJSON() ([]byte, error) {
	type Alias Weapon

	log.Printf("marshaling %v\n", u)
	b, e := json.Marshal(&struct {
		Damage_ string `json:"damage"`
		*Alias
	}{
		Damage_:      u.Damage_.String(),
		Alias:        (*Alias)(u),
	})

	log.Printf("marshaled %v, %v\n", string(b), e)
	return b, e
}

func (u *Weapon) UnmarshalJSON(data []byte) error {
	log.Printf("unmarshaling %v\n", string(data))
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
	log.Printf("unmarshaled aux struxt %v\n", aux)

	d, err := ParseDamage(aux.Damage_)
	if err != nil {
		log.Printf("failed to parse aux damage %v\n", aux.Damage_)
		return err
	}
	log.Printf("parsed aux damage %s\n", d)
	u.Damage_ = d
	return nil
}

func (u *RangedWeapon) MarshalJSON() ([]byte, error) {
	type Alias RangedWeapon

	log.Printf("marshaling %v\n", u)
	b, e := json.Marshal(&struct {
		RangeBand_ string `json:"damage"`
		*Alias
	}{
		RangeBand_:      u.RangeBand_.String(),
		Alias:           (*Alias)(u),
	})

	log.Printf("marshaled %v, %v\n", string(b), e)
	return b, e
}

func (u *RangedWeapon) UnmarshalJSON(data []byte) error {
	log.Printf("unmarshaling %v\n", string(data))
	type Alias RangedWeapon
	aux := &struct {
		RangeBand_ string `json:"damage"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		log.Printf("failed to unmarshal into aux struxt %v\n", aux)
		return err
	}
	log.Printf("unmarshaled aux struxt %v\n", aux)

	rb, err := ParseRangeBand(aux.RangeBand_)
	if err != nil {
		log.Printf("failed to parse aux range band %v\n", aux.RangeBand_)
		return err
	}
	log.Printf("parsed aux range band %s\n", rb)
	u.RangeBand_ = rb
	return nil
}
