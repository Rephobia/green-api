package validatorx

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(s any) []string {
	err := validator.New().Struct(s)
	if err == nil {
		return nil
	}

	var errs []string

	for _, e := range func() validator.ValidationErrors {
		var target validator.ValidationErrors

		_ = errors.As(err, &target)

		return target
	}() {
		errs = append(errs, mapError(e))
	}

	return errs
}

func mapError(e validator.FieldError) string {
	switch e.Field() {
	case "IdInstance":
		return "idInstance обязателен"
	case "APITokenInstance":
		return "apiTokenInstance обязателен"
	case "ChatID":
		return "chatId обязателен"
	case "Message":
		return "message обязателен"
	default:
		return "некорректное поле"
	}
}
