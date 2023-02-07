package sqlstore

import (
	u "github.com/MartIden/go-hexagonal/domain/user"
)

type UserRepo struct {
	store *Store
}

func (ur UserRepo) Create(model *u.UserModelInDB) error {
	return nil
}

func (ur UserRepo) Read(uid uint) (*u.UserModelInDB, error) {
	return &u.UserModelInDB{}, nil
}

func (ur UserRepo) ReadAll() ([]*u.UserModelInDB, error) {
	var users []*u.UserModelInDB = []*u.UserModelInDB{
		&u.UserModelInDB{},
	}
	return users, nil
}

func (ur UserRepo) Update() (*u.UserModelInDB, error) {
	return &u.UserModelInDB{}, nil
}

func (ur UserRepo) Delete(uid uint) (*u.UserModelInDB, error) {
	return &u.UserModelInDB{}, nil
}