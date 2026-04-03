package utils

import "os"

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func EnsureFile(filename string) error {
	if !FileExists(filename) {
		file, err := os.Create(filename)

		if err != nil {
			return err
		}

		defer file.Close()

		file.WriteString("[]")
	}

	return nil
}
