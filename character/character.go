package character

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

type Character struct {
	Name string
	Ancestry Ancestry
	Strength int
	Dexterity int
	Constitution int
	Intelligence int
	Wisdom int
	Charisma int
	Background Background
	Class Class
	ClassFeatures []ClassFeature

}

// Create walks the user through creating a new character, by asking the user  to provide information about the
// character they would like to create
func (c *Character) Create() {

}
