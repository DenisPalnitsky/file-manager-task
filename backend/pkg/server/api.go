package server

import "github.com/DenisPalnitsky/file-manager-task/backend/pkg/dtos"

type FileListGetRequest struct {
}

type FileListGetResponse struct {
	Files []*dtos.Metadata `json:"files"`
}

type FilePostRequest struct {
	File []byte `json:"file"`
}

type FilePostResponse struct {
	Message string `json:"message"`
}
