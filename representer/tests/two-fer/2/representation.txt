package twofer

import "fmt"

func PLACEHOLDER_1(PLACEHOLDER_2 string) string {
	if PLACEHOLDER_2 == "" {
		PLACEHOLDER_2 = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", PLACEHOLDER_2)
}
