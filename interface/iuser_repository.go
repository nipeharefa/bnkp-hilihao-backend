package interfaces

import m "bnkp-hilihao.org/model"

type UserRepository interface {
	CreateUser(user *m.User)
	GetAllUser() ([]m.User, error)
	FindByID(ID string) (m.User, error)
	UpdateUser() error
}
