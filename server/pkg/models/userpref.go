package models

type UserPref struct {
	Id            int    `json:"id" gorm:"primaryKey"`
	Setting       string `json:"setting"`
	Configuration string `json:"configuration"`
}
