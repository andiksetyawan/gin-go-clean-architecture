package storage

import "io"

type Storage interface {
	Upload(fileID string, file io.Reader, bucket string) error
	//Download()
}
