package utils

import (
	"os"
	"path/filepath"
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
