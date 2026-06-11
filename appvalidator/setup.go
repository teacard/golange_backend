package appvalidator

import (
	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

// Init 註冊自訂 validator tag
func Init() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return
	}
	// positiveInt：整數必須 >= 1
	_ = v.RegisterValidation("positiveInt", func(fl validator.FieldLevel) bool {
		return fl.Field().Int() >= 1
	})
}
