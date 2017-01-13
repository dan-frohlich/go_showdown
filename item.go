package main

type Item interface {
	Description() string
	Cost() float64
}

type Weapon interface {
	Item
	Damage() string
}

type Ranged interface {
	RangeBand() string
}

type PersistedItem struct {
	Description string
}
