package model

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Gender    string    `json:"gender"`
	Address   string    `json:"address"`
	NoHP      string    `json:"no_hp"`
	Division  Division  `json:"division_id"`
	Role      string    `json:"role"`
	ImageUser string    `json:"image_user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsDeleted bool      `json:"is_deleted"`
}

func (u User) IsValidRole() bool {
	return u.Role == "it support" || u.Role == "user" || u.Role == "manager maintenance" || u.Role == "techician"

}
