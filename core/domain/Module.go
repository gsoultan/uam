package domain

type Module struct {
	Model
	Name  string `gorm:"UNIQUE_INDEX;UNIQUE;NOT NULL;Type:"Varchar(50)" json:"name"`
	Pages []Page `json:"pages"`
}
