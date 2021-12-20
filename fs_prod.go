// +build !debug

package main

import (
	"embed"
	"io/fs"
)

//go:embed templates/*
var embedFs embed.FS

func GetTemplates() fs.FS {
	return embedFs
}
