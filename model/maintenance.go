package model

import "time"

type Maintenance struct {
	MaintenanceID     string    `json:"maintenance_id"`
	Equipment         Equipment `json:"equipment_id"`
	Priority          string    `json:"priority"`
	WorkingDate       time.Time `json:"working_date"`
	EstimasionDate    time.Time `json:"estimasion_date"`
	ActualDate        time.Time `json:"actual_date"`
	TechicianId       User      `json:"techician_id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	Solusion          string    `json:"solusion"`
	ImageMaintenance  string    `json:"image_maintenance"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	IsDeleted         bool      `json:"is_deleted"`
}
