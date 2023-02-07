package setting

import (
	"context"
	"database/sql"
	"time"

	"go-template/client/redis"
	"go-template/controllers"
	settingRequest "go-template/controllers/setting"
	settingEntity "go-template/entities/setting"
	errHelper "go-template/helpers/errors"
	"go-template/helpers/interfacepkg"
	interfaceHelper "go-template/helpers/interfacepkg"

	"go-template/logger"
	"go-template/services"

	"github.com/jmoiron/sqlx"
)

type SettingRepo interface {
	CreateTx(ctx context.Context) (tx *sqlx.Tx, err error)
	SelectAll(ctx context.Context, parameters *settingRequest.SettingParameter) (data []settingEntity.Setting, err error)
	FindAll(ctx context.Context, parameters *settingRequest.SettingParameter) (data []settingEntity.Setting, count int, err error)
	FindOne(ctx context.Context, parameters *settingRequest.SettingParameter) (data settingEntity.Setting, err error)
	Create(ctx context.Context, tx *sqlx.Tx, data *settingEntity.Setting) (res int, err error)
	Update(ctx context.Context, tx *sqlx.Tx, data *settingEntity.Setting) error
	Delete(ctx context.Context, tx *sqlx.Tx, id int) error
}

type settingService struct {
	settingRepo SettingRepo
	redis       redis.IRedis
}

func NewSettingService(repo SettingRepo, redis redis.IRedis) *settingService {
	return &settingService{
		settingRepo: repo,
		redis:       redis,
	}
}

func (service settingService) SelectAll(ctx context.Context, parameters *settingRequest.SettingParameter) (res []settingEntity.Setting, err error) {
	parameters.Offset, parameters.Limit, parameters.Page, parameters.OrderBy, parameters.Sort =
		services.SetPaginationParameter(parameters.Page, parameters.Limit, settingEntity.MapOrderBy[parameters.OrderBy], parameters.Sort, settingEntity.OrderBy, settingEntity.OrderByString)

	res, err = service.settingRepo.SelectAll(ctx, parameters)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "select all setting query")
		return
	}

	return
}

func (service settingService) FindAll(ctx context.Context, parameters *settingRequest.SettingParameter) (res []settingEntity.Setting, pagination services.Pagination, err error) {
	parameters.Offset, parameters.Limit, parameters.Page, parameters.OrderBy, parameters.Sort =
		services.SetPaginationParameter(parameters.Page, parameters.Limit, settingEntity.MapOrderBy[parameters.OrderBy], parameters.Sort, settingEntity.OrderBy, settingEntity.OrderByString)

	res, count, err := service.settingRepo.FindAll(ctx, parameters)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "find all setting query")
		return
	}
	pagination = services.SetPaginationResponse(parameters.Page, parameters.Limit, count)

	return
}

func (service settingService) FindOne(ctx context.Context, parameters *settingRequest.SettingParameter) (res settingEntity.Setting, err error) {
	if parameters.Name != "" {
		service.redis.Get(ctx, redis.Setting+parameters.Name, &res)
		if res.ID != 0 {
			return
		}
	}

	res, err = service.settingRepo.FindOne(ctx, parameters)
	if err == sql.ErrNoRows {
		logger.ErrorWithStack(ctx, err, errHelper.DataNotFound.Message)
		return res, errHelper.DataNotFound.Error
	} else if err != nil {
		logger.ErrorWithStack(ctx, err, err.Error())
		return
	}

	err = service.redis.Set(ctx, redis.Setting+parameters.Name, &res, 60*time.Minute)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "set redis")
		return
	}

	return
}

func (service settingService) FindDefaultSetting(ctx context.Context) (res settingEntity.SettingDefault, err error) {
	data, err := service.settingRepo.FindOne(ctx, &settingRequest.SettingParameter{
		Name: settingEntity.SettingTypeDefault,
	})
	if err != nil {
		logger.ErrorWithStack(ctx, err, err.Error())
		return
	}

	// unmarshall value json string to struct
	interfacepkg.UnmarshallCb(data.Value, &res)

	return
}

func (service settingService) FindIncentiveBudgetSetting(ctx context.Context) (res settingEntity.SettingIncentiveBudget, err error) {
	data, err := service.settingRepo.FindOne(ctx, &settingRequest.SettingParameter{
		Name: settingEntity.SettingTypeIncentiveBudget,
	})
	if err != nil {
		logger.ErrorWithStack(ctx, err, err.Error())
		return
	}

	// unmarshall value json string to struct
	interfacepkg.UnmarshallCb(data.Value, &res)

	return
}

func (service settingService) FindIncentiveBudgetApprovalSetting(ctx context.Context) (res settingEntity.SettingIncentiveBudgetApproval, err error) {
	data, err := service.settingRepo.FindOne(ctx, &settingRequest.SettingParameter{
		Name: settingEntity.SettingTypeIncentiveBudgetApproval,
	})
	if err != nil {
		logger.ErrorWithStack(ctx, err, err.Error())
		return
	}

	// unmarshall value json string to struct
	interfacepkg.UnmarshallCb(data.Value, &res)

	return
}

func (service settingService) checkSettingDefault(ctx context.Context, data string) (err error) {
	values := settingEntity.SettingDefault{}
	err = interfaceHelper.ConvertStrToInterface(data, &values)
	if err != nil {
		err = errHelper.InvalidDataSetting.Error
		logger.ErrorWithStack(ctx, err, "error convert value")
		return
	}

	if values.AsanaEmailAssignee == "" {
		err = errHelper.AsanaEmailAssigneeRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if values.AsanaCommentMentionID == "" {
		err = errHelper.AsanaCommentMentionIDRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if len(values.AsanaEmailFollowers) == 0 {
		err = errHelper.AsanaEmailFollowersRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	return nil
}

func (service settingService) checkSettingIncentiveBudget(ctx context.Context, data string) (err error) {
	values := settingEntity.SettingIncentiveBudget{}
	err = interfaceHelper.ConvertStrToInterface(data, &values)
	if err != nil {
		err = errHelper.InvalidDataSetting.Error
		logger.ErrorWithStack(ctx, err, "error convert value")
		return
	}

	if values.OriginAccountNumber == "" {
		err = errHelper.OriginAccountNumberRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if values.DestinationAccountNumber == "" {
		err = errHelper.DestinationAccountNumberRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if values.OriginDescription == "" {
		err = errHelper.OriginDescriptionRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if values.DestinationDescription == "" {
		err = errHelper.DestinationDescriptionRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	return nil
}

func (service settingService) checkSettingIncentiveBudgetApproval(ctx context.Context, data string) (err error) {
	values := settingEntity.SettingIncentiveBudgetApproval{}
	err = interfaceHelper.ConvertStrToInterface(data, &values)
	if err != nil {
		err = errHelper.InvalidDataSetting.Error
		logger.ErrorWithStack(ctx, err, "error convert value")
		return
	}

	if values.DebitAccountNumber == "" {
		err = errHelper.DebitAccountNumberRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	if values.CreditAccountNumber == "" {
		err = errHelper.CreditAccountNumberRequired.Error
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidationValue.Message)
		return
	}

	return nil
}

func (service settingService) checkData(ctx context.Context, input *settingEntity.Setting) (err error) {
	oldData, err := service.settingRepo.FindOne(ctx, &settingRequest.SettingParameter{
		Name: input.Name,
	})

	if err != nil && err != sql.ErrNoRows {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorFindOne.Message)
		return
	}

	if oldData.ID != 0 && oldData.ID != input.ID {
		err = errHelper.SettingRegistered.Error
		logger.ErrorWithStack(ctx, err, "name registered")
		return
	}

	if input.Name == settingEntity.SettingTypeDefault {
		err = service.checkSettingDefault(ctx, input.Value)
		if err != nil {
			logger.ErrorWithStack(ctx, err, "check setting default")
			return
		}
	} else if input.Name == settingEntity.SettingTypeIncentiveBudget {
		err = service.checkSettingIncentiveBudget(ctx, input.Value)
		if err != nil {
			logger.ErrorWithStack(ctx, err, "check setting incentive budget")
			return
		}
	} else if input.Name == settingEntity.SettingTypeIncentiveBudgetApproval {
		err = service.checkSettingIncentiveBudgetApproval(ctx, input.Value)
		if err != nil {
			logger.ErrorWithStack(ctx, err, "check setting incentive budget approval")
			return
		}
	} else {
		err = errHelper.InvalidName.Error
		logger.ErrorWithStack(ctx, err, "error invalid name")
		return
	}

	return nil
}

func (service settingService) Create(ctx context.Context, input *settingEntity.Setting) (res int, err error) {
	if err = service.checkData(ctx, input); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCheckData.Message)
		return
	}

	tx, err := service.settingRepo.CreateTx(ctx)
	if err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCreateTx.Message)
		return
	}
	defer tx.Rollback()

	res, err = service.settingRepo.Create(ctx, tx, input)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "create setting query")
		return
	}

	tx.Commit()

	return
}

func (service settingService) checkDetails(ctx context.Context, input *settingEntity.Setting) (res settingEntity.Setting, err error) {
	res, err = service.settingRepo.FindOne(ctx, &settingRequest.SettingParameter{
		DefaultParameter: controllers.DefaultParameter{ID: input.ID},
	})
	if err != nil {
		err = errHelper.DataNotFound.Error
		logger.ErrorWithStack(ctx, err, "old data not found")
		return
	}

	return
}

func (service settingService) Update(ctx context.Context, input *settingEntity.Setting) (err error) {
	dataSetting, err := service.checkDetails(ctx, input)
	if err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCheckDetails.Message)
		return
	}

	if err = service.checkData(ctx, input); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCheckData.Message)
		return
	}

	tx, err := service.settingRepo.CreateTx(ctx)
	if err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCreateTx.Message)
		return
	}
	defer tx.Rollback()

	err = service.settingRepo.Update(ctx, tx, input)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "update setting query")
		return
	}
	tx.Commit()

	// delete redis
	go service.redis.Del(context.Background(), redis.Setting+dataSetting.Name)

	return
}

func (service settingService) Delete(ctx context.Context, id int) (err error) {
	setting := settingEntity.Setting{ID: id}
	dataSetting, err := service.checkDetails(ctx, &setting)
	if err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCheckDetails.Message)
		return
	}

	tx, err := service.settingRepo.CreateTx(ctx)
	if err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorCreateTx.Message)
		return
	}
	defer tx.Rollback()

	err = service.settingRepo.Delete(ctx, tx, id)
	if err != nil {
		logger.ErrorWithStack(ctx, err, "delete setting query")
		return
	}

	tx.Commit()

	// delete redis
	go service.redis.Del(context.Background(), redis.Setting+dataSetting.Name)

	return
}
