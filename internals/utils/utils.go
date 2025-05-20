package utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/munnaMia/from-list/internals/model"
)

// Create file and directory if not exist.
func CreateIfNotExist(path string) error {
	absFilePath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Creating a directory if not exist
	dir := filepath.Dir(absFilePath)

	_, err = os.Stat(dir)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	//create file if not exist
	_, err = os.Stat(absFilePath)

	if os.IsNotExist(err) {
		file, err := os.Create(absFilePath)
		if err != nil {
			return err
		}

		defer file.Close()
		_, err = file.Write([]byte("[]"))
		if err != nil {
			return err
		}
	}
	return nil // return nil if every thing goes well
}

func WriteJson(formData model.User, storagePath string) error {
	err := CreateIfNotExist(storagePath) // check the file exist or not before writing
	if err != nil {
		return err
	}

	// read database
	records, err := ReadJson(storagePath)
	if err != nil {
		return err
	}

	records = append(records, formData)

	jsonFromData, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		return err
	}

	absFilePath, err := filepath.Abs(storagePath)
	if err != nil {
		return err
	}

	file, err := os.Create(absFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonFromData)
	if err != nil {
		return err
	}
	return nil
}

func ReadJson(storagePath string) ([]model.User, error) {
	var data []model.User

	file, err := os.Open(storagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	readFile, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(readFile, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
