package filereader

import fr "github.com/ek-170/loglyzer/internal/domain/filereader"

func InitFileReader(conf fr.FileReaderConfig) fr.FileReader{
	switch conf.FileReadMode {
	case "sftp":
		return NewSftpFileReader(conf)
	case "local":
		return NewLocalFileReader(conf)
	default:
		return NewLocalFileReader(conf)
	}
}