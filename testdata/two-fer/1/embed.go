package twofer

import (
	// Needed to embed the expected files for tests.
	_ "embed"
)

// Mapping contains the expected mapping.
//go:embed expected_mapping.json
var Mapping []byte

// Representation contains the expected representation.
//go:embed expected_representation.txt
var Representation []byte
