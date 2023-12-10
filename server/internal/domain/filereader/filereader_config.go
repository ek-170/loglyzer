package filereader

type FileReaderConfig struct {
	// common conf
	FileReadMode string
	Path string
	// ssh conf
	SshKeyPath   string
	UserName     string
	Password     string
	Host         string
	Port         *int
}
