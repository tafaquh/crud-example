package login

import (
	"github.com/tafaquh/crud-example/database/schema"
	"github.com/tafaquh/crud-example/utils/errors"
)

func (s *LoginRequest) GetUserByEmail() (*schema.Users, *errors.RestErr) {
	var user schema.Users
	if err := schema.Database.First(&user, "email = ?", s.Email); err.Error != nil {
		return nil, errors.NewInternalServerError("error when trying to get user by confirmation token: " + err.Error.Error())
	}
	if user.StatusID < 2 {
		return nil, errors.NewBadRequestError("You must confirm your email first")
	}
	return &user, nil
}
