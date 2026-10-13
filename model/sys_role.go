package model

import "gorm.io/gorm"

type SysRole struct {
	gorm.Model
	Name    string `gorm:"column:name;type:varchar(100);" josn:"name"`
	IsAdmin int8   `gorm:"column:is_admin;type:tinyint(1);" josn:"is_admin"`
	Sort    int64  `gorm:"column:sort;type:int(11);default:0;" json:"sort"`
	Remarks string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

func (s *SysRole) TableName() string {
	return "sys_role"
}

func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysRole)).Select([]string{"id", "name", "is_admin", "sort", "remarks", "created_at", "updated_at"})
	if keyword != "" {
		tx.Where("name LIKE ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}
