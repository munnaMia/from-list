package utils

import (
	"log"
	"os"
	"path/filepath"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Create file and directory if not exist.
func CreateIfNotExist(path string) {
	absFilePath, err := filepath.Abs(path)
	Check(err)

	// Creating a directory if not exist
	dir := filepath.Dir(absFilePath)

	_, err = os.Stat(dir)
	Check(err)

	if os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		Check(err)
	}

	//create file if not exist
	_, err = os.Stat(absFilePath)

	if os.IsNotExist(err) {
		file, err := os.Create(absFilePath)
		Check(err)
		defer file.Close()
		_, err = file.Write([]byte("[]"))
		Check(err)
	}

}
