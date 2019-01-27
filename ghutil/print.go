package ghutil

import (
	"encoding/json"
	"fmt"
)

// PrettyPrint prints a struct in human-readable way
func PrettyPrint(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("%s", jsonData)
	return nil
}
