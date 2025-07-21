package local

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sahilrana7582/go-storage/pkg/storage"
)

type LocalStorage struct {
	*storage.Client
}

type Config struct {
	LocalPath  string
	RemotePath string
}

func NewLocalConfig(localPath, remotePath string) Config {
	return Config{
		LocalPath:  localPath,
		RemotePath: remotePath,
	}
}

func NewLocal(config Config) storage.Storage {
	return LocalStorage{
		Client: &storage.Client{
			LocalPath:  config.LocalPath,
			RemotePath: config.RemotePath,
		},
	}
}

func (l LocalStorage) Upload(fileName string) error {

	_, err := os.Stat(filepath.Join(l.LocalPath, fileName))
	if os.IsNotExist(err) {
		return err
	}

	err = copyFile(filepath.Join(l.LocalPath, fileName), filepath.Join(l.RemotePath, fileName))
	if err != nil {
		return err
	}

	return nil
}

func (l LocalStorage) UploadAllFiles(dirPath string) error {

	levelDeep := 0
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error accessing path:", err)
			return err
		}

		if !info.IsDir() {
			stringSlice := strings.Split(path, "/")

			dst := strings.Join(stringSlice[len(stringSlice)-levelDeep:len(stringSlice)], "/")

			err = copyFile(path, filepath.Join(l.RemotePath, dst))

			if err != nil {
				return err
			}
		} else {
			levelDeep++
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func copyFile(src, dst string) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	fileIn, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fileIn.Close()

	fileOut, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer fileOut.Close()

	_, err = io.Copy(fileOut, fileIn)
	if err != nil {
		return err
	}

	return nil
}
