package core

import (
	"os"
	"path/filepath"
)

type Collection struct {
	name    string
	path    string
	entries []os.DirEntry
}

func NewCollection() *Collection {
	return &Collection{
		name:    "",
		path:    "",
		entries: make([]os.DirEntry, 0),
	}
}

func (collection *Collection) Name() (name string) {
	return collection.name
}

func (collection *Collection) Path() (path string) {
	return collection.path
}

func (collection *Collection) Entries() (entries []os.DirEntry) {
	return collection.entries
}

func (collection *Collection) Read(path string) *Collection {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("Invalid Collection Path")
	}

	name := filepath.Base(path)

	return &Collection{
		name:    name,
		path:    path,
		entries: entries,
	}
}

func (collection *Collection) Create(name string, path string) (newCollection *Collection, message string, success bool) {
	message, success = validateCollection(name, path)

	if !success {
		return nil, message, success
	}

	path, err := filepath.Abs(filepath.Join(path, name))

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.Mkdir(path, os.ModePerm)

	if err != nil {
		return nil, "Failed To Create Collection", false
	}

	newCollection = collection.Read(path)

	return newCollection, message, success
}

func (collection *Collection) Delete(path string) {
	path, err := filepath.Abs(path)

	if err != nil {
		panic("Invalid Path Provided")
	}

	err = os.RemoveAll(path)

	if err != nil {
		panic("Invalid Collection Path")
	}
}

func validateCollection(name string, path string) (message string, success bool) {
	entries, err := os.ReadDir(path)

	if err != nil {
		panic("File Validation Failed")
	}

	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == name {
			return "Collection Already Exists", false
		}
	}

	return "Collection Added Successfully", true
}
