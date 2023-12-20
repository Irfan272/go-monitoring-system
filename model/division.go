package model

import "time"

type Division struct {
	DivisonID    string    `json:"division_id"`
	NameDivision string    `json:"name_division"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	IsDeleted    bool      `json:"is_deleted"`
}
