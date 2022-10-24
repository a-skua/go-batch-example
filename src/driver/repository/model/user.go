package model

import (
	"bytes"
	"encoding/json"
	"github.com/a-skua/go-batch-example/pkg/user"
	"io"
)

type User struct {
	reader io.Reader
}

func ParseUser(reader io.Reader) *User {
	return &User{reader}
}

func NewUser(u *user.User) (*User, error) {
	type value struct {
		ID       user.ID       `json:"id"`
		Name     user.Name     `json:"name"`
		Password user.Password `json:"password"`
	}

	body := struct {
		User value `json:"user"`
	}{
		User: value{
			ID:       u.ID,
			Name:     u.Name,
			Password: u.Password,
		},
	}

	bin, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	return &User{
		reader: bytes.NewBuffer(bin),
	}, nil
}

func (u *User) Read(p []byte) (int, error) {
	return u.reader.Read(p)
}

func (u *User) Entity() (*user.User, error) {
	body := struct {
		User struct {
			ID       user.ID       `json:"id"`
			Name     user.Name     `json:"name"`
			Password user.Password `json:"password"`
		} `json:"user"`
	}{}
	err := json.NewDecoder(u.reader).Decode(&body)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID:       body.User.ID,
		Name:     body.User.Name,
		Password: body.User.Password,
	}, nil
}
