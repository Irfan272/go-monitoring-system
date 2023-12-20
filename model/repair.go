package model

import "time"

type Repair struct {
	RepairID          string    `json:"repair_id"`
	Title             string    `json:"title"`
	Equipment         Equipment `json:"equipment_id"`
	User              User      `json:"user_id"`
	Description       string    `json:"description"`
	Priority          string    `json:"priority"`
	LocationEquipment string    `json:"location_equipment"`
	Image_Damage      string    `json:"image_damage"`
	WorkingDate       time.Time `json:"working_date"`
	EstimasionDate    time.Time `json:"estimasion_date"`
	ActualDate        time.Time `json:"actual_date"`
	RepairStatus      string    `json:"repair_status"`
	Solusion          string    `json:"solusion "`
	TechicianId       User      `json:"techician_id"`
	Image_Repair      string    `json:"image_repair"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	IsDeleted         bool      `json:"is_deleted"`
}
