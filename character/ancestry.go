package character

import (
	"os"
	"path/filepath"
	"strings"
)

type Ancestry struct {
	Name string
	HitPoints int
	Size string
	Speed int
	Boosts []string
	Flaws []string
	Languages []string
	Features []string
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
