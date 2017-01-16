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
	ID_          string `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int `json:"cost" yaml:"cost"`
}

//Named
func (i *Item) Name() string { return i.Name_ }

func (i *Item) Description() string { return i.Description_ }

func (i *Item) Cost() int { return i.Cost_ }

type Weapon struct {
	Item
	Damage string `json:"damage"`
}

type RangedWeapon struct {
	Weapon
	RangeBand string `json:"range_band"`
}
