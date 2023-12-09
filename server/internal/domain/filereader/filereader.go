package filereader

import "io"

type FileReader interface {
	ReadFile() (io.Reader, error)
}