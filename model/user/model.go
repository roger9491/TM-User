package user


type User struct {
	ID			int64 		`gorm:"index" json:"id"`
	UserName    string     	`gorm:"column:username;not null" json:"username"`
	Password    string     	`gorm:"column:password;not null" json:"password"`
}

// 覆蓋預設表名
func (u *User) TableName() string {
    return "user"
}

// 回傳user訊息
type UserInfo struct {
	UserID 		int64	`json:"userid"`
	JwtToken	string  `json:"jwttoken"`
}
