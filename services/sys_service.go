package services

import (
	// "sync"
	"fmt"
	"golangapi/models"
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