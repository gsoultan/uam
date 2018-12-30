package domain

type Role struct {
	Model
	Name          string         `gorm:"Type:Varchar(25);NOT NULL;UNIQUE_INDEX;UNIQUE" json:"name"`
	RoleInModules []RoleInModule `json:"modules"`
}

type RoleInModule struct {
	Model
	RoleID   int64 `gorm:"INDEX" json:"-"`
	Role     Role
	ModuleID int64 `gorm:"INDEX" json:"-"`
	Module   Module
	Read     bool
	Create   bool
	Edit     bool
	Delete   bool
}
