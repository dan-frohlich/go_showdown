package character

import "showdown/identity"

type Power struct {
	ID_          identity.ID `json:"id" yaml:"id"`
	Name_        string `json:"name" yaml:"name"`
	Short_Name_  string `json:"short_name" yaml:"short_name"`
	Description_ string `json:"description" yaml:"description"`
	Cost_        int `json:"cost" yaml:"cost"`
	CV_          int `json:"cv" yaml:"cv"`
	Range_       string  `json:"range" yaml:"range"`
	Duration_    string  `json:"range" yaml:"range"`
}

//Identifiable
func (i *Power) ID() identity.ID { return i.ID_ }

//Named
func (i *Power) Name() string { return i.Name_ }

func (i *Power) ShortName() string { return i.Short_Name_ }

func (i *Power) Description() string { return i.Description_ }

func (i *Power) Cost() int { return i.Cost_ }

func (i *Power) CV() int { return i.CV_ }

func (i *Power) Range() string { return i.Range_ }

func (i *Power) Duration() string { return i.Duration_ }
