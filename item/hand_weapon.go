package item

import (
	"encoding/json"
	"log"
	"errors"
	"fmt"
	"reflect"
)

type HandWeapon struct {
	Weapon
	Reach_ int `json:"reach" yaml:"reach"`
	Parry_ int `json:"parry" yaml:"parry"`
}

func (w *HandWeapon) Parry() int { return w.Parry_ }

func (w *HandWeapon) Reach() int { return w.Reach_ }

func (u *HandWeapon) UnmarshalJSON(data []byte) error {
	return u.unmarshalJSONTheDorkyWay(data)
}

func (u *HandWeapon) unmarshalJSONTheDorkyWay(data []byte) error {
	//defer log.Printf("unmarshaled HandWeapon as %v\n", u)

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

	parry_interface := stuff["parry"]
	if parry_interface != nil {
		if rof_int, ok := parry_interface.(float64); ok {
			u.Parry_ = int(rof_int)
		} else {
			message := fmt.Sprintf("Parry [%v] must be an int but was %v",
				parry_interface, reflect.TypeOf(parry_interface))
			return errors.New(message)
		}
	} else {
		u.Parry_ = 0 //default to Parry 0
	}
	
	reach_interface := stuff["reach"]
	if reach_interface != nil {
		if rof_int, ok := reach_interface.(float64); ok {
			u.Reach_ = int(rof_int)
		} else {
			message := fmt.Sprintf("reach [%v] must be an int but was %v",
				reach_interface, reflect.TypeOf(reach_interface))
			return errors.New(message)
		}
	} else {
		u.Reach_ = 0 //default to reach 0
	}

	//log.Printf("found %v - %v\n", weapon_, stuff)
	return nil
}

//func (u *HandWeapon) unmarshalJSONTheCoolWay(data []byte) error {
//	defer log.Printf("unmarshaled HandWeapon as %v\n", u)
//
//	log.Printf("unmarshaling HandWeapon %v\n", string(data))
//	type Alias HandWeapon
//	aux := &struct {
//		Range_ string `json:"range"`
//		*Alias
//	}{
//		Alias: (*Alias)(u),
//	}
//
//	if err := json.Unmarshal(data, &aux); err != nil {
//		log.Printf("failed to unmarshal into aux struxt %v\n", aux)
//		return err
//	}
//	log.Printf("read Item: %v ", aux.Weapon.Item)
//	log.Printf("read weapon: %v ", aux.Weapon)
//	log.Printf("read range: %s ", aux.Range_)
//	log.Printf("unmarshaled aux struxt for HandWeapon %v - %v\n", aux.Range_, aux.Alias)
//
//	rb, err := ParseRangeBand(aux.Range_)
//	if err != nil {
//		log.Printf("failed to parse aux range band %v\n", aux.Range_)
//		return err
//	}
//	log.Printf("parsed aux range band %s\n", rb)
//	u.Range_ = rb
//	return nil
//}
