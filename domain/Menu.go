package domain

type Menu struct {
	Model
	Icon    string `gorm:"Type:Varchar(30)" json:"icon"`
	MenuID  *int64 `gorm:"INDEX" json:"-"`
	SubMenu []Menu `json:"sub_menu, omitEmpty"`
	PageID  *int64 `gorm:"INDEX" json:"-"`
	Page    Page   `json:"page"`
}
