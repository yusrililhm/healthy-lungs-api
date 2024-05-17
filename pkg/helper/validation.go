package helper

import (
	"expert_systems_api/pkg/exception"

	"github.com/asaskevich/govalidator"
)

func ValidationStruct(s any) exception.Exception {

	if _, err := govalidator.ValidateStruct(s); err != nil {
		return exception.NewBadRequestError(err.Error())
	}

	return nil
}
