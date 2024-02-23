package model

import "time"

type Pjp struct {
	ID            int       `gorm:"type:int;primary_key"`
	PjpCode       int       `gorm:"column:pjp_code;type:int;uniqueIndex;not null"`
	SaleType      string    `gorm:"column:sale_type;type:varchar(125);not null"`
	TeamSalesMan  string    `gorm:"column:team_salesman;type:varchar(125)"`
	SalesManID    string    `gorm:"column:salesman_id"`
	SalesmanName  string    `gorm:"column:salesman_name;type:varchar(125)"`
	WarehouseID   string    `gorm:"column:warehouse_id"`
	WarehouseName string    `gorm:"column:warehouse_name;type:varchar(125)"`
	PjpMode       string    `gorm:"column:pjp_mode;type:varchar(125);null"`
	Status        string    `gorm:"column:status;type:varchar(125)"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type Tabler interface {
	TableName() string
}

func (Pjp) TableName() string {
	return "permanent_journey_plans"
}
