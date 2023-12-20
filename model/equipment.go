package model

import "time"

type Equipment struct {
	EquipmentID     string    `json:"equipment_id"`
	NameEquipment   string    `json:"name_equipment"`
	EquipmentBrand  string    `json:"equipment_brand"`
	Producer        string    `json:"producer"`
	Division        Division  `json:"division_id"`
	ProductionYear  time.Time `json:"production_year"`
	LimitYear       time.Time `json:"limit_year"`
	EquipmentAge    string    `json:"equipment_age"`
	CondisionStatus string    `json:"condision_status"`
	Image_Equipment string    `json:"image_equipment"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	IsDeleted       bool      `json:"is_deleted"`
}
