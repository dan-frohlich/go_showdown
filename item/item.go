package item

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
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Cost int `json:"cost"`
}

type Weapon struct {
	Item
	Damage string `json:"damage"`
}

type RangedWeapon struct {
	Weapon
	RangeBand string `json:"range_band"`
}
