package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ndhanushkodi/jsonflattener/flatten"
)

func main() {
	flattenedObject, err := flatten.FlattenJSON(os.Stdin)
	if err != nil {
		panic(fmt.Sprintf("Failed to flatten: %v", err))
	}

	flattenedJSON, err := json.Marshal(flattenedObject)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal JSON: %v", err))
	}

	fmt.Println(string(flattenedJSON))
}
