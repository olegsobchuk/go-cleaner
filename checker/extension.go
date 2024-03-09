package checker

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

func extension(fileName string) string {
	ext := filepath.Ext(fileName)
	ext = strings.TrimPrefix(ext, ".")
	return strings.ToLower(ext)
}

func IsExtMatch(file fs.DirEntry, exts []string) bool {
	ext := extension(file.Name())
	return slices.Contains(exts, ext)
}
