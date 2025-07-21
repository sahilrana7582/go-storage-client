package storage

type Storage interface {
	Upload(localPath, remotePath string) error
	Download(remotePath, localPath string) error
	Delete(remotePath string) error
	List(remotePath string) ([]string, error)
	ViewFile(fileName, remotePath string) (string, error)
}

type Client struct {
	LocalPath  string
	RemotePath string
}
