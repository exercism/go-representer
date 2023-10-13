package representer

import (
	"strings"
	"testing"
	"runtime"

	declarations1 "github.com/exercism/go-representer/testdata/declarations/1"
	declarations2 "github.com/exercism/go-representer/testdata/declarations/2"
	declarations3 "github.com/exercism/go-representer/testdata/declarations/3"
	declarations4 "github.com/exercism/go-representer/testdata/declarations/4"
	hamming1 "github.com/exercism/go-representer/testdata/hamming/1"
	hamming2 "github.com/exercism/go-representer/testdata/hamming/2"
	hamming3 "github.com/exercism/go-representer/testdata/hamming/3"
	hamming4 "github.com/exercism/go-representer/testdata/hamming/4"
	raindrops1 "github.com/exercism/go-representer/testdata/raindrops/1"
	raindrops2 "github.com/exercism/go-representer/testdata/raindrops/2"
	raindrops3 "github.com/exercism/go-representer/testdata/raindrops/3"
	raindrops4 "github.com/exercism/go-representer/testdata/raindrops/4"
	raindrops5 "github.com/exercism/go-representer/testdata/raindrops/5"
	raindrops6 "github.com/exercism/go-representer/testdata/raindrops/6"
	raindrops7 "github.com/exercism/go-representer/testdata/raindrops/7"
	twofer1 "github.com/exercism/go-representer/testdata/two-fer/1"
	twofer2 "github.com/exercism/go-representer/testdata/two-fer/2"
	"github.com/matryer/is"
	"github.com/stretchr/testify/assert"
)

func sanitizeResult(s string) string {
	result := s

	if runtime.GOOS == "windows" {
		result = strings.ReplaceAll(result, "\r\n", "\n")
	}

	return result
}

func TestExtract(t *testing.T) {
	tests := []struct {
		path        string
		wantRepr    []byte
		wantMapping []byte
	}{
		{
			path:        "testdata/two-fer/1",
			wantRepr:    twofer1.Representation,
			wantMapping: twofer1.Mapping,
		},
		{
			path:        "testdata/two-fer/2",
			wantRepr:    twofer2.Representation,
			wantMapping: twofer2.Mapping,
		},
		{
			path:        "testdata/hamming/1",
			wantRepr:    hamming1.Representation,
			wantMapping: hamming1.Mapping,
		},
		{
			path:        "testdata/hamming/2",
			wantRepr:    hamming2.Representation,
			wantMapping: hamming2.Mapping,
		},
		{
			path:        "testdata/hamming/3",
			wantRepr:    hamming3.Representation,
			wantMapping: hamming3.Mapping,
		},
		{
			path:        "testdata/hamming/4",
			wantRepr:    hamming4.Representation,
			wantMapping: hamming4.Mapping,
		},
		{
			path:        "testdata/raindrops/1",
			wantRepr:    raindrops1.Representation,
			wantMapping: raindrops1.Mapping,
		},
		{
			path:        "testdata/raindrops/2",
			wantRepr:    raindrops2.Representation,
			wantMapping: raindrops2.Mapping,
		},
		{
			path:        "testdata/raindrops/3",
			wantRepr:    raindrops3.Representation,
			wantMapping: raindrops3.Mapping,
		},
		{
			path:        "testdata/raindrops/4",
			wantRepr:    raindrops4.Representation,
			wantMapping: raindrops4.Mapping,
		},
		{
			path:        "testdata/raindrops/5",
			wantRepr:    raindrops5.Representation,
			wantMapping: raindrops5.Mapping,
		},
		{
			path:        "testdata/raindrops/6",
			wantRepr:    raindrops6.Representation,
			wantMapping: raindrops6.Mapping,
		},
		{
			path:        "testdata/raindrops/7",
			wantRepr:    raindrops7.Representation,
			wantMapping: raindrops7.Mapping,
		},
		{
			path:        "testdata/declarations/1",
			wantRepr:    declarations1.Representation,
			wantMapping: declarations1.Mapping,
		},
		{
			path:        "testdata/declarations/2",
			wantRepr:    declarations2.Representation,
			wantMapping: declarations2.Mapping,
		},
		{
			path:        "testdata/declarations/3",
			wantRepr:    declarations3.Representation,
			wantMapping: declarations3.Mapping,
		},
		{
			path:        "testdata/declarations/4",
			wantRepr:    declarations4.Representation,
			wantMapping: declarations4.Mapping,
		},
	}
	for _, tt := range tests {
		t.Run(strings.TrimPrefix(tt.path, "testdata/"), func(t *testing.T) {
			asrt := is.New(t)
			repr, mapping, err := Extract(tt.path)
			asrt.NoErr(err)

			assert.Equal(t, sanitizeResult(string(tt.wantRepr)), sanitizeResult(string(repr)))
			assert.Equal(t, sanitizeResult(string(tt.wantMapping)), sanitizeResult(string(mapping)))
		})
	}
}
