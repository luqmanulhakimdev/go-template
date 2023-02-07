package setting

import (
	"context"
	"database/sql"
	"encoding/json"
	constantController "go-template/controllers/constants"
	settingEntity "go-template/entities/setting"
	errHelper "go-template/helpers/errors"
	"go-template/helpers/str"
	"go-template/logger"
	"go-template/response"
	"go-template/services"
	"net/http"

	"github.com/gorilla/mux"
)

type SettingService interface {
	SelectAll(context.Context, *SettingParameter) ([]settingEntity.Setting, error)
	FindAll(context.Context, *SettingParameter) ([]settingEntity.Setting, services.Pagination, error)
	FindOne(context.Context, *SettingParameter) (settingEntity.Setting, error)
	Create(ctx context.Context, input *settingEntity.Setting) (int, error)
	Update(ctx context.Context, input *settingEntity.Setting) error
	Delete(ctx context.Context, id int) error
}

type settingController struct {
	settingServices SettingService
}

func NewSettingController(settingServices SettingService) settingController {
	return settingController{settingServices: settingServices}
}

func (ctrl settingController) InitializeRoutes(userRouter *mux.Router, adminRouter *mux.Router, staticRouter *mux.Router) {
	// static
	staticRouter.HandleFunc("/all", ctrl.SelectAll).Methods(http.MethodGet)
	staticRouter.HandleFunc("", ctrl.FindAll).Methods(http.MethodGet)
	staticRouter.HandleFunc("/one", ctrl.FindOne).Methods(http.MethodGet)
	staticRouter.HandleFunc("", ctrl.Create).Methods(http.MethodPost)
	staticRouter.HandleFunc(constantController.ARGUMENT_ID, ctrl.Update).Methods(http.MethodPut)
	staticRouter.HandleFunc(constantController.ARGUMENT_ID, ctrl.Delete).Methods(http.MethodDelete)
}

// Select All Setting
// @Summary      Show All Settings
// @Description  Get all setting data from database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        search   	query	string  false  "Search by setting name"
// @Param        order_by   query   string  false  "Order data"
// @Param        sort   	query   string  false  "Sort data"
// @Success      200  {object}  object{success=bool,data=[]SettingResponse{}}
// @Failure      400  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      401  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      404  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Router       /v1/apiStatic/setting/all [get]
func (ctrl settingController) SelectAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error
	parameters := new(SettingParameter)
	parameters.Search = r.FormValue("search")
	parameters.OrderBy = r.FormValue("order_by")
	parameters.Sort = r.FormValue("sort")
	logger.Info(ctx, "Select All Total: %v", parameters)

	result, err := ctrl.settingServices.SelectAll(ctx, parameters)
	if err != nil && err != sql.ErrNoRows { // no rows error come from count query on find all
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, FromServices(result), nil)
}

// Find All Setting
// @Summary      Show All Settings with Pagination
// @Description  Get all setting data from database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        page   	query	string  false  "page"
// @Param        limit   	query	string  false  "limit data"
// @Param        search   	query	string  false  "Search by setting name"
// @Param        order_by   query   string  false  "Order data"
// @Param        sort   	query   string  false  "Sort data"
// @Success      200  {object}  object{success=bool,data=[]SettingResponse{},meta=services.Pagination{}}
// @Failure      400  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      401  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      404  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Router       /v1/apiAdmin/setting [get]
func (ctrl settingController) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error
	parameters := new(SettingParameter)
	parameters.Search = r.FormValue("search")
	parameters.Page = str.StringToInt(r.FormValue("page"))
	parameters.Limit = str.StringToInt(r.FormValue("limit"))
	parameters.OrderBy = r.FormValue("order_by")
	parameters.Sort = r.FormValue("sort")
	logger.Info(ctx, "Find All Total: %v", parameters)

	result, meta, err := ctrl.settingServices.FindAll(ctx, parameters)
	if err != nil && err != sql.ErrNoRows { // no rows error come from count query on find all
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, FromServices(result), meta)
}

// Find One Setting
// @Summary      Find One Setting
// @Description  Get one setting data from database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   	query	string  false  "setting id"
// @Param        name   query   string  false  "setting name"
// @Success      200  {object}  object{success=bool,data=SettingResponse{}}
// @Failure      400  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      401  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Failure      404  {object}  object{success=bool,data=interface{},stat_msg=string}
// @Router       /v1/apiStatic/setting/one [get]
func (ctrl settingController) FindOne(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error
	parameters := new(SettingParameter)
	parameters.ID = str.StringToInt(r.FormValue("id"))
	parameters.Name = r.FormValue("name")
	logger.Info(ctx, "Find One: %v", parameters)

	result, err := ctrl.settingServices.FindOne(ctx, parameters)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, FromService(result), nil)
}

// Create Setting
// @Summary      Create Setting Data
// @Description  Create setting data to database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        name   	body   	string  		true  	"setting name"
// @Param        value  	body   	interface{} 	true  	"setting value"
// @Success      200  {object}  object{success=bool,data=int}
// @Failure      400  {object}  object{success=bool,data=int,stat_msg=string}
// @Failure      401  {object}  object{success=bool,data=int,stat_msg=string}
// @Failure      404  {object}  object{success=bool,data=int,stat_msg=string}
// @Router       /v1/apiAdmin/setting [post]
func (ctrl settingController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	reqCreateSetting := new(SettingRequest)
	if err := json.NewDecoder(r.Body).Decode(reqCreateSetting); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.CannotDecodeJson.Message)
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	if err := reqCreateSetting.Validate(); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidateRequest.Message)
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	data := reqCreateSetting.ToService()
	res, err := ctrl.settingServices.Create(ctx, data)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, res, nil)
}

// Update Setting
// @Summary      Update Setting
// @Description  Edit setting data from database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   		path	int  			true  	"setting id"
// @Param        name   	body   	string  		true  	"setting name"
// @Param        value  	body   	interface{} 	true  	"setting value"
// @Success      200  {object}  object{success=bool}
// @Failure      400  {object}  object{success=bool,stat_msg=string}
// @Failure      401  {object}  object{success=bool,stat_msg=string}
// @Failure      404  {object}  object{success=bool,stat_msg=string}
// @Router       /v1/apiStatic/setting/id/{id} [put]
func (ctrl settingController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := str.StringToInt(vars["id"])
	reqUpdateSetting := new(SettingRequest)
	if err := json.NewDecoder(r.Body).Decode(reqUpdateSetting); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.CannotDecodeJson.Message)
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	if err := reqUpdateSetting.Validate(); err != nil {
		logger.ErrorWithStack(ctx, err, errHelper.ErrorValidateRequest.Message)
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	data := reqUpdateSetting.ToService()
	data.ID = id
	err := ctrl.settingServices.Update(ctx, data)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, nil, nil)
}

// Delete Setting
// @Summary      Delete Setting
// @Description  Delete setting data from database neobank
// @Tags         API Setting
// @Security 	 ApiKeyAuth
// @Accept       json
// @Produce      json
// @Param        id   		path	int  			true  	"setting id"
// @Success      200  {object}  object{success=bool}
// @Failure      400  {object}  object{success=bool,stat_msg=string}
// @Failure      401  {object}  object{success=bool,stat_msg=string}
// @Failure      404  {object}  object{success=bool,stat_msg=string}
// @Router       /v1/apiAdmin/setting/id/{id} [delete]
func (ctrl settingController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	id := str.StringToInt(vars["id"])
	err := ctrl.settingServices.Delete(ctx, id)
	if err != nil {
		response.RespondError(w, http.StatusBadRequest, err)
		return
	}

	response.RespondSuccess(w, http.StatusOK, nil, nil)
}
