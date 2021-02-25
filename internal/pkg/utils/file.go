package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
)

func EnsureBaseDir(dst string) error {
	baseDir := path.Dir(dst)
	info, err := os.Stat(baseDir)
	if err == nil && info.IsDir() {
		return nil
	}
	return os.MkdirAll(baseDir, 0755)
}

func CreateFile(file *multipart.FileHeader, dst string) error {
	if err := EnsureBaseDir(dst); err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func DeleteFile(path string) error {
	var err = os.Remove(path)
	return err
}

func ByteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cb", float64(b)/float64(div), "KMGTPE"[exp])
}
