package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const template = `{ 
	"method": "GET",
	"url": "http://localhost:<PORT>"
}`

func AddAPI(name string, path string) (message string, success bool) {
	message, success = validate("API", name, path)

	if !success {
		return message, success
	}

	file := fmt.Sprintf("%s%s", name, ".json")

	err := os.WriteFile(filepath.Join(path, file), []byte(template), os.ModePerm)

	if err != nil {
		return "Failed To Add API", false
	}

	return message, success
}

func AddCollection(name string, path string) (message string, success bool) {
	message, success = validate("COLLECTION", name, path)

	if !success {
		return message, success
	}

	err := os.Mkdir(filepath.Join(path, name), os.ModePerm)

	if err != nil {
		return "Failed To Add Collection", false
	}

	return message, success
}

func validate(kind string, name string, path string) (message string, success bool) {
	entries, err := os.ReadDir(path)

	if err != nil {
		return "Invalid Path Provided", false
	}

	for _, entry := range entries {
		if kind == "COLLECTION" && entry.IsDir() && entry.Name() == name {
			return "Collection Already Exists", false
		}

		file := fmt.Sprintf("%s%s", name, ".json")

		if strings.Contains(entry.Name(), file) {
			return "API Already Exists", false
		}
	}

	if kind == "COLLECTION" {
		return "Collection Added Successfully", true
	}

	return "API Added Successfully", true
}
