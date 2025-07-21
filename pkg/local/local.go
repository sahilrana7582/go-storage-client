package local

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

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

func copyFile(src, dst string) error {

	fileIn, err := os.Open(src)
	if err != nil {
		return err
	}

	defer func(fileIn *os.File) {
		err := fileIn.Close()
		if err != nil {
			fmt.Println("Error creating file:", err)
		}
	}(fileIn)

	fileOut, err := os.Create(dst)
	if err != nil {
		return err
	}

	_, err = io.Copy(fileOut, fileIn)
	if err != nil {
		err := fileOut.Close()
		if err != nil {
			return err
		}
		return err
	}

	return fileOut.Close()
}
