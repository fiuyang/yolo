package service

import (
	"scyllax-pjp/data/request"
	"scyllax-pjp/data/response"
	"scyllax-pjp/exception"
	"scyllax-pjp/model"
	"scyllax-pjp/repository"

	_ "github.com/go-playground/validator/v10"
)

type PjpServiceImpl struct {
	PjpRepository repository.PjpRepository
}

func NewPjpServiceImpl(pjpRepository repository.PjpRepository) PjpService {
	return &PjpServiceImpl{
		PjpRepository: pjpRepository,
	}
}

func (service *PjpServiceImpl) Create(pjp request.CreatePjpRequest) {
	dataset := model.Pjp{
		PjpCode:       pjp.PjpCode,
		TeamSalesMan:  pjp.TeamSalesMan,
		SaleType:      pjp.SaleType,
		SalesManID:    pjp.SalesManID,
		SalesmanName:  "superimi",
		WarehouseID:   pjp.WarehouseID,
		WarehouseName: "superimi",
		Status:        pjp.Status,
	}

	service.PjpRepository.Save(dataset)
}

func (service *PjpServiceImpl) Delete(pjpId int) {
	service.PjpRepository.Delete(pjpId)
}

func (service *PjpServiceImpl) FindAll(limit int, page int, filters map[string]interface{}) (response.Pagination, error) {

	paginateResponse := response.Pagination{}

	result := service.PjpRepository.FindAll(limit, page, filters)

	var payload []response.PjpResponse
	for _, value := range result {
		pjp := response.PjpResponse{
			ID:            value.ID,
			PjpCode:       value.PjpCode,
			TeamSalesMan:  value.TeamSalesMan,
			SaleType:      value.SaleType,
			SalesManID:    value.SalesManID,
			SalesmanName:  value.SalesmanName,
			WarehouseID:   value.WarehouseID,
			WarehouseName: value.WarehouseName,
			Status:        value.Status,
		}
		payload = append(payload, pjp)
	}

	totalData, _ := service.PjpRepository.Count()

	paginateResponse.Data = payload
	paginateResponse.Meta.Page = page
	paginateResponse.Meta.Limit = limit
	paginateResponse.Meta.TotalData = totalData
	paginateResponse.Meta.TotalPage = totalData / int64(limit)

	return paginateResponse, nil
}

func (service *PjpServiceImpl) FindById(pjpId int) (response.PjpResponse, error) {
	dataset, err := service.PjpRepository.FindById(pjpId)

	if err != nil {
		return response.PjpResponse{}, exception.NewNotFoundError(err.Error())
	}

	response := response.PjpResponse{
		ID:            dataset.ID,
		PjpCode:       dataset.PjpCode,
		TeamSalesMan:  dataset.TeamSalesMan,
		SaleType:      dataset.SaleType,
		SalesManID:    dataset.SalesManID,
		SalesmanName:  dataset.SalesmanName,
		WarehouseID:   dataset.WarehouseID,
		WarehouseName: dataset.WarehouseName,
		Status:        dataset.Status,
	}
	return response, nil
}

func (service *PjpServiceImpl) Update(pjp request.UpdatePjpRequest) {
	dataset, err := service.PjpRepository.FindById(pjp.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	dataset.PjpCode = pjp.PjpCode
	service.PjpRepository.Update(dataset)
}
