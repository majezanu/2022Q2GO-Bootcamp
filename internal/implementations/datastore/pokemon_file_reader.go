package datastore

import (
	"io"
	"majezanu/capstone/internal/contracts/datastore"
	"os"
)

type pokemonFileReader struct {
	CsvPath string
	File    *os.File
}

func (p pokemonFileReader) Open() (io.ReadWriter, error) {
	var err error
	p.File, err = os.Open(p.CsvPath)
	return p.File, err
}

func (p pokemonFileReader) Close() error {
	if p.File == nil {
		return nil
	}
	return p.File.Close()
}

func NewPokemonFileReader(csvPath string) datastore.OpenerCloser {
	return &pokemonFileReader{csvPath, nil}
}
