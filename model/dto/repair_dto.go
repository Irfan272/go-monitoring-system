package dto

import "time"

type RepairRequestDTO struct {
	RepairID          string    `json:"repair_id"`
	Title             string    `json:"title"`
	Equipment         string    `json:"equipment_id"`
	UserID            string    `json:"user_id"`
	Description       string    `json:"description"`
	Priority          string    `json:"priority"`
	LocationEquipment string    `json:"location_equipment"`
	Image_Damage      string    `json:"image_damage"`
	WorkingDate       time.Time `json:"working_date"`
	EstimasionDate    time.Time `json:"estimasion_date"`
	ActualDate        time.Time `json:"actual_date"`
	RepairStatus      string    `json:"repair_status"`
	Solusion          string    `json:"solusion "`
	TechicianId       string    `json:"techician_id"`
	Image_Repair      string    `json:"image_repair"`
	IsDeleted         bool      `json:"is_deleted"`
}
