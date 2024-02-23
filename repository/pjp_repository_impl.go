package repository

import (
	"errors"
	"fmt"
	"scyllax-pjp/data/request"
	"scyllax-pjp/data/response"
	"scyllax-pjp/helper"
	"scyllax-pjp/model"

	"gorm.io/gorm"
)

type PjpRepositoryImpl struct {
	Db *gorm.DB
}

func NewPjpRepositoryImpl(Db *gorm.DB) PjpRepository {
	return &PjpRepositoryImpl{Db: Db}
}

func (repo *PjpRepositoryImpl) Delete(pjpId int) {
	var pjp model.Pjp
	result := repo.Db.Where("id = ?", pjpId).Delete(&pjp)
	helper.ErrorPanic(result.Error)
}

func (repo *PjpRepositoryImpl) FindById(pjpId int) (pjps model.Pjp, err error) {
	var pjp model.Pjp
	result := repo.Db.Find(&pjp, pjpId)

	if result.Error != nil {
		return pjps, result.Error
	}
	
	if result.RowsAffected == 0 {
		return pjps, errors.New("pjp is not found")
	}

	return pjp, nil
}

func (repo *PjpRepositoryImpl) Save(pjp model.Pjp) {
	result := repo.Db.Create(&pjp)
	helper.ErrorPanic(result.Error)
}

func (repo *PjpRepositoryImpl) Update(pjp model.Pjp) {
	var updatePjp = request.UpdatePjpRequest{
		ID:      pjp.ID,
		PjpCode: pjp.PjpCode,
	}

	result := repo.Db.Model(&pjp).Updates(updatePjp)
	helper.ErrorPanic(result.Error)
}

func (repo *PjpRepositoryImpl) FindAll(limit int, page int, filters map[string]interface{}) []model.Pjp {
	var pjp []model.Pjp
	var totalRows int64

	query := repo.Db.Model(&pjp)

	for field, value := range filters {
		switch v := value.(type) {
		case string:
			if v != "" {
				query = query.Where(fmt.Sprintf("%s = ?", field), v)
			}
		case int:
			if v != 0 {
				query = query.Where(fmt.Sprintf("%s = ?", field), v)
			}
		}
	}

	result := query.Scopes(response.Scopes(page, limit)).Find(&pjp).Count(&totalRows)
	helper.ErrorPanic(result.Error)

	return pjp
}

func (repo *PjpRepositoryImpl) Count() (int64, error) {
	var pjp []model.Pjp
	var totalRows int64
	result := repo.Db.Find(&pjp).Count(&totalRows)
	helper.ErrorPanic(result.Error)
	return totalRows, nil
}
