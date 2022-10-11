package models

type Income struct {
	Id           int    `json:"id,omitempty" gorm:"primaryKey"`
	IncomeTitle  string `json:"incometitle,omitempty"`
	IncomeAmount string `json:"incomeamount,omitempty"`
	Month        string `json:"month,omitempty"`
	Year         string `json:"year,omitempty"`
	Category     string `json:"category,omitempty"`
}
