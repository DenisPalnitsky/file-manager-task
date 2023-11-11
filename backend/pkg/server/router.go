package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/DenisPalnitsky/file-manager-task/backend/cmd/version"
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg"
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/dtos"
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/rest"
	"github.com/DenisPalnitsky/file-manager-task/backend/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"

	"github.com/wI2L/fizz"
)

type FileStorage interface {
	AddFile(path string, m *dtos.Metadata) error
	ListFiles() ([]*dtos.Metadata, error)
}

type Router struct {
	fizz        *fizz.Fizz
	port        int
	fileStorage FileStorage
}

func (r *Router) Run() {
	slog.Info("Listening", "port", r.port)
	err := r.fizz.Engine().Run(fmt.Sprint(":", r.port))
	if err != nil {
		panic(err)
	}
}

func NewRouter(config *rest.HttpConfig, debugMode bool) *Router {
	f := rest.NewFizzRouter(config, pkg.ServiceName, version.Version, debugMode)

	r := &Router{
		fizz:        f,
		port:        config.Port,
		fileStorage: storage.NewFileSystemStorage(),
	}
	r.init()
	return r
}

func (r *Router) init() {
	filesGroup := r.fizz.Group("/files", "List of files", "Files operations")
	filesGroup.POST("", []fizz.OperationOption{
		fizz.Summary("Upload a file"),
		fizz.Description("Upload a file to the server. Use multipart/form-data encoding"),
	}, tonic.Handler(r.filePostHandler, http.StatusOK))

	filesGroup.GET("", nil, tonic.Handler(r.fileGetHandler, http.StatusOK))
}

func (r *Router) fileGetHandler(c *gin.Context, req *FileListGetRequest) (*FileListGetResponse, error) {
	slog.Info("Listing files")

	m, err := r.fileStorage.ListFiles()
	if err != nil {
		return nil, err
	}

	return &FileListGetResponse{Files: m}, nil

}

func (r *Router) filePostHandler(c *gin.Context) (*FilePostResponse, error) {
	file, err := c.FormFile("file")
	if err != nil {
		return nil, rest.HttpError{
			HttpCode: http.StatusBadRequest,
			Message:  "Error getting file",
		}
	}

	log := slog.With("req", "uploading file request").With("filename", file.Filename)

	if err != nil {
		log.Error("Error getting file info", err)
		return nil, rest.HttpError{
			HttpCode: http.StatusBadRequest,
			Message:  "Error getting file info",
		}
	}

	tmpDir := os.TempDir()
	t, err := os.CreateTemp(tmpDir, "upload-*.txt")
	if err != nil {
		log.Error("Error creating temp file", "error", err)
		return nil, rest.HttpError{
			HttpCode: http.StatusInternalServerError,
			Message:  "Error creating temp file",
		}
	}

	// Upload the file to specific dst.
	err = c.SaveUploadedFile(file, t.Name())
	if err != nil {
		log.Error("Error saving file", "error", err)
		return nil, rest.HttpError{
			HttpCode: http.StatusInternalServerError,
			Message:  "Saving file error",
		}
	}

	err = r.fileStorage.AddFile(t.Name(), &dtos.Metadata{Filename: file.Filename})
	if err != nil {
		log.Error("Error adding file to storage", "error", err)
		return nil, rest.HttpError{
			HttpCode: http.StatusInternalServerError,
			Message:  "Error adding file to storage",
		}
	}

	return &FilePostResponse{Message: "File uploaded successfully"}, nil
}
