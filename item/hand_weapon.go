package item

type HandWeapon struct {
	Weapon
	Reach_ int `json:"reach" yaml:"reach"`
	Parry_ int `json:"parry" yaml:"parry"`
}

func (w *HandWeapon) Parry() int { return w.Parry_ }

func (w *HandWeapon) Reach() int { return w.Reach_ }
