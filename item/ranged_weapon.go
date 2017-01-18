package item

import (
	"encoding/json"
	"log"
	//"errors"
	"errors"
	"fmt"
	"reflect"
)

type RangedWeapon struct {
	Weapon
	Range_ Range `json:"range" yaml:"range"`
	RoF_   int   `json:"RoF" yaml:"RoF"`
}

//Ranged
func (w *RangedWeapon) Range() Range { return w.Range_ }

func (w *RangedWeapon) RateOfFire() int { return w.RoF_ }

func (u *RangedWeapon) MarshalJSON() ([]byte, error) {
	type Alias RangedWeapon

	//log.Printf("marshaling %v\n", u)
	b, e := json.Marshal(&struct {
		Range_ string `json:"range"`
		*Alias
	}{
		Range_: u.Range_.String(),
		Alias:  (*Alias)(u),
	})

	//log.Printf("marshaled %v, %v\n", string(b), e)
	return b, e
}

func (u *RangedWeapon) UnmarshalJSON(data []byte) error {
	return u.unmarshalJSONTheDorkyWay(data)
}

func (u *RangedWeapon) unmarshalJSONTheDorkyWay(data []byte) error {
	//defer log.Printf("unmarshaled RangedWeapon as %v\n", u)

	var stuff map[string]interface{}
	var weapon_ Weapon

	if err := json.Unmarshal(data, &weapon_); err != nil {
		log.Printf("failed to unmarshal into weapon_ %v\n", weapon_)
		return err
	}

	u.Weapon = weapon_

	if err := json.Unmarshal(data, &stuff); err != nil {
		log.Printf("failed to unmarshal into map %v\n", stuff)
		return err
	}

	if range_string_, ok := stuff["range"].(string); ok {
		var err error
		if u.Range_, err = ParseRangeBand(range_string_); err != nil {
			log.Printf("failed to parse aux range band %v\n", range_string_)
			return err
		}
	} else {
		return errors.New("range must be a string")
	}

	rof_interface := stuff["RoF"]
	if rof_interface != nil {
		if rof_int, ok := rof_interface.(float64); ok {
			u.RoF_ = int(rof_int)
		} else {
			message := fmt.Sprintf("RoF [%v] must be an int but was %v",
				rof_interface, reflect.TypeOf(rof_interface))
			return errors.New(message)
		}
	} else {
		u.RoF_ = 1 //default to RoF 1
	}

	//log.Printf("found %v - %v\n", weapon_, stuff)
	return nil
}

func (u *RangedWeapon) unmarshalJSONTheCoolWay(data []byte) error {
	defer log.Printf("unmarshaled RangedWeapon as %v\n", u)

	log.Printf("unmarshaling RangedWeapon %v\n", string(data))
	type Alias RangedWeapon
	aux := &struct {
		Range_ string `json:"range"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		log.Printf("failed to unmarshal into aux struxt %v\n", aux)
		return err
	}
	log.Printf("read Item: %v ", aux.Weapon.Item)
	log.Printf("read weapon: %v ", aux.Weapon)
	log.Printf("read range: %s ", aux.Range_)
	log.Printf("unmarshaled aux struxt for RangedWeapon %v - %v\n", aux.Range_, aux.Alias)

	rb, err := ParseRangeBand(aux.Range_)
	if err != nil {
		log.Printf("failed to parse aux range band %v\n", aux.Range_)
		return err
	}
	log.Printf("parsed aux range band %s\n", rb)
	u.Range_ = rb
	return nil
}
