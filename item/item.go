package item

import (
	"showdown/identity"
)

type Element struct {
	ID_          identity.ID `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Short_Name_  string `json:"short_name" yaml:"short_name"`
	Description_ string `json:"description" yaml:"description"`
	Notes_       string `json:"notes" yaml:"notes"`
}

type Item struct {
	Element
	Count_       int `json:"count" yaml:"count"`
	Cost_        int `json:"cost" yaml:"cost"`
	SpecialCost_ int `json:"special_cost" yaml:"special_cost"`
}

//Identifiable
func (i *Element) ID() identity.ID { return i.ID_ }

//Named
func (i *Element) Name() string { return i.Name_ }

func (i *Element) ShortName() string { return i.Short_Name_ }

func (i *Element) Description() string { return i.Description_ }

func (i *Element) Notes() string { return i.Notes_ }

//Costly
func (i *Item) Cost() int {
	return i.Cost_
}

func (i *Item) SpecialCost() int { return i.SpecialCost_ }
