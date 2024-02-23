package request

type CreatePjpRequest struct {
	PjpCode      int    `validate:"required,numeric,unique=permanent_journey_plans;pjp_code" json:"pjp_code"`
	TeamSalesMan string `validate:"required"   json:"salesman_team"`
	SaleType     string `validate:"required"   json:"sale_type"`
	SalesManID   string `validate:"required"   json:"salesman_id"`
	WarehouseID  string `validate:"required"   json:"warehouse_id"`
	Status       string `validate:"required"   json:"is_active"`
}

type UpdatePjpRequest struct {
	ID           int    `validate:"required"`
	PjpCode      int    `validate:"required,numeric,max=4"   json:"pjp_code"`
	TeamSalesMan string `validate:"required"   json:"salesman_team"`
	SaleType     string `validate:"required"   json:"sale_type"`
	SalesManID   string `validate:"required"   json:"salesman_id"`
	WarehouseID  string `validate:"required"   json:"warehouse_id"`
	Status       string `validate:"required"   json:"is_active"`
}
