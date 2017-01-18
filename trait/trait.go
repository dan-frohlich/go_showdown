package trait

import "showdown/item"

type Trait struct {
	item.Element
	Level_ int `json:"level" yaml:"level"`
}

func (t *Trait) Level() int { return t.Level_ }
