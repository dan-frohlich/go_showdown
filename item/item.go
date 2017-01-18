package item

type Identifiable interface {
	ID() string
}

type Named interface {
	Name() string
	Description() string
	Notes() string
}

type Costly interface {
	Named
	Cost() int
	SpecialCost() int
}

type Damaging interface {
	Costly
	Damage() string
	ArmorPiercing() int
}

type Ranged interface {
	Damaging
	Range() string
	RateOfFire() int
}

type Hand interface {
	Damaging
	Parry() int
	Reach() int
}

type Armored interface {
	Costly
	Armor() int
	Parry() int
}

type Item struct {
	ID_          string `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int    `json:"cost" yaml:"cost"`
	SpecialCost_ int    `json:"special_cost" yaml:"special_cost"`
	Notes_       string `json:"notes" yaml:"notes"`
}

//Named
func (i *Item) Name() string { return i.Name_ }

func (i *Item) Description() string { return i.Description_ }

func (i *Item) Notes() string { return i.Notes_ }

//Identifiable
func (i *Item) ID() string { return i.ID_ }

//Costly
func (i *Item) Cost() int { return i.Cost_ }

func (i *Item) SpecialCost() int { return i.SpecialCost_ }
