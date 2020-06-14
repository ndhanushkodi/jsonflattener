package flatten

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// FlattenJSON takes input from any type that implements the io.Reader
// interface, and returns a flattened JSON object.
func FlattenJSON(input io.Reader) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to read: %v", err)
	}

	jsonObject := make(map[string]interface{})
	err = json.Unmarshal(bytes, &jsonObject)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal JSON: %v", err)
	}

	output := make(map[string]interface{})
	flatten(jsonObject, "", output)

	return output, nil
}

// flatten is first called with the original JSON object and recursively adds
// fully flattened keys to the output JSON object.
func flatten(input interface{}, flattenedKey string, output map[string]interface{}) {
	switch input.(type) {
	default: // fully flattened key value pair (base case)
		output[flattenedKey] = input
	case map[string]interface{}: // JSON object
		// handles including keys with an empty object value
		// otherwise, flatten gets called with an empty value and the pair is omitted
		// i.e {"key": {}}
		if len(input.(map[string]interface{})) == 0 && flattenedKey != "" {
			output[flattenedKey] = input
			return
		}
		// for json objects with 1+ key value pairs
		// i.e {"key": {...}}
		for k, v := range input.(map[string]interface{}) {
			flatten(v, constructFlattenedKey(flattenedKey, k), output)
		}
	}
}

func constructFlattenedKey(name, k string) string {
	if name == "" {
		return k
	}
	return fmt.Sprintf("%s.%s", name, k)
}
