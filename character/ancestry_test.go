package character

import "testing"

func TestAncestryList(t *testing.T) {
	// Test fetching the list of Ancestries from the directory. This should return a list of names.
	ancestry_names := AncestryList("../ancestries")
	if len(ancestry_names) != 9 {
		t.Errorf("incorrect number of ancestries fetched: got %d, wanted %d.", len(ancestry_names), 9)
	}

	// Check a couple of known, core, ancestries that should exist
	if !contains("Human", ancestry_names) {
		t.Errorf("ancestries did not contain expected ancestry: wanted Human")
	}

	if !contains("Gnome", ancestry_names) {
		t.Errorf("ancestries did not contain expected ancestry: wanted Gnome")
	}

	if !contains("Elf", ancestry_names) {
		t.Errorf("ancestries did not contain expected ancestry: wanted Elf")
	}
}

func TestPrintAncestryInfo(t *testing.T) {
	ancestryDir := "../ancestries"
	ancestryName := "Dwarf"

	dwarfAncestry := GetAncestryInfo(ancestryName, ancestryDir)

	if dwarfAncestry.Name != "Dwarf" {
		t.Errorf("ancestry did not properly load: wanted name=%s, got name=%s", "Dwarf", dwarfAncestry.Name)
	}

	if len(dwarfAncestry.Heritages) != 5 {
		t.Errorf("heritages not properly populated: wanted %d, got %d", 5, len(dwarfAncestry.Heritages))
	}
}

// contains checks a given slice to see if contains the specified key. Returns true if key is found,
// false otherwise
func contains(key string, arr []string) bool {
	for i := range arr {
		if arr[i] == key {
			return true
		}
	}
	return false
}

