package representer_test

import (
	"strings"
	"testing"

	"github.com/exercism/go-representer/representer"
	hamming1 "github.com/exercism/go-representer/representer/tests/hamming/1"
	hamming2 "github.com/exercism/go-representer/representer/tests/hamming/2"
	hamming3 "github.com/exercism/go-representer/representer/tests/hamming/3"
	hamming4 "github.com/exercism/go-representer/representer/tests/hamming/4"
	raindrops1 "github.com/exercism/go-representer/representer/tests/raindrops/1"
	raindrops2 "github.com/exercism/go-representer/representer/tests/raindrops/2"
	raindrops3 "github.com/exercism/go-representer/representer/tests/raindrops/3"
	raindrops4 "github.com/exercism/go-representer/representer/tests/raindrops/4"
	raindrops5 "github.com/exercism/go-representer/representer/tests/raindrops/5"
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
		{
			path:        "tests/hamming/2",
			wantRepr:    hamming2.Representation,
			wantMapping: hamming2.Mapping,
		},
		{
			path:        "tests/hamming/3",
			wantRepr:    hamming3.Representation,
			wantMapping: hamming3.Mapping,
		},
		{
			path:        "tests/hamming/4",
			wantRepr:    hamming4.Representation,
			wantMapping: hamming4.Mapping,
		},
		{
			path:        "tests/raindrops/1",
			wantRepr:    raindrops1.Representation,
			wantMapping: raindrops1.Mapping,
		},
		{
			path:        "tests/raindrops/2",
			wantRepr:    raindrops2.Representation,
			wantMapping: raindrops2.Mapping,
		},
		{
			path:        "tests/raindrops/3",
			wantRepr:    raindrops3.Representation,
			wantMapping: raindrops3.Mapping,
		},
		{
			path:        "tests/raindrops/4",
			wantRepr:    raindrops4.Representation,
			wantMapping: raindrops4.Mapping,
		},
		{
			path:        "tests/raindrops/5",
			wantRepr:    raindrops5.Representation,
			wantMapping: raindrops5.Mapping,
		},
	}
	for _, tt := range tests {
		t.Run(strings.TrimPrefix(tt.path, "tests/"), func(t *testing.T) {
			asrt := is.New(t)
			got, err := representer.Extract(tt.path)
			asrt.NoErr(err)

			repr, err := got.RepresentationBytes()
			asrt.NoErr(err)
			assert.Equal(t, string(tt.wantRepr), string(repr))

			mapping, err := got.MappingBytes()
			asrt.NoErr(err)
			assert.Equal(t, string(tt.wantMapping), string(mapping))
		})
	}
}
