package services

import (
	// "sync"
	"fmt"
	"golangapi/models"
	"io"
	"os"
	"path/filepath"

	"gorm.io/gorm"
)

type SysServices struct {
	DB 		*gorm.DB
	// downloadMutex sync.Mutex
}

func NewSysService(db *gorm.DB) *SysServices {
	return &SysServices{
		DB: db,
	}
}

func (ss *SysServices) CreateDirectory(req *models.CreateDirectoryRequest) error {
	err := os.Mkdir(req.DirectoryName, 0755) // 0755 merupakan permission dari si directory
	if err != nil {
		return fmt.Errorf("gagal membuat directory: %v", err)
	}
	return nil
}

func (ss *SysServices) CreateFile(req *models.CreateFileRequest) (string, error) {
	// misalkan folder tidak ada, dia bikin folder otomatis
	// misalkan folder ada, dia gabakal bikin lagi
	if err := os.MkdirAll(req.DirectoryName, 0755); err != nil {
		return "", fmt.Errorf("gagal membuat directory: %v", err)
	}

	// filepath 
	// testing/file.txt
	filepath := filepath.Join(req.DirectoryName, req.FileName)
	
	// fungsinya buat bikin file. contoh file.txt
	file, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("gagal membuat file: %v", err)
	}

	// dijalankan ketika semuanya sukses
	// file close ini menutup si file...
	defer file.Close()

	// write string untuk menulis sebuah file
	_, err = file.WriteString(req.Content)
	if err != nil {
		return "", fmt.Errorf("gagal menulis file: %v", err)
	}
	
	return filepath, nil
}

func (ss *SysServices) UploadFile(filename string, src io.Reader) (any, error) {
	const uploadDir = "uploads"

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return nil, fmt.Errorf("gagal membuat folder : %v", err)
	}

	cleanFilename := filepath.Base(filename)
	pathFolder := filepath.Join(uploadDir, cleanFilename)

	dst, err := os.Create(pathFolder)

	if err != nil {
		return nil, fmt.Errorf("gagal membuat file : %v", err)
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan file : %v", err)
	}

	response := map[string]any{
		"message":  "File berhasil diupload!",
		"filename": cleanFilename,
		"path":     pathFolder,
	}

	return response, nil
}