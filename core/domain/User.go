package domain

type User struct {
	Model
	UserName  string `gorm:"Type:Varchar(200);Not Null;UNIQUE_INDEX;UNIQUE" json:"user_name, omitEmpty"`
	Email     string `gorm:"Type:Varchar(200);Not Null;UNIQUE_INDEX;UNIQUE" json:"email, omitEmpty"`
	Password  string `gorm:"Type:Varchar(50);Not Null" json:"-"`
	Hash      string `gorm:"TYpe:Varchar(50);Not Null" json:"-"`
	Mobile    string `gorm:"Type:Varchar(250);Not Null;UNIQUE_INDEX;UNIQUE" json:"mobile, omitEmpty"`
	Token     string `gorm:"Type:Varchar(250)" json:"token, omitEmpty"`
	RoleID    *int64 `gorm:"INDEX" json:"-"`
	Role      Role
	AccountID int64   `gorm:"INDEX" json:"-"`
	Account   Account `json:"account, omitEmpty"`
}
