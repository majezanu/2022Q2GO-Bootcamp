package datastore

import "io"

type OpenerCloser interface {
	OpenToRead() (io.ReadWriter, error)
	OpenToWrite() (io.ReadWriter, error)
	Close() error
}
