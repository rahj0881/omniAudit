// Package tutil provides test utilities.
//
//nolint:gochecknoglobals,gomnd,gosec // These not an issue for simple testing library.
package tutil

import (
	"encoding/json"
	"flag"
	"os"
	"path"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

var (
	update = flag.Bool("update", false, "Create or update golden files, instead of comparing them")
	clean  = flag.Bool("clean", false, "Deletes the testdata folder before updating (noop of update==false)")
)

var cleanOnce sync.Once

// WithFilename configures a custom golden test filename.
func WithFilename(name string) func(*string) {
	return func(filename *string) {
		*filename = name
	}
}

// RequireGoldenBytes asserts that a golden testdata file exists containing the exact data.
// This is heavily inspired from https://github.com/sebdah/goldie.
func RequireGoldenBytes(t *testing.T, data []byte, opts ...func(*string)) {
	t.Helper()

	filename := strings.ReplaceAll(t.Name(), "/", "_") + ".golden"
	for _, opt := range opts {
		opt(&filename)
	}
	filename = path.Join("testdata", filename)

	if *update {
		if *clean {
			cleanOnce.Do(func() {
				_ = os.RemoveAll("testdata")
			})
		}

		require.NoError(t, os.MkdirAll("testdata", 0o755))

		_ = os.Remove(filename)
		require.NoError(t, os.WriteFile(filename, data, 0o644))

		return
	}

	expected, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		t.Fatalf("golden file does not exist, %s, generate by running with -update", filename)
		return
	}

	require.Equalf(t, string(expected), string(data), "Golden file mismatch, %s", filename)
}

// RequireGoldenJSON asserts that a golden testdata file exists containing the JSON serialized form of the data object.
// This is heavily inspired from https://github.com/sebdah/goldie.
func RequireGoldenJSON(t *testing.T, data any, opts ...func(*string)) {
	t.Helper()

	if _, ok := data.(proto.Message); ok {
		require.Fail(t, "RequireGoldenJSON should not be used with protobuf messages, use RequireGoldenProto")
	}

	b, err := json.MarshalIndent(data, "", " ")
	require.NoError(t, err)

	RequireGoldenBytes(t, b, opts...)
}
