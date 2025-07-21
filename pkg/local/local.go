package local

import "github.com/sahilrana7582/go-storage/pkg/storage"

type LocalStorage struct {
	*storage.Client
}

type Config struct {
	LocalPath  string
	RemotePath string
}

func New(config Config) storage.Storage {
	return LocalStorage{
		Client: &storage.Client{
			LocalPath:  config.LocalPath,
			RemotePath: config.RemotePath,
		},
	}
}

func (l LocalStorage) Upload(fileName string) error {

}
