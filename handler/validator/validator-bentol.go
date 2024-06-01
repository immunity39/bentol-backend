package validator

import (
	"app/domain/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func SetupValidator() error {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("task_status", ValidateTaskStatus); err != nil {
			return err
		}
	}
	return nil
}
func ValidateTaskStatus(fl validator.FieldLevel) bool {
	return model.TaskStatusMap[model.TaskStatus(fl.Field().String())]
}
