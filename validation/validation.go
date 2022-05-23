package validation

import (
	"github.com/RakibSiddiquee/go-fiber-jwt-auth/models"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// type ErrorResponse struct {
// 	FailedField string
// 	Value       string
// }

var (
	uni *ut.UniversalTranslator
)

func ValidateStruct(user models.User) map[string]string {
	var errors = make(map[string]string)

	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	var validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		//errors = errs.Translate(trans)
		// result
		// {
		// 	"User.Email": "Email must be a valid email address",
		// 	"User.Name": "Name must be at least 3 characters in length"
		// }

		//fmt.Println(errs.Translate(trans))
		for _, err := range errs {
			errors[err.Field()] = err.Translate(trans)
			//fmt.Println("sdfdf", err.Namespace())
			//var element ErrorResponse
			//element.FailedField = err.Field()
			// element.Tag = err.Tag()
			//element.Value = err.Translate(trans)
			//errors = append(errors, &element)
			//errors = append(map[string]string{err.Field(), err.Translate(trans)})
		}
		// fmt.Println(err.Namespace())
		// fmt.Println(err.Field())
		// fmt.Println(err.StructNamespace())
		// fmt.Println(err.StructField())
		// fmt.Println(err.Tag())
		// fmt.Println(err.ActualTag())
		// fmt.Println(err.Kind())
		// fmt.Println(err.Type())
		// fmt.Println(err.Value())
		// fmt.Println(err.Param())
		// fmt.Println()
		// }
	}

	// errors result
	// {
	// 	"Email": "Email must be a valid email address",
	// 	"Name": "Name must be at least 3 characters in length"
	// }

	return errors
}
