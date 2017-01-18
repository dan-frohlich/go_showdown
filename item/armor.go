package item

type Armor struct {
	Item
	Armor_ int `json:"armor" yaml:"armor"`
	Parry_ int `json:"parry" yaml:"parry"`
}

func (a *Armor) Armor() int { return a.Armor_ }

func (a *Armor) Parry() int { return a.Parry_ }
