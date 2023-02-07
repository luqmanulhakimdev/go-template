package setting

import (
	"fmt"
	"go-template/controllers"
	settingEntity "go-template/entities/setting"
	"go-template/helpers/interfacepkg"

	"github.com/go-playground/validator"
)

type SettingParameter struct {
	Name string `query:"name"`
	controllers.DefaultParameter
}

type SettingRequest struct {
	Name  string      `json:"name" validate:"required"`
	Value interface{} `json:"value" validate:"required"`
}

func (i *SettingRequest) Validate() error {
	err := validator.New().Struct(i)
	if err != nil {
		for _, er := range err.(validator.ValidationErrors) {
			return fmt.Errorf("%v %v", er.Field(), er.ActualTag())
		}
	}
	return nil
}

func (r *SettingRequest) ToService() (res *settingEntity.Setting) {
	res = &settingEntity.Setting{
		Name:  r.Name,
		Value: interfacepkg.Marshal(r.Value),
	}
	return
}
