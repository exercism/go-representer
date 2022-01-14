package raindrops

import (
	"strconv"
	"strings"
)

func PLACEHOLDER_1(PLACEHOLDER_2 int) string {
	PLACEHOLDER_3 := strings.Builder{}
	if PLACEHOLDER_2%3 == 0 {
		PLACEHOLDER_3.WriteString("Pling")
	}
	if PLACEHOLDER_2%5 == 0 {
		PLACEHOLDER_3.WriteString("Plang")
	}
	if PLACEHOLDER_2%7 == 0 {
		PLACEHOLDER_3.WriteString("Plong")
	}
	if PLACEHOLDER_3.Len() == 0 {
		PLACEHOLDER_3.WriteString(strconv.Itoa(PLACEHOLDER_2))
	}
	return PLACEHOLDER_3.String()
}
