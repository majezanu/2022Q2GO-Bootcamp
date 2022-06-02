package datastore

import "io"

type ReadWriteCloser interface {
	Read() (io.Reader, error)
	Write() error
	Close() error
}
