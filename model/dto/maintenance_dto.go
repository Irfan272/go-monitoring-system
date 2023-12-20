package dto

import "time"

type MaintenanceRequestDTO struct {
	MaintenanceID     string    `json:"maintenance_id"`
	Equipment         string    `json:"equipment_id"`
	Priority          string    `json:"priority"`
	WorkingDate       time.Time `json:"working_date"`
	EstimasionDate    time.Time `json:"estimasion_date"`
	ActualDate        time.Time `json:"actual_date"`
	TechicianId       string    `json:"techician_id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	Solusion          string    `json:"solusion"`
	ImageMaintenance  string    `json:"image_maintenance"`
	IsDeleted         bool      `json:"is_deleted"`
}
