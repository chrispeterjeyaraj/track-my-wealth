package models

type Saving struct {
	Id           int    `json:"id,omitempty" gorm:"primaryKey"`
	SavingTitle  string `json:"savingtitle,omitempty"`
	SavingAmount string `json:"savingamount,omitempty"`
	Month        string `json:"month,omitempty"`
	Year         string `json:"year,omitempty"`
	Category     string `json:"category,omitempty"`
}
