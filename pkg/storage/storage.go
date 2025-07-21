package storage

type Storage interface {
	Upload(fileName string) error
	UploadAllFiles(dirPath string) error
	// Download(remotePath, localPath string) error
	Delete(remotePath, fileName string) error
	Prune(retentionDays int) error
}

type Client struct {
	LocalPath  string
	RemotePath string
}
