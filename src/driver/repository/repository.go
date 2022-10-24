package repository

import (
	"errors"
	"github.com/a-skua/go-batch-example/driver/repository/model"
	"github.com/a-skua/go-batch-example/pkg/user"
	"github.com/a-skua/go-batch-example/pkg/user/upload"
	"net/http"
	"net/url"
)

type Repository interface {
	upload.Repository
}

type repository struct {
	baseURL url.URL
}

func New(url *url.URL) Repository {
	return &repository{*url}
}

func (r *repository) CreateUser(u *user.User) (*user.User, error) {
	user, err := model.NewUser(u)
	if err != nil {
		return nil, err
	}

	url := r.baseURL // copy
	url.Path += "/user"
	resp, err := http.Post(url.String(), "application/json", user)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error response")
	}
	user = model.ParseUser(resp.Body)
	return user.Entity()
}
