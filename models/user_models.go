package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Job       string    `json:"job"`
	Notes     string    `json:"notes"`
	Status    int       `json:"status"`
	Flag      int       `json:"flag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *User) TableName() string {
	return "user"
}

type UserInput struct {
	Name  string `json:"name"`
	Job   string `json:"job"`
	Notes string `json:"notes"`
}

func (ui *UserInput) TableName() string {
	return "user"
}
