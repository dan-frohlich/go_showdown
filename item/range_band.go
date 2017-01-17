package item

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)

type Range struct {
	short int
}

func (rb Range) Short() int {
	return rb.short
}

func (rb Range) String() string {
	s := rb.short
	m := 2 * rb.short
	l := 4 * rb.short
	return fmt.Sprintf("%d/%d/%d", s, m, l)
}

func ParseRangeBand(s string) (Range, error) {
	items := strings.Split(s, "/")
	if len(items) != 3 {
		return Range{}, errors.New("'" + s + "' is not a range band.")
	}
	s0 := strings.TrimSpace(items[0])
	i0, err := strconv.ParseInt(s0, 10, 0)
	return Range{int(i0)}, err
}
