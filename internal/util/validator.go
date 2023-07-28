package util

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(request interface{}) []*ValidateResponse {

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("fa")

	validate = validator.New()

	fa_translations.RegisterDefaultTranslations(validate, trans)

	var errors []*ValidateResponse

	err := validate.Struct(request)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			var element ValidateResponse
			element.Field = err.Field()
			element.Message = err.Translate(trans)
			errors = append(errors, &element)
		}
	}

	return errors
}
