package character

import (
	"fmt"
)

type Background struct {
	Name string
	Description string
	Boosts []string
	Skill []Skill
	Feat Feat
}

type Class struct {
	Name string
	Description string
	KeyAbility string
	HitPoints int
	Perception string
	Fortitude string
	Reflex string
	Will string
	Skills []string
	Attacks []string
	Defenses []string
	DC string
	ClassFeatures []ClassFeature
}

type ClassFeature struct {
	Name string
	Description string
	GrantedFeat Feat
	Level int
}

type Skill struct {
	Name string
	KeyAbility string
	UntrainedActions []string
	TrainedActions []string
}

type Feat struct {
	Name string
	Level int
	Prerequisite string
	Types []string
	Description string
	ActionType string
}

func (f *Feat) Print() {
	fmt.Printf("%-5s | %-5s | Level %-5d\n", f.Name, f.ActionType, f.Level)
	fmt.Println("----------------------------------------")
	fmt.Printf("Prerequisites: %s", f.Prerequisite)
	fmt.Println()
	fmt.Println(f.Description)
}

type Character struct {
	Name string
	Ancestry Ancestry
	Heritage Heritage
	Strength int
	Dexterity int
	Constitution int
	Intelligence int
	Wisdom int
	Charisma int
	Background Background
	Class Class
	ClassFeatures []ClassFeature
	Feats []Feat

}

func (c *Character) BoostAbility(abilityName string) {
	if abilityName == "strength" {
		c.Strength += 2
	} else if abilityName == "dexterity" {
		c.Dexterity += 2
	} else if abilityName == "constitution" {
		c.Constitution += 2
	} else if abilityName == "intelligence" {
		c.Intelligence += 2
	} else if abilityName == "wisdom" {
		c.Wisdom += 2
	} else if abilityName == "charisma" {
		c.Charisma += 2
	}
	fmt.Printf("Boosted %s by 2 points\n", abilityName)
}

func (c *Character) FlawAbility(abilityName string) {
	if abilityName == "strength" {
		c.Strength -= 2
	} else if abilityName == "dexterity" {
		c.Dexterity -= 2
	} else if abilityName == "constitution" {
		c.Constitution -= 2
	} else if abilityName == "intelligence" {
		c.Intelligence -= 2
	} else if abilityName == "wisdom" {
		c.Wisdom -= 2
	} else if abilityName == "charisma" {
		c.Charisma -= 2
	}
	fmt.Printf("Flawed %s by 2 points\n", abilityName)
}

// Create walks the user through creating a new character, by asking the user  to provide information about the
// character they would like to create
func Create() {
	var newCharacter Character
	ancestry := AncestrySelector()

	newCharacter.Ancestry = ancestry

	// Apply boosts and flaws
	ancestry.ApplyBoostsFlaws(newCharacter)

	// Choose a heritage
	heritage := ancestry.HeritageSelector()
	newCharacter.Heritage = heritage

	// Choose Ancestry feat
	feat := ancestry.FeatSelector()
	newCharacter.Feats = append(newCharacter.Feats, feat)
}
