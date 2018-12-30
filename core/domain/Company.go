package domain

type Company struct {
	Model
	Name string `gorm:"UNIQUE_INDEX UNIQUE" json:"company_name"`
}
