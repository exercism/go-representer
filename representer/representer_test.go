package representer_test

import (
	"strings"
	"testing"

	"github.com/exercism/go-representer/representer"
	hamming1 "github.com/exercism/go-representer/representer/tests/hamming/1"
	twofer1 "github.com/exercism/go-representer/representer/tests/two-fer/1"
	twofer2 "github.com/exercism/go-representer/representer/tests/two-fer/2"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	tests := []struct {
		path        string
		wantRepr    []byte
		wantMapping []byte
	}{
		{
			path:        "tests/two-fer/1",
			wantRepr:    twofer1.Representation,
			wantMapping: twofer1.Mapping,
		},
		{
			path:        "tests/two-fer/2",
			wantRepr:    twofer2.Representation,
			wantMapping: twofer2.Mapping,
		},
		{
			path:        "tests/hamming/1",
			wantRepr:    hamming1.Representation,
			wantMapping: hamming1.Mapping,
		},
	}
	for _, tt := range tests {
		t.Run(strings.TrimPrefix(tt.path, "tests/"), func(t *testing.T) {
			asrt := is.New(t)
			got, err := representer.Extract(tt.path)
			asrt.NoErr(err)

			repr, err := got.RepresentationBytes()
			asrt.NoErr(err)
			assert.Equal(t, strings.TrimSpace(string(tt.wantRepr)), string(repr))

			mapping, err := got.MappingBytes()
			asrt.NoErr(err)
			assert.Equal(t, strings.TrimSpace(string(tt.wantMapping)), string(mapping))
		})
	}
}
