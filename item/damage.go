package item

import "fmt"

type Damage struct {
	dice_count   int
	dice_sides   int
	damage_bonus int
	base_stat    string
}

func (d Damage) Max() int { return d.dice_sides*d.dice_count + d.damage_bonus }

func (d Damage) String() string {
	if d.dice_count == 0 || d.dice_sides == 0 {
		return ""
	}
	if d.dice_count != 1 {
		if d.base_stat != "" {
			if d.damage_bonus != 0 {
				return fmt.Sprintf("%s+%dd%d+%d", d.base_stat, d.dice_count, d.dice_sides, d.damage_bonus)
			} else {
				return fmt.Sprintf("%s+%dd%d", d.base_stat, d.dice_count, d.dice_sides)
			}
		} else {
			if d.damage_bonus != 0 {
				return fmt.Sprintf("%dd%d+%d", d.dice_count, d.dice_sides, d.damage_bonus)
			} else {
				return fmt.Sprintf("%dd%d", d.dice_count, d.dice_sides)
			}
		}
	}else{
		if d.base_stat != "" {
			if d.damage_bonus != 0 {
				return fmt.Sprintf("%s+d%d+%d", d.base_stat, d.dice_sides, d.damage_bonus)
			} else {
				return fmt.Sprintf("%s+d%d", d.base_stat, d.dice_sides)
			}
		} else {
			if d.damage_bonus != 0 {
				return fmt.Sprintf("d%d+%d", d.dice_sides, d.damage_bonus)
			} else {
				return fmt.Sprintf("d%d", d.dice_sides)
			}
		}
	}
}

func ParseDamage(source string) (Damage, error) {
	return Damage{dice_count: 2, dice_sides: 6, damage_bonus: 0, base_stat: ""}, nil
}