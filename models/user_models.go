package models

import (
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

type SysUser struct {
	Id          int64          `json:"id"`
	Username    string         `json:"username"`
	Password    string         `json:"password"`
	Avatar      string         `json:"avatar"`
	Email       string         `json:"email"`
	Phonenumber string         `json:"phonenumber"`
	LoginDate   time.Time      `json:"login_date"`
	Status      string         `json:"status"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Remark      string         `json:"remark"`
}

func (s *SysUser) TableName() string {
	return "sys_user"
}
