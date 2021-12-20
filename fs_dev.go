// +build debug

package main

import (
	"io/fs"
	"os"
)

func GetTemplates() fs.FS {
	wd, _ := os.Getwd()

	return os.DirFS(wd)
}
