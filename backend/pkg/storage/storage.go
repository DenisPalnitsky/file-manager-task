package storage

import (
	"log/slog"
	"os"

	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/dtos"
)

type FileSystemStorage struct {
	dir string
}

func NewFileSystemStorage() *FileSystemStorage {
	err := os.RemoveAll("./tmp")
	if err != nil {
		slog.Error("Error removing folder", err)
	}
	err = os.Mkdir("./tmp", 0755)
	if err != nil {
		slog.Error("Error creating folder", err)
	}
	return &FileSystemStorage{
		dir: "./tmp",
	}
}

func (s *FileSystemStorage) AddFile(path string, m *dtos.Metadata) error {
	// copy file to dir
	// Read all content of src to data, may cause OOM for a large file.
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	// Write data to dst
	err = os.WriteFile(s.dir+"/"+m.Filename, data, 0644)
	return err
}

func (s *FileSystemStorage) ListFiles() ([]*dtos.Metadata, error) {
	files, err := os.ReadDir(s.dir)
	res := []*dtos.Metadata{}
	if err != nil {
		return res, err
	}

	for _, f := range files {
		res = append(res, &dtos.Metadata{
			Filename: f.Name(),
		})
	}
	return res, nil
}
