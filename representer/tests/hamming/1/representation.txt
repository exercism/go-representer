package hamming

import "errors"

func PLACEHOLDER_1(PLACEHOLDER_2, PLACEHOLDER_3 string) (int, error) {
	if len(PLACEHOLDER_2) != len(PLACEHOLDER_3) {
		return 0, errors.New("strings have different length")
	}
	PLACEHOLDER_4 := 0
	for PLACEHOLDER_5 := 0; PLACEHOLDER_5 < len(PLACEHOLDER_2); PLACEHOLDER_5++ {
		if PLACEHOLDER_2[PLACEHOLDER_5] != PLACEHOLDER_3[PLACEHOLDER_5] {
			PLACEHOLDER_4++
		}
	}
	return PLACEHOLDER_4, nil
}
