package readfile

import "os"

func ReadFile(fileName string) (string, error) {
	fileContent, err := os.ReadFile(fileName)

	return string(fileContent), err
}
