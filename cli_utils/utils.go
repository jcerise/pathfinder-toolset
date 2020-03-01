package cli_utils

import (
	"fmt"
	"log"
	"strings"
)


// ChooseFromList prints a list of choices to the terminal, and prompts the user to select a valid
// numeric choice from those provided. If an invalid choice is provided, the user is notified, and
// asked to choose a valid choice. This will repeat until a valid choice is selected. The list index
// of the valid choice selected is returned.
func ListChoice(text string, iterable []string) int {
	fmt.Println(text)

	// Print the list provided, with a numeric index for each item.
	for i := range iterable {
		fmt.Printf("%d. %s\n", i + 1, iterable[i])
	}
	fmt.Print("> ")

	var choice int
	_, err := fmt.Scanf("%d", &choice)

	if err != nil {
		log.Fatal(err)
	}

	for choice > len(iterable) {
		fmt.Println("Please make a valid selection.")
		fmt.Print("> ")
		_, err = fmt.Scanf("%d", &choice)

		if err != nil {
			log.Fatal(err)
		}
	}

	return choice
}

// Confirmation prints a confirmation message to the screen. It accepts a default (either y or n), and will provide
// that value if nothing is input. The choice is returned.
func Confirmation(text string, defaultChoice bool) bool {
	if defaultChoice {
		fmt.Printf("%s (Y/n)", text)
	} else {
		fmt.Printf("%s (y/N)", text)
	}

	var confChoice string
	_, err := fmt.Scanf("%s", &confChoice)

	if err != nil {
		if err.Error() == "unexpected newline" {
			return defaultChoice
		} else {
			log.Fatal(err)
		}
	}

	validChoice := false
	var choice string

	for !validChoice {
		choice = strings.ToLower(confChoice)
		if choice == "y" || choice == "yes" || choice == "n" || choice == "no" {
			validChoice = true
		} else {
			if defaultChoice {
				fmt.Println("Please enter (Y)es or (n)o ")
			} else {
				fmt.Println("Please enter (y)es or (N)o ")
			}
			_, _ = fmt.Scan(&confChoice)
		}
	}
	var decision bool
	if choice == "yes" || choice == "y" {
		decision = true
	} else if choice == "no" || choice == "n" {
		decision = false
	}

	return decision
}
