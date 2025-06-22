package models

type CreateDirectoryRequest struct {
	DirectoryName string `json:"directory_name" binding:"required"`
}

type CreateFileRequest struct {
	DirectoryName string `json:"directory_name" binding:"required"`
	FileName string `json:"file_name" binding:"required"`
	Content string `json:"content"`
}

type FileUploadResponse struct {
	Message string `json:"message"`
	Filename string `json:"filename"`
	Path string `json:"path"`
}