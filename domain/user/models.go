package user

import (
	"github.com/MartIden/go-hexagonal/utils/schemas"
)

type UserModel struct {
	Email             string
	EncryptedPassword string
}

type UserModelInDB struct {
	schemas.IDSchemaMixin
	UserModel
}

type UserModelInCreate struct {
	UserModel
}

type UserModelInResponse struct {
	UserModelInDB
}