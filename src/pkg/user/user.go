package user

type ID int

type Name string

type Password string

type User struct {
	ID       ID
	Name     Name
	Password Password
}

func New(name Name, pw Password) *User {
	return &User{
		Name:     name,
		Password: pw,
	}
}
