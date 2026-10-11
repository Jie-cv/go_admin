package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(50);" json:"user_name"`
	Password  string `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone     string `gorm:"column:phone;type:char(11)" json:"phone"`
	WxUnoioId string `gorm:"columm:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId  string `gorm:"columm:wx_open_id;type:varchar(255);" json:"wx_open_id"`
	Avatar    string `gorm:"columen:avatar;type:varchar(255);" json:"avatar"`
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func GetUserByUsernamePwd(username string, pwd string) (*SysUser, error) {
	user := new(SysUser)
	err := DB.Where("username = ? AND password = ?", username, pwd).First(user).Error
	return user, err
}
