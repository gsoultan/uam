package domain

type Page struct {
	Model
	Name     string `gorm:"Type:Varchar(255);UNIQUE_INDEX;UNIQUE;Not Null" json:"name"`
	Path     string `gorm:"Type:Varchar(500);Not Null" json:"path"`
	Pages    []Page `json:"pages"`
	PageID   int64  `gorm:"INDEX" json:"-"`
	SubPage  []Page `json:"sub_page"`
	Type     string `gorm:"Type:Varchar(10);Not Null" json:"type"` //Form, View, Table
	ModuleID int64  `gorm:"INDEX" json:"-"`
}
