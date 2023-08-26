package storage

import (
	"os"
	"path/filepath"
)

var storagePath string
var storageName = "uploads"

func SetupStorage() {
	storage, _ := filepath.Abs(storageName)

	storagePath = storage
}

func SetupTestStorage() {
	storage, _ := filepath.Abs(storageName + "_test")

	storagePath = storage
}

func GetStoragePath() string {
	return storagePath
}

func ClearTestStorage() {
	storage, _ := filepath.Abs(storageName + "_test")
	os.RemoveAll(storage)
}
