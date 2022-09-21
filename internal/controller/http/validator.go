package http

import (
	"net/http"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ess "github.com/go-playground/validator/v10/translations/es"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator  *validator.Validate
	translator ut.Translator
}

func NewValidator(validator *validator.Validate) *CustomValidator {
	en := es.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("es")
	ess.RegisterDefaultTranslations(validator, trans)

	return &CustomValidator{
		validator:  validator,
		translator: trans,
	}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		errs := err.(validator.ValidationErrors)

		return echo.NewHTTPError(http.StatusBadRequest, errs.Translate(cv.translator))
	}

	return nil
}
