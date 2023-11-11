package storage

import (
	"testing"

	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/dtos"
	"github.com/stretchr/testify/assert"
)

func TestFileSystemStorage_AddFile(t *testing.T) {

	s := NewFileSystemStorage()
	err := s.AddFile("/home/denis/go/src/github.com/DenisPalnitsky/file-manager-task/backend/pkg/storage/testdata/testfile.txt", &dtos.Metadata{Filename: "testfile.txt"})
	assert.NoError(t, err)

	err = s.AddFile("/home/denis/go/src/github.com/DenisPalnitsky/file-manager-task/backend/pkg/storage/testdata/testfile.txt", &dtos.Metadata{Filename: "testfile1.txt"})
	assert.NoError(t, err)

	files, err := s.ListFiles()
	if assert.NoError(t, err) {
		assert.Equal(t, 2, len(files))
		assert.Equal(t, "testfile.txt", files[0].Filename)
		assert.Equal(t, "testfile1.txt", files[1].Filename)
	}
}
