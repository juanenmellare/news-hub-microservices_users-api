package api

import (
	"github.com/gin-gonic/gin"
	"news-hub-microservices_users-api/errors"
	"news-hub-microservices_users-api/utils"
)

func validateStringField(fieldName string, field *string, notValidFields *[]string) {
	if field == nil || *field == "" {
		*notValidFields = append(*notValidFields, fieldName)
	}
}

func validateNotValidFieldsSlice(notValidFields []string) {
	if len(notValidFields) > 0 {
		panic(errors.NewRequestFieldsShouldNotBeEmptyError(notValidFields))
	}
}

func marshallRequestBody(context *gin.Context, i interface{}) {
	if err := context.BindJSON(&i); err != nil {
		panic(errors.NewBadRequestApiError(err.Error()))
	}
}

type Request interface {
	MarshallAndValidate(context *gin.Context)
}

type CreateUserRequest struct {
	FirstName      *string `json:"firstName"`
	LastName       *string `json:"lastName"`
	Email          *string `json:"email"`
	Password       *string `json:"password"`
	PasswordRepeat *string `json:"passwordRepeat"`
}

func (r *CreateUserRequest) MarshallAndValidate(context *gin.Context) {
	marshallRequestBody(context, r)
	notValidFields := utils.NewStringSlice()
	validateStringField("firstName", r.FirstName, &notValidFields)
	validateStringField("lastName", r.LastName, &notValidFields)
	validateStringField("email", r.Email, &notValidFields)
	validateStringField("password", r.Password, &notValidFields)
	validateStringField("passwordRepeat", r.PasswordRepeat, &notValidFields)
	validateNotValidFieldsSlice(notValidFields)
	if *r.Password != *r.PasswordRepeat {
		panic(errors.NewBadRequestApiError("the fields 'password' and 'passwordRepeat' doesn't match"))
	}
}
