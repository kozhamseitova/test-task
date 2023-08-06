package entity

type UserFilter struct {
	BirthDateAsc *bool   `json:"birth_date_asc" query:"birth_date_asc"`
	City         *string `json:"city" query:"city"`
	Search       *string `json:"search"`
	Page         int     `json:"page" query:"page" binding:"required"`
	Amount       int     `json:"amount" query:"amount" binding:"required"`
}
