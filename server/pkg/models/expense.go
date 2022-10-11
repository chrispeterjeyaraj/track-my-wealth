package models

type Expense struct {
	Id            int    `json:"id,omitempty" gorm:"primaryKey"`
	ExpenseTitle  string `json:"expensetitle,omitempty"`
	ExpenseAmount string `json:"expenseamount,omitempty"`
	Month         string `json:"month,omitempty"`
	Year          string `json:"year,omitempty"`
	Category      string `json:"category,omitempty"`
}
