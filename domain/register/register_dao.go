package register

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/tafaquh/crud-example/domain/email"
	"github.com/tafaquh/crud-example/utils"
	"github.com/tafaquh/crud-example/utils/errors"
)

func (s *RegisterUserRequest) CreateUser() (*RegisterAdminResponse, *errors.RestErr) {
	s.ConfirmationToken = generateConfirmationToken()
	// send email verification
	// if err := s.SendEmail(); err != nil {
	// 	return nil, err
	// }
	// TODO: save to database
	if err := s.CreateUserToDB(); err != nil {
		return nil, err
	}
	return &RegisterAdminResponse{Email: s.Email}, nil
}

func Confirmation(token string) *errors.RestErr {
	user, err := GetUserByConfirmationToken(token)
	if err != nil {
		return err
	}
	user.ConfirmationToken = ""
	user.UpdatedAt = time.Now()
	if err := UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *RegisterUserRequest) SendEmail() *errors.RestErr {
	fmt.Printf("payload: %+v", s)
	templateData := email.TemplateEmailRequest{
		TemplateName: "confirmEmail.html",
		Recipients:   []string{s.Email},
		Subject:      "Register Confirmation",
		Data: email.RegisterConfirmationRequest{
			FullName: s.FirstName + " " + s.LastName,
			Role:     "admin",
			Url:      "https://www.garapin.com/api/v1/user/register/confirmation/" + s.ConfirmationToken,
		},
	}
	htmlTemplate, err := utils.CreateTemplate(&templateData)
	if err == nil {
		err := utils.SendEmail(templateData.Recipients, templateData.Subject, htmlTemplate)
		if err != nil {
			fmt.Printf("\nerror: %+v\n\n", err)
			return &errors.RestErr{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			}
		}
	}
	return nil
}

func generateConfirmationToken() string {
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}
	return token + time.Now().Format("20060102150405")
}
