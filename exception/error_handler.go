package exception

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"scyllax-pjp/data/response"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(err error, ctx *gin.Context, requestType interface{}) {
	if notFoundError(err, ctx) {
		return
	}

	var e *ValidationError
	if errors.As(err, &e) {
		validationErrors(e.err, ctx, e.requestType)
		return
	}

	internalServerError(err, ctx)
}

func validationErrors(err error, ctx *gin.Context, requestType interface{}) bool {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		report := make(map[string]string)

		for _, e := range castedObject {
			field, _ := reflect.TypeOf(requestType).FieldByName(e.StructField())
			fieldName := field.Tag.Get("json")
			switch e.Tag() {
			case "required":
				report[fieldName] = fmt.Sprintf("%s is required", fieldName)
			case "email":
				report[fieldName] = fmt.Sprintf("%s is not valid email", fieldName)
			case "gte":
				report[e.Field()] = fmt.Sprintf("%s value must be greater than %s", fieldName, e.Param())
			case "lte":
				report[fieldName] = fmt.Sprintf("%s value must be lower than %s", fieldName, e.Param())
			case "unique":
				report[fieldName] = fmt.Sprintf("%s has already been taken %s", fieldName, e.Param())
			case "max":
				report[fieldName] = fmt.Sprintf("%s value must be lower than %s", fieldName, e.Param())
			case "min":
				report[fieldName] = fmt.Sprintf("%s value must be greater than %s", fieldName, e.Param())
			case "numeric":
				report[fieldName] = fmt.Sprintf("%s value must be numeric", fieldName)
			}
		}

		ctx.JSON(http.StatusBadRequest, response.Error{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Errors: report,
		})
		return true
	}
	return false
}

func notFoundError(err interface{}, ctx *gin.Context) bool {
	exception, ok := err.(*NotFoundError)
	if ok {
		ctx.JSON(http.StatusNotFound, response.Error{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Errors: exception.Error(),
		})
		return true
	}
	return false
}

func internalServerError(err interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, response.Error{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Errors: err,
	})
}
