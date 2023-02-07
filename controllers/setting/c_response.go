package setting

import (
	settingEntity "go-template/entities/setting"
	"go-template/helpers/interfacepkg"
)

type SettingResponse struct {
	ID        int         `json:"id" example:"1"`
	Name      string      `json:"name" example:"setting_name"`
	Value     interface{} `json:"value"`
	CreatedAt string      `json:"created_at" example:"2022-06-22 11:34:19.214 +0700"`
	UpdatedAt string      `json:"updated_at" example:"2022-06-22 11:34:19.214 +0700"`
}

func FromService(res settingEntity.Setting) SettingResponse {
	return SettingResponse{
		ID:        res.ID,
		Name:      res.Name,
		Value:     interfacepkg.Unmarshall(res.Value),
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}
}

func FromServices(entities []settingEntity.Setting) []SettingResponse {
	res := []SettingResponse{}
	for i := range entities {
		res = append(res, FromService(entities[i]))
	}
	return res
}
