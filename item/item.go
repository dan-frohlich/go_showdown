package item

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
	Range() string
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
