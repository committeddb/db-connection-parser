package parser_test

import (
	"strings"
	"testing"

	"github.com/committeddb/db-connection-parser/pkg/parser"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  map[string]parser.Datastore
	}{
		"empty": {input: "", want: nil},
		"simple": {
			input: "datastores:\n  -\n    name: foo\n    url: bar",
			want:  map[string]parser.Datastore{"foo": &parser.NullDatastore{}},
		},
		"two": {
			input: "datastores:\n  -\n    name:   foo\n    url: bar\n  -\n    name: baz\n    url: qux\n",
			want: map[string]parser.Datastore{
				"foo": &parser.NullDatastore{},
				"baz": &parser.NullDatastore{},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parser.Parse(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
