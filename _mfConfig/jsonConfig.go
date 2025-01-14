package _mfConfig

import (
	"encoding/json"
	"fmt"
	"os"
)

func decodeJSONConfig(v interface{}, filename string) error {

	fmt.Println("Decoding JSON...")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(v)

}
