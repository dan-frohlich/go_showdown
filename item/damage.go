package item

import "fmt"

type Damage struct {
	dice_count   int
	dice_sides   int
	damage_bonus int
	base_stat    string
	shotgun_min  int
	shotgun_max  int
}

func (d Damage) Max() int { return d.dice_sides*d.dice_count + d.damage_bonus }

func (d Damage) ShotgunDamage() bool { return d.shotgun_min != 0 && d.shotgun_max != 0 }

func ParseDamage(source string) (Damage, error) {
	return Damage{dice_count: 2, dice_sides: 6, damage_bonus: 0, base_stat: ""}, nil
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
