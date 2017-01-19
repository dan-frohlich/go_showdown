package character

import "showdown/identity"

type Edge struct {
	ID_          identity.ID `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Short_Name_  string `json:"short_name" yaml:"short_name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int `json:"cost" yaml:"cost"`
}

type Hindrance Edge

//Identifiable
func (i *Edge) ID() identity.ID { return i.ID_ }

//Named
func (i *Edge) Name() string { return i.Name_ }

func (i *Edge) ShortName() string { return i.Short_Name_ }

func (i *Edge) Description() string { return i.Description_ }

func (i *Edge) Cost() int { return i.Cost_ }
