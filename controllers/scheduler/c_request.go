package scheduler

import (
	"fmt"

	"github.com/go-playground/validator"
)

type SchedulerRequest struct {
	Type string `json:"type" validate:"required"`
}

func (i *SchedulerRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}

	return nil
}
