package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateStruct valida uma struct usando as tags de validação
func ValidateStruct(s interface{}) map[string]string {
	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	errors := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errors[field] = getErrorMessage(err)
	}

	return errors
}

// getErrorMessage retorna uma mensagem de erro amigável
func getErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("O campo %s é obrigatório", strings.ToLower(err.Field()))
	case "min":
		return fmt.Sprintf("O campo %s deve ter no mínimo %s caracteres", strings.ToLower(err.Field()), err.Param())
	case "max":
		return fmt.Sprintf("O campo %s deve ter no máximo %s caracteres", strings.ToLower(err.Field()), err.Param())
	default:
		return fmt.Sprintf("O campo %s é inválido", strings.ToLower(err.Field()))
	}
}
