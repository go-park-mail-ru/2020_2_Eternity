package utils

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
)

func ValidProfile(profile api.SignUp) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Email, validation.Required, is.EmailFormat),
		validation.Field(&profile.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric),
		validation.Field(&profile.Password, validation.Required, validation.Length(8, 50), is.Alphanumeric),
		validation.Field(&profile.Name, is.Alphanumeric),
		validation.Field(&profile.Surname, is.Alphanumeric),
	)
}

func ValidUpdate(profile api.UpdateUser) error {
	return validation.ValidateStruct(&profile,
		validation.Field(&profile.Email, is.EmailFormat),
		validation.Field(&profile.Username, validation.Length(5, 50), is.Alphanumeric),
		validation.Field(&profile.Name, is.Alphanumeric),
		validation.Field(&profile.Surname, is.Alphanumeric),
	)
}

func ValidUsername(user api.UserAct) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username, validation.Required, validation.Length(5, 50), is.Alphanumeric),
	)
}

func ValidPasswords(pswds api.UpdatePassword) error {
	return validation.ValidateStruct(&pswds,
		validation.Field(&pswds.OldPassword, validation.Required, validation.Length(8, 50), is.Alphanumeric),
		validation.Field(&pswds.NewPassword, validation.Required, validation.Length(8, 50), is.Alphanumeric),
	)
}

