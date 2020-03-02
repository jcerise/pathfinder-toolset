package character

import (
	"PathfinderToolset/cli_utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)


type Heritage struct {
	Name string
	Ancestry Ancestry
	Description string
	Ability string
}

func (h *Heritage) Print() {
	fmt.Printf("%-5s\n", h.Name)
	fmt.Println("----------------------------------------")
	fmt.Println(h.Description)
	fmt.Println()
	fmt.Println(h.Ability)
}

type Ancestry struct {
	Name string `json:"name"`
	Description string `json:"description"`
	HitPoints int `json:"hitpoints"`
	Size string `json:"size"`
	Speed int `json:"speed"`
	Boosts []string `json:"boosts"`
	Flaws []string `json:"flaws"`
	Languages []string `json:"languages"`
	Features map[string]string `json:"features"`
	Heritages []Heritage `json:"heritages"`
	Feats []Feat `json:"feats"`
}

func (a *Ancestry) Print() {
	fmt.Printf("%-5s | HP: %-5d | Size: %-5s | Speed: %-5d\n", a.Name, a.HitPoints, a.Size, a.Speed)
	fmt.Println("----------------------------------------")
	fmt.Println(a.Description)
	fmt.Println()
	fmt.Printf("Boosts:    %s\n", a.Boosts)
	fmt.Printf("Flaws:     %s\n", a.Flaws)
	fmt.Printf("Languages: %s\n", a.Languages)
	fmt.Println()
	fmt.Printf("a Features:\n")
	for featureName, featureDesc := range a.Features {
		fmt.Printf("     %s: %s\n", featureName, featureDesc)
	}
}

// ApplyBoostsFlaws will take any boosts listed for the ancestry, and add a +2 bonus to
// a characters core abilities. The same will be done for for flaws, instead applying a -2
// detraction. If a boost is listed as `free`, the user will be prompted to apply to the boost
// to one of the six core abilities.
func (a *Ancestry) ApplyBoostsFlaws(character Character) {
	abilities := []string{"strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}
	// First, apply boosts
	for _, boost := range a.Boosts {
		if boost == "free" {
			// Assume anything else is a free boost. Let the user decide what they want to boost
			choice := cli_utils.ListChoice("Choose an ability to boost:", abilities)
			character.BoostAbility(abilities[choice - 1])
		} else {
			character.BoostAbility(boost)
		}
	}

	// Next, apply flaws
	for _, flaw := range a.Flaws {
		if flaw == "free" {
			// Assume anything else is a free boost. Let the user decide what they want to boost
			choice := cli_utils.ListChoice("Choose an ability to flaw:", abilities)
			character.FlawAbility(abilities[choice - 1])
		} else {
			character.FlawAbility(flaw)
		}
	}
}

func (a *Ancestry) heritageList() []string {
	var heritageNames []string

	for _, item := range a.Heritages {
		heritageNames = append(heritageNames, item.Name)
	}

	return heritageNames
}

func (a *Ancestry) featList() []string {
	var featNames []string

	for _, item := range a.Feats {
		featNames = append(featNames, item.Name)
	}

	return featNames
}

func (a *Ancestry) HeritageSelector() Heritage {
	heritageChosen := false
	var heritage Heritage

	for !heritageChosen {
		heritageChoice := cli_utils.ListChoice("Choose a Heritage: ", a.heritageList())
		curHeritage := a.Heritages[heritageChoice-1]
		curHeritage.Print()
		fmt.Println()
		decision := cli_utils.Confirmation("Would you like to choose this Heritage?", true)

		if decision {
			heritageChosen = true
			heritage = curHeritage
		} else {
			heritageChosen = false
		}
	}

	return heritage
}

func (a *Ancestry) FeatSelector() Feat {
	featChosen := false
	var feat Feat

	for !featChosen {
		featChoice := cli_utils.ListChoice("Choose an Ancestry Feat: ", a.featList())
		curfeat := a.Feats[featChoice-1]
		curfeat.Print()
		fmt.Println()
		decision := cli_utils.Confirmation("Would you like to choose this Ancestry Feat?", true)

		if decision {
			featChosen = true
			feat = curfeat
		} else {
			featChosen = false
		}
	}

	return feat
}

// ancestryList returns a list of all available ancestries to choose from when creating a character, or looking up
// information. It scans a directory called `ancestries` located at the the root of the program, and compiles a list
// of file names. Each ancestry file should be named according to the ancestry name, and have a .ancestry extension.
func AncestryList(ancestryDir string) []string {
	var ancestryNames []string

	root := ancestryDir
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// We only care about .ancestry files. All others in the directory will be ignored
		if filepath.Ext(path) == ".ancestry" {
			ancestryName := strings.Split(info.Name(), ".")[0]
			ancestryNames = append(ancestryNames, ancestryName)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	return ancestryNames
}

func GetAncestryInfo(ancestryName, ancestryDir string) Ancestry {
	fileName := ancestryDir + "/" + ancestryName + ".ancestry"
	ancestryFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	ancestryData := Ancestry{}
	err = json.Unmarshal(ancestryFile, &ancestryData)

	if err != nil {
		log.Fatal(err)
	}

	return ancestryData
}

func AncestrySelector() Ancestry {
	ancestryChosen := false
	var ancestry Ancestry
	for !ancestryChosen {
		ancestries := AncestryList("ancestries")
		ancestryChoice := cli_utils.ListChoice("Choose your characters ancestry: ", ancestries)

		curAncestry := GetAncestryInfo(ancestries[ancestryChoice - 1], "ancestries")
		curAncestry.Print()
		fmt.Println()
		decision := cli_utils.Confirmation("Would you like to choose this ancestry?", true)

		if decision {
			ancestryChosen = true
			ancestry = curAncestry
		} else {
			ancestryChosen = false
		}
	}
	return ancestry
}
