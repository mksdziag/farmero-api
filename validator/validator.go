package validator

import (
	"errors"
	"log"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type AppValidator struct {
	validator *validator.Validate
	trans     ut.Translator
}

func (av *AppValidator) Validate(i interface{}) error {
	if err := av.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)

		var errorMessage string
		for _, e := range errs {
			errorMessage += e.Translate(av.trans)
		}

		return errors.New(errorMessage)
	}
	return nil
}

func NewAppValidator() *AppValidator {
	en := en.New()
	uni := ut.New(en, en)

	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("translator not found")
	}

	v := validator.New()

	registerTranslations(v, trans)

	en_translations.RegisterDefaultTranslations(v, trans)

	return &AppValidator{validator: v, trans: trans}
}

type ValidatorTranslation struct {
	Name           string
	MessagePattern string
	IncludeParam   bool
}

var validatorTranslations = []ValidatorTranslation{
	{Name: "required", MessagePattern: "{0} is required", IncludeParam: false},
	{Name: "min", MessagePattern: "{0} have to be at least {1} characters long.", IncludeParam: true},
	{Name: "max", MessagePattern: "{0} have to be at most {1} characters long.", IncludeParam: true},
}

func registerTranslations(v *validator.Validate, trans ut.Translator) {
	for _, item := range validatorTranslations {
		currentItem := item
		v.RegisterTranslation(
			currentItem.Name,
			trans,
			func(ut ut.Translator) error {
				return ut.Add(currentItem.Name, currentItem.MessagePattern, true)
			},
			func(ut ut.Translator, fe validator.FieldError) string {
				var param string
				if currentItem.IncludeParam {
					param = fe.Param()
				}

				t, _ := ut.T(currentItem.Name, fe.Field(), param)
				return t
			})

	}
}
