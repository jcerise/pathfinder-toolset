package character

import (
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

func PrintAncestryInfo(ancestry Ancestry) {
	fmt.Printf("%-5s | HP: %-5d | Size: %-5s | Speed: %-5d\n", ancestry.Name, ancestry.HitPoints, ancestry.Size, ancestry.Speed)
	fmt.Println("----------------------------------------")
	fmt.Println(ancestry.Description)
	fmt.Println()
	fmt.Printf("Boosts:    %s\n", ancestry.Boosts)
	fmt.Printf("Flaws:     %s\n", ancestry.Flaws)
	fmt.Printf("Languages: %s\n", ancestry.Languages)
	fmt.Println()
	fmt.Printf("Ancestry Features:\n")
	for featureName, featureDesc := range ancestry.Features {
		fmt.Printf("     %s: %s\n", featureName, featureDesc)
	}

}
