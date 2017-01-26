package library

import (
	"showdown/identity"
	"showdown/character"
	"github.com/ghodss/yaml"
)

type EdgeLibrary map[identity.ID]character.Edge

var edge_library EdgeLibrary = make(EdgeLibrary)

func loadEdgeLibrary() {
	src := `
showdown.2009.edge.ace:
  id: showdown.2009.edge.ace
  name: Ace
showdown.2009.edge.acrobat:
  id: showdown.2009.edge.acrobat
  name: Acrobat
showdown.2009.edge.alertness:
  id: showdown.2009.edge.alertness
  name: Alertness
showdown.2009.edge.ambidextrous:
  id: showdown.2009.edge.ambidextrous
  name: Ambidextrous
showdown.2009.edge.arcane_background:
  id: showdown.2009.edge.arcane_background
  name: Arcane Background
showdown.2009.edge.arcane_resistance:
  id: showdown.2009.edge.arcane_resistance
  name: Arcane Resistance
showdown.2009.edge.beast_bond:
  id: showdown.2009.edge.beast_bond
  name: Beast Bond
showdown.2009.edge.beast_master:
  id: showdown.2009.edge.beast_master
  name: Beast Master
showdown.2009.edge.berserk:
  id: showdown.2009.edge.berserk
  name: Berserk
showdown.2009.edge.block:
  id: showdown.2009.edge.block
  name: Block
showdown.2009.edge.brawny:
  id: showdown.2009.edge.brawny
  name: Brawny
showdown.2009.edge.champion:
  id: showdown.2009.edge.champion
  name: Champion
showdown.2009.edge.combat_reflexes_:
  id: showdown.2009.edge.combat_reflexes_
  name: Combat Reflexes
showdown.2009.edge.common_bond_:
  id: showdown.2009.edge.common_bond_
  name: Common Bond
showdown.2009.edge.danger_sense:
  id: showdown.2009.edge.danger_sense
  name: Danger Sense
showdown.2009.edge.dead_shot:
  id: showdown.2009.edge.dead_shot
  name: Dead Shot
showdown.2009.edge.dodge:
  id: showdown.2009.edge.dodge
  name: Dodge
showdown.2009.edge.fervor:
  id: showdown.2009.edge.fervor
  name: Fervor
showdown.2009.edge.first_strike:
  id: showdown.2009.edge.first_strike
  name: First Strike
showdown.2009.edge.fleet-footed:
  id: showdown.2009.edge.fleet-footed
  name: Fleet-Footed
showdown.2009.edge.florentine:
  id: showdown.2009.edge.florentine
  name: Florentine
showdown.2009.edge.frenzy:
  id: showdown.2009.edge.frenzy
  name: Frenzy
showdown.2009.edge.giant_killer:
  id: showdown.2009.edge.giant_killer
  name: Giant Killer
showdown.2009.edge.great_luck:
  id: showdown.2009.edge.great_luck
  name: Great Luck
showdown.2009.edge.hard_to_kill:
  id: showdown.2009.edge.hard_to_kill
  name: Hard to Kill
showdown.2009.edge.harder_to_kill:
  id: showdown.2009.edge.harder_to_kill
  name: Harder to Kill
showdown.2009.edge.healer:
  id: showdown.2009.edge.healer
  name: Healer
showdown.2009.edge.hold_the_line!:
  id: showdown.2009.edge.hold_the_line!
  name: Hold the Line!
showdown.2009.edge.holy/unholy_warrior:
  id: showdown.2009.edge.holy/unholy_warrior
  name: Holy/Unholy Warrior
showdown.2009.edge.improved_arcane_res_:
  id: showdown.2009.edge.improved_arcane_res_
  name: Improved Arcane Res
showdown.2009.edge.improved_block::
  id: showdown.2009.edge.improved_block:
  name: Improved Block:
showdown.2009.edge.improved_first-strike:
  id: showdown.2009.edge.improved_first-strike
  name: Improved First-Strike
showdown.2009.edge.improved-frenzy:
  id: showdown.2009.edge.improved-frenzy
  name: Improved-Frenzy
showdown.2009.edge.improved_level_headed::
  id: showdown.2009.edge.improved_level_headed:
  name: Improved Level Headed:
showdown.2009.edge.improved_nerves_of_steel::
  id: showdown.2009.edge.improved_nerves_of_steel:
  name: Improved Nerves of Steel:
showdown.2009.edge.improved_sweep::
  id: showdown.2009.edge.improved_sweep:
  name: Improved Sweep:
showdown.2009.edge.improved_tough_as_nails::
  id: showdown.2009.edge.improved_tough_as_nails:
  name: Improved Tough as Nails:
showdown.2009.edge.improved_trademark_weapon:
  id: showdown.2009.edge.improved_trademark_weapon
  name: Improved Trademark Weapon
showdown.2009.edge.improvedoved_dodge:
  id: showdown.2009.edge.improvedoved_dodge
  name: Improvedoved Dodge
showdown.2009.edge.inspire::
  id: showdown.2009.edge.inspire:
  name: Inspire:
showdown.2009.edge.jack-of-all-trades:
  id: showdown.2009.edge.jack-of-all-trades
  name: Jack-of-all-Trades
showdown.2009.edge.level_headed:
  id: showdown.2009.edge.level_headed
  name: Level Headed
showdown.2009.edge.luck:
  id: showdown.2009.edge.luck
  name: Luck
showdown.2009.edge.marksman:
  id: showdown.2009.edge.marksman
  name: Marksman
showdown.2009.edge.master_of_arms:
  id: showdown.2009.edge.master_of_arms
  name: Master of Arms
showdown.2009.edge.mentalist:
  id: showdown.2009.edge.mentalist
  name: Mentalist
showdown.2009.edge.mighty_blow:
  id: showdown.2009.edge.mighty_blow
  name: Mighty Blow
showdown.2009.edge.mr._fix_it:
  id: showdown.2009.edge.mr._fix_it
  name: Mr. Fix It
showdown.2009.edge.natural_leader:
  id: showdown.2009.edge.natural_leader
  name: Natural Leader
showdown.2009.edge.nerves_of_steel:
  id: showdown.2009.edge.nerves_of_steel
  name: Nerves of Steel
showdown.2009.edge.nerves_of_steel:
  id: showdown.2009.edge.nerves_of_steel
  name: Nerves of Steel
showdown.2009.edge.no_mercy:
  id: showdown.2009.edge.no_mercy
  name: No Mercy
showdown.2009.edge.occult:
  id: showdown.2009.edge.occult
  name: Occult
showdown.2009.edge.power_surge:
  id: showdown.2009.edge.power_surge
  name: Power Surge
showdown.2009.edge.quick*:
  id: showdown.2009.edge.quick*
  name: Quick*
showdown.2009.edge.rock_and_roll!:
  id: showdown.2009.edge.rock_and_roll!
  name: Rock and Roll!
showdown.2009.edge.soul_drain:
  id: showdown.2009.edge.soul_drain
  name: Soul Drain
showdown.2009.edge.steady_hands:
  id: showdown.2009.edge.steady_hands
  name: Steady Hands
showdown.2009.edge.strong_willed:
  id: showdown.2009.edge.strong_willed
  name: Strong Willed
showdown.2009.edge.sweep:
  id: showdown.2009.edge.sweep
  name: Sweep
showdown.2009.edge.tough_as_nails:
  id: showdown.2009.edge.tough_as_nails
  name: Tough as Nails
showdown.2009.edge.trademark_weapon:
  id: showdown.2009.edge.trademark_weapon
  name: Trademark Weapon
showdown.2009.edge.two-fisted:
  id: showdown.2009.edge.two-fisted
  name: Two-Fisted
showdown.2009.edge.weapon_master:
  id: showdown.2009.edge.weapon_master
  name: Weapon Master
showdown.2009.edge.wizard:
  id: showdown.2009.edge.wizard
  name: Wizard
showdown.2009.edge.woodsman:
  id: showdown.2009.edge.woodsman
  name: Woodsman
`

	if err := yaml.Unmarshal([]byte(src), &edge_library); err != nil {
		panic(err)
	}
}
