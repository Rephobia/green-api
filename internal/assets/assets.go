package assets

import (
	"embed"
	"fmt"
	"io/fs"
)

//go:embed frontend/*
var frontendFiles embed.FS

func GetFrontendFiles() (*fs.FS, error) {
	content, err := fs.Sub(frontendFiles, "frontend")
	if err != nil {
		return nil, fmt.Errorf("can't load frontend files: %w", err)
	}

	return &content, nil
}
