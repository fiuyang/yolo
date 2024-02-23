package service

import (
	"scyllax-pjp/data/request"
	"scyllax-pjp/data/response"
)

type PjpService interface {
	Create(pjp request.CreatePjpRequest)
	Update(pjp request.UpdatePjpRequest)
	Delete(pjpId int)
	FindById(pjpId int) (response.PjpResponse, error)
	FindAll(limit int, page int, filters map[string]interface{}) (response.Pagination, error)
}
