package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteJsonFile(jsonObject interface{}, path string) error {
	b, err := json.Marshal(jsonObject)
	if err != nil {
		return err
	}

	return WriteStringFile(b, path)
}

func WriteStringFile(data []byte, path string) error {
	os.Mkdir(filepath.Dir(path), 0755)

	// write to file
	err := ioutil.WriteFile(path, data, 0644)

	if err != nil {
		return err
	}
	return nil
}
