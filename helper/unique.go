package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"reflect"
	"scyllax-pjp/model"
	"strings"
)

var modelMap = map[string]reflect.Type{
	"permanent_journey_plans": reflect.TypeOf(model.Pjp{}),
}

func ValidateUnique(db *gorm.DB, fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	tableName := getModelFromTag(fl)
	fmt.Println("Validation: for table", tableName)

	exists := UniqueExistsInTable(db, value, tableName)

	return !exists
}

func UniqueExistsInTable(db *gorm.DB, value interface{}, tableName string) bool {
	parts := strings.Split(tableName, ";")
	modelName := parts[0]
	columnName := parts[1]
	modelType, ok := modelMap[modelName]
	if !ok {
		fmt.Printf("Unknown model name: %s\n", parts)
		return false
	}

	modelInstance := reflect.New(modelType).Interface()
	fmt.Printf("model instance: %s\n", value)
	if err := db.Table(modelName).Where(columnName+" = ?", value).First(modelInstance).Error; err != nil {
		return false
	}
	return true
}

func getModelFromTag(fl validator.FieldLevel) string {
	// Assuming 'validate' tag is in the format "unique=tableName;columnName"
	validateTag := fl.Param()

	parts := strings.Split(validateTag, "=")
	if len(parts) >= 2 {
		return parts[1]
	}

	return validateTag
}
