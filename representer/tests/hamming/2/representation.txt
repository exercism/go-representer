package hamming

import "errors"

func PLACEHOLDER_1(PLACEHOLDER_2, PLACEHOLDER_3 string) (int, error) {
	if len(PLACEHOLDER_2) != len(PLACEHOLDER_3) {
		return 0, errors.New("strings have different length")
	}
	PLACEHOLDER_4 := 0
	PLACEHOLDER_5 := []rune(PLACEHOLDER_3)
	for PLACEHOLDER_6, PLACEHOLDER_7 := range PLACEHOLDER_2 {
		if PLACEHOLDER_7 != PLACEHOLDER_5[PLACEHOLDER_6] {
			PLACEHOLDER_4++
		}
	}
	return PLACEHOLDER_4, nil
}
