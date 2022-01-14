package hamming

import (
	// Needed to embed the expected files for tests.
	_ "embed"
)

// Mapping contains the expected mapping.
//go:embed mapping.json
var Mapping []byte

// Representation contains the expected representation.
//go:embed representation.txt
var Representation []byte
