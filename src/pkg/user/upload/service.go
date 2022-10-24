package upload

import (
	"errors"
	"fmt"
	"github.com/a-skua/go-batch-example/pkg/user"
)

var ErrEOL = errors.New("End Of Line")

type Data interface {
	Next() (*user.User, error)
}

type Repository interface {
	CreateUser(*user.User) (*user.User, error)
}

func Upload(data Data, repo Repository) error {
	for {
		user, err := data.Next()
		if err == ErrEOL {
			return nil
		}

		if err != nil {
			return err
		}

		user, err = repo.CreateUser(user)
		if err != nil {
			return err
		}
		fmt.Println(user)
	}
}
