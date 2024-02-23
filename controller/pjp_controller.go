package controller

import (
	"net/http"
	"scyllax-pjp/data/request"
	"scyllax-pjp/data/response"
	"scyllax-pjp/exception"
	"scyllax-pjp/helper"
	"scyllax-pjp/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type PjpController struct {
	pjpService service.PjpService
	Validate   *validator.Validate
}

func NewPjpController(service service.PjpService, validate *validator.Validate) *PjpController {
	return &PjpController{
		pjpService: service,
		Validate:   validate,
	}
}

// CreateTags		godoc
// @Summary			Create pjp
// @Description		Save pjp data in Db.
// @Param			pjp body request.CreatePjpRequest true "Create pjp"
// @Produce			application/json
// @Tags			pjp
// @Success			200 {object} response.Response{}
// @Router			/pjp [post]
func (controller *PjpController) Create(ctx *gin.Context) {
	log.Info().Msg("create pjp")
	createPjpRequest := request.CreatePjpRequest{}

	err := ctx.ShouldBindJSON(&createPjpRequest)
	helper.ErrorPanic(err)

	if err := controller.Validate.Struct(createPjpRequest); err != nil {
		exception.ErrorHandler(err, ctx, createPjpRequest)
		return
	}

	controller.pjpService.Create(createPjpRequest)
	webResponse := response.Response{
		Code:    http.StatusCreated,
		Status:  "Ok",
		Message: "Created successfully",
		Data:    nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, webResponse)
}

// UpdateTags		godoc
// @Summary			Update pjp
// @Description		Update pjp data.
// @Param			pjpId path string true "update pjp by id"
// @Param			pjp body request.UpdatePjpRequest true  "Update pjp"
// @Tags			pjp
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/pjp/{pjpId} [patch]
func (controller *PjpController) Update(ctx *gin.Context) {
	log.Info().Msg("update pjp")
	updatePjpRequest := request.UpdatePjpRequest{}
	err := ctx.ShouldBindJSON(&updatePjpRequest)
	helper.ErrorPanic(err)

	if err := controller.Validate.Struct(updatePjpRequest); err != nil {
		exception.ErrorHandler(err, ctx, updatePjpRequest)
		return
	}

	pjpId := ctx.Param("pjpId")
	id, err := strconv.Atoi(pjpId)
	helper.ErrorPanic(err)
	updatePjpRequest.ID = id

	controller.pjpService.Update(updatePjpRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags		godoc
// @Summary			Delete pjp
// @Param		    pjpId path string true "delete pjp by id"
// @Description		Remove pjp data by id.
// @Produce			application/json
// @Tags			pjp
// @Success			200 {object} response.Response{}
// @Router			/pjp/{pjpId} [delete]
func (controller *PjpController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete pjp")
	pjpId := ctx.Param("pjpId")
	id, err := strconv.Atoi(pjpId)
	helper.ErrorPanic(err)
	controller.pjpService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTags 		godoc
// @Summary				Get Single pjp by id.
// @Param				pjpId path string true "update pjp by id"
// @Description			Return the pjp whoes pjpId value mathes id.
// @Produce				application/json
// @Tags				pjp
// @Success				200 {object} response.Response{}
// @Router				/pjp/{pjpId} [get]
func (controller *PjpController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid pjp")
	pjpId := ctx.Param("pjpId")
	id, err := strconv.Atoi(pjpId)
	helper.ErrorPanic(err)

	// data := controller.pjpService.FindById(id)
	data, err := controller.pjpService.FindById(id)
	if err != nil {
		exception.ErrorHandler(err, ctx, nil)
		return
	}
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   data,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All pjp.
// @Description		Return list of pjp.
// @Param		    limit query string false "Limit"
// @Param		    page query string false "Page"
// @Param		    pjp_code query string false "Pjp Code"
// @Produce		    application/json
// @Tags			pjp
// @Success         200 {object} response.Pagination{}
// @Router			/pjp [get]
func (controller *PjpController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll pjp")
	limit, _ := strconv.Atoi(ctx.Query("limit"))

	page, _ := strconv.Atoi(ctx.Query("page"))

	filters := make(map[string]interface{})
	filters["pjp_code"] = ctx.Query("pjp_code")

	if limit < 1 {
		limit = 10
	}

	if page < 1 {
		page = 1
	}
	responses, err := controller.pjpService.FindAll(limit, page, filters)
	helper.ErrorPanic(err)

	webResponse := response.Pagination{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   responses.Data,
		Meta:   responses.Meta,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
