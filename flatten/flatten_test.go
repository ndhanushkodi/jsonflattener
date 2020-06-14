package flatten

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestFlatten(t *testing.T) {
	type testcase struct {
		name     string
		input    io.Reader
		expected map[string]interface{}
	}

	tests := []testcase{

		{
			name: "Empty JSON",
			input: strings.NewReader(
				`{}`,
			),
			expected: map[string]interface{}{},
		},

		{
			name: "Non-nested JSON",
			input: strings.NewReader(
				`{
					"hi": "", 
					"foo": "bar",
					"baz": true,
					"boom": 1
				}`,
			),
			// json.Unmarshal stores numbers in json as float64 in a golang object
			// https://golang.org/pkg/encoding/json/#Unmarshal
			expected: map[string]interface{}{
				"hi":   "",
				"foo":  "bar",
				"baz":  true,
				"boom": float64(1),
			},
		},

		{
			name: "Nested JSON",
			input: strings.NewReader(
				`{
					"baz": 2, 
					"nested": {
						"a": "b",
						"c": {
							"d": "e",
							"f": 1.234,
							"g": {
								"h": {
									"i": "j"
								}
							}
						}
					}
				}`,
			),
			expected: map[string]interface{}{
				"baz":            float64(2),
				"nested.a":       "b",
				"nested.c.d":     "e",
				"nested.c.f":     float64(1.234),
				"nested.c.g.h.i": "j",
			},
		},

		{
			name: "Nested JSON with empty value",
			input: strings.NewReader(
				`{
					"a": 1,
					"b": true,
					"c": {}
				}`,
			),
			expected: map[string]interface{}{
				"a": float64(1),
				"b": true,
				"c": map[string]interface{}{},
			},
		},
	}

	for _, tc := range tests {
		actual, err := FlattenJSON(tc.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(tc.expected, actual) {
			t.Errorf(`Testcase: %s
            Expected: %v
            Actual:   %v
            `, tc.name, tc.expected, actual)
		}
	}

}
