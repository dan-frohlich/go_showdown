package item

import (
	"encoding/json"
	"log"
)

type Identifiable interface {
	ID() string
}

type Named interface {
	Name() string
	Description() string
}

type Costly interface {
	Named
	Cost() int
}

type Damaging interface {
	Costly
	Damage() string
}

type Ranged interface {
	Damaging
	RangeBand() string
}

type Item struct {
	ID_          string `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int `json:"cost" yaml:"cost"`
}

//Named
func (i *Item) Name() string { return i.Name_ }

func (i *Item) Description() string { return i.Description_ }

//Identifiable
func (i *Item) ID() string { return i.ID_ }

//Costly
func (i *Item) Cost() int { return i.Cost_ }

type Weapon struct {
	ID_          string `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int `json:"cost" yaml:"cost"`
	Damage_      Damage `json:"damage" yaml:"damage"`
}

//Damaging
func (w *Weapon) Damage() Damage { return w.Damage_ }

type RangedWeapon struct {
	Weapon
	RangeBand string `json:"range_band" yaml:"range_band"`
}

func (u *Weapon) MarshalJSON() ([]byte, error) {
	type Local_Alias Weapon

	log.Printf("marshaling %v\n", u)
	b, e := json.Marshal(&struct {
		Damage_ string `json:"damage"`
		*Local_Alias
	}{
		Damage_:      u.Damage_.String(),
		Local_Alias:  (*Local_Alias)(u),
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
