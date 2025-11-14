package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type SupplyUsage struct {
	ID             string          `gorm:"default:uuid_generate_v4()"`
	BatchID        string          `gorm:"column:batch_id"`
	StudentID      string          `gorm:"column:student_id"`
	ClassSessionID string          `gorm:"column:class_session_id"`
	Quantity       int             `gorm:"column:quantity"`
	UnitPrice      decimal.Decimal `gorm:"column:unit_price"`
	TotalPrice     decimal.Decimal `gorm:"column:total_price"`
	UsedAt         time.Time       `gorm:"column:used_at"`
}

func (SupplyUsage) TableName() string {
	return "supply_usages"
}
