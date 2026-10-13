package model

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	UserName  string `gorm:"column:username;type:varchar(50);" json:"user_name"`
	Password  string `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone     string `gorm:"column:phone;type:char(11)" json:"phone"`
	WxUnoioId string `gorm:"column:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);" json:"wx_open_id"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"`
	Sex       string `gorm:"column:sex;type:varchar(1);" json:"sex"`
	Email     string `gorm:"column:email;type:varchar(20);" json:"email"`
	Remarks   string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func GetUserByUsernamePwd(username string, pwd string) (*SysUser, error) {
	user := new(SysUser)
	err := DB.Where("username = ? AND password = ?", username, pwd).First(user).Error
	return user, err
}

func GetUserList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysUser)).Select("id,username,phone,avatar,created_at,updated_at,remarks,email,sex")
	if keyword != "" {
		tx.Where("username LIKE ?", "%"+keyword+"%")
	}
	return tx
}
