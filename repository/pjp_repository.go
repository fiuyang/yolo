package repository

import "scyllax-pjp/model"

type PjpRepository interface {
	Save(pjp model.Pjp)
	Update(pjp model.Pjp)
	Delete(pjpId int)
	FindById(pjpId int) (pjps model.Pjp, err error)
	FindAll(limit int, page int, filters map[string]interface{}) ([]model.Pjp)
	Count() (int64, error)
}