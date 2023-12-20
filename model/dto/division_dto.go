package dto

type DivisionRequestDTO struct {
	DivisonID    string `json:"division_id"`
	NameDivision string `json:"name_division"`
	Description  string `json:"description"`
	IsDeleted    bool   `json:"is_deleted"`
}
