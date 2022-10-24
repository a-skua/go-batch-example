package file

import (
	"bytes"
	"encoding/csv"
	"errors"
	"github.com/a-skua/go-batch-example/pkg/user"
	"github.com/a-skua/go-batch-example/pkg/user/upload"
	"io"
	"os"
)

type UserData interface {
	upload.Data
}

type userData struct {
	reader *csv.Reader
}

func (data *userData) first() bool {
	return data.reader.InputOffset() == 0
}

func NewUserData(filename string) (UserData, error) {
	bin, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &userData{
		reader: csv.NewReader(bytes.NewBuffer(bin)),
	}, nil
}

func (data *userData) Next() (*user.User, error) {
	if data.first() {
		// skip header
		_, err := data.reader.Read()
		if err != nil {
			return nil, err
		}
	}

	row, err := data.reader.Read()
	if err == io.EOF {
		return nil, upload.ErrEOL
	}
	if err != nil {
		return nil, err
	}

	const columns = 2
	if len(row) != columns {
		return nil, errors.New("unexpected columns")
	}

	const (
		name     = 0
		password = 1
	)

	return user.New(
		user.Name(row[name]),
		user.Password(row[password]),
	), nil
}
