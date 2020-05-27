package models

type Order struct {
	ID          int     `json:"id" gorm:"primary_key"`
	CutomerName string  `json:"cutomerName"`
	OrderAmount float64 `json:"orderAmount"`
	Items       []*Item `json:"items" gorm:"foreignKey:OrderID"`
}
