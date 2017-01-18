package item

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Damage struct {
	dice_count   int
	dice_sides   int
	damage_bonus int
	base_stat    string
	shotgun_min  int
	shotgun_max  int
	initialized  bool
}

func (d Damage) Max() int { return d.dice_sides*d.dice_count + d.damage_bonus }

func (d Damage) ShotgunDamage() bool { return d.shotgun_min != 0 && d.shotgun_max != 0 }

func ParseDamage(source string) (Damage, error) {
	//todo actually implement the parser for these cases...
	//Str+d8
	//2d6+1
	//Str+d6+1
	//validate
	r, _ := regexp.Compile("^[ \\t]*((([a-z A-Z]{3}\\+){1}|)[0-9]*d[0-9]+)([\\+\\-]{1}[0-9]+|)[ \\t]*$")

	if !r.MatchString(source) {
		return Damage{}, errors.New(fmt.Sprintf("could not parse damage '%s'", source))
	}

	src2 := strings.TrimSpace(source)

	split1 := strings.Split(src2, "d")
	if len(split1) != 2 {
		return Damage{}, errors.New(fmt.Sprintf("no 'd' in %s ?\n", src2))
	}
	d := Damage{dice_count: 1, initialized: true}
	left := split1[0]
	right := split1[1]
	//log.Printf("splitting [%s] on [%s] => [%s] [%s]", src2, "d", left, right)

	if left != "" {
		if err := parseLeft(&d, src2, left); err != nil {
			return d, err
		}
	}
	if i, err := strconv.ParseInt(right, 10, 64); err == nil {
		d.dice_sides = int(i)
	} else if strings.Contains(right, "+") {
		bonus_split := strings.Split(right, "+")
		if i, err := strconv.ParseInt(bonus_split[0], 10, 64); err == nil {
			d.dice_sides = int(i)
		}
		if i, err := strconv.ParseInt(bonus_split[1], 10, 64); err == nil {
			d.damage_bonus = int(i)
		}
	} else if strings.Contains(right, "-") {
		bonus_split := strings.Split(right, "-")
		if i, err := strconv.ParseInt(bonus_split[0], 10, 64); err == nil {
			d.dice_sides = int(i)
		}
		if i, err := strconv.ParseInt(bonus_split[1], 10, 64); err == nil {
			d.damage_bonus = int(i) * (-1)
		}
	}

	//log.Printf("returning %v", d)
	return d, nil
}

func parseLeft(d *Damage, src2, left string) error {
	if i, err := strconv.ParseInt(left, 10, 64); err == nil {
		d.dice_count = int(i)
	} else {
		//must begin with Str+
		left_len := len(left)
		if left_len >= 4 {
			d.base_stat = left[:3]
			if left_len > 4 {
				if i, err := strconv.ParseInt(left[4:], 10, 64); err == nil {
					d.dice_count = int(i)
				} else {
					return errors.New(fmt.Sprintf("failed parsing dice count for %s in %s", left[4:], src2))
				}
			}
		}
	}
	return nil
}

func (d Damage) String() string {
	return fmt.Sprintf("%s%s%s", d.Prefix(), d.DieCode(), d.Suffix())
}

func (d Damage) Prefix() string {
	if d.base_stat != "" {
		return d.base_stat + "+"
	} else {
		return ""
	}
}

func (d Damage) Suffix() string {
	if d.damage_bonus > 0 {
		return fmt.Sprintf("+%d", d.damage_bonus)
	} else if d.damage_bonus < 0 {
		return fmt.Sprintf("%d", d.damage_bonus)
	} else {
		return ""
	}
}

func (d Damage) DieCode() string {
	if d.ShotgunDamage() {
		return fmt.Sprintf("%d-%dd%d", d.shotgun_min, d.shotgun_max, d.dice_sides)
	} else if d.dice_count != 1 {
		return fmt.Sprintf("%dd%d", d.dice_count, d.dice_sides)
	} else {
		return fmt.Sprintf("d%d", d.dice_sides)
	}
}
