package models

import "time"

type Activities struct {
	Id             int       `json:"id"`
	FkUserID       int       `json:"fk_user_id"`
	FkCategoriesID int       `json:"fk_categories_id"`
	ExpenseDate    time.Time `json:"expense_date"`
	Expense        int       `json:"expense"`
	TotalExpense   int       `json:"total_expense"`
	Notes          string    `json:"notes"`
	Status         int       `json:"status"`
	Flag           int       `json:"flag"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (a *Activities) TableName() string {
	return "activities"
}
