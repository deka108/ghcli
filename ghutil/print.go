package ghutil

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(data interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	fmt.Printf("%s", jsonData)
	return nil
}
