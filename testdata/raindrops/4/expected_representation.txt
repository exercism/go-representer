package raindrops

import (
	"strconv"
)

func PLACEHOLDER_1(PLACEHOLDER_2 int) string {
	PLACEHOLDER_3 := ""
	if PLACEHOLDER_4(PLACEHOLDER_2, 3) {
		PLACEHOLDER_3 += "Pling"
	}
	if PLACEHOLDER_4(PLACEHOLDER_2, 5) {
		PLACEHOLDER_3 += "Plang"
	}
	if PLACEHOLDER_4(PLACEHOLDER_2, 7) {
		PLACEHOLDER_3 += "Plong"
	}
	if PLACEHOLDER_3 == "" {
		return strconv.Itoa(PLACEHOLDER_2)
	}
	return PLACEHOLDER_3
}
func PLACEHOLDER_4(PLACEHOLDER_5, PLACEHOLDER_6 int) bool {
	return PLACEHOLDER_5%PLACEHOLDER_6 == 0
}
