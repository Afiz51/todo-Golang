package validate

import (
	"fmt"

	"github.com/Afiz51/TodoGoRest/handlers"
	"github.com/go-playground/validator"
)

func UserStructLevelValidation(sl validator.StructLevel) {
	fmt.Println("UserStructLevelValidation called")
	user := sl.Current().Interface().(handlers.CreateUserRequest)

	if len(user.FirstName) == 0 && len(user.LastName) == 0 {
		fmt.Println("UserStructLevelValidation called")
		sl.ReportError(user.FirstName, "FirstName", "fname", "fnameorlname", "")
		sl.ReportError(user.LastName, "LastName", "lname", "fnameorlname", "")
	}
}
