package response

type PjpResponse struct {
	ID            int       `json:"id"`
	PjpCode       int       `json:"pjp_code"`
	SaleType      string    `json:"sale_type"`
	TeamSalesMan  string    `json:"team_salesman"`
	SalesManID    string    `json:"salesman_id"`
	SalesmanName  string    `json:"salesman_name"`
	WarehouseID   string    `json:"warehouse_id"`
	WarehouseName string    `json:"warehouse_name"`
	PjpMode       string    `json:"pjp_mode"`
	Status        string    `json:"status"`
}
