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

	api := fmt.Sprintf("%s%s", name, ".json")

	path, err := filepath.Abs(filepath.Join(path, api))

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.WriteFile(path, []byte(template), os.ModePerm)

	if err != nil {
		return "Failed To Add API", false
	}

	return message, success
}

func DeleteAPI(path string) {
	path, err := filepath.Abs(path)

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.Remove(path)

	if err != nil {
		panic("Invalid API Path")
	}
}

func AddCollection(name string, path string) (message string, success bool) {
	message, success = validate("COLLECTION", name, path)

	if !success {
		return message, success
	}

	path, err := filepath.Abs(filepath.Join(path, name))

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.Mkdir(path, os.ModePerm)

	if err != nil {
		return "Failed To Add Collection", false
	}

	return message, success
}

func DeleteCollection(path string) {
	path, err := filepath.Abs(path)

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.RemoveAll(path)

	if err != nil {
		panic("Invalid Collection Path")
	}
}

func validate(kind string, name string, path string) (message string, success bool) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic(err.Error())
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
