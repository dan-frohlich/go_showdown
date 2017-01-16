package item

import (
	"fmt"
	"strings"
	"errors"
	"strconv"
)

type RangeBand struct {
	short int
}

func (rb RangeBand) Short() int {
	return rb.short
}

func (rb RangeBand) String() string {
	s := rb.short
	m := 2 * rb.short
	l := 3 * rb.short
	return fmt.Sprintf("%d/%d/%d", s, m, l)
}

func ParseRangeBand(s string) (RangeBand, error) {
	items := strings.Split(s, "/")
	if len(items) != 3 {
		return RangeBand{}, errors.New("'" + s + "' is not a range band.")
	}
	s0 := strings.TrimSpace(items[0])
	i0, err := strconv.ParseInt(s0, 10, 0)
	return RangeBand{int(i0)}, err
}
