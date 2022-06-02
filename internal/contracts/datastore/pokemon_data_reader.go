package datastore

import "io"

type OpenerCloser interface {
	Open() (io.ReadWriter, error)
	Close() error
}
