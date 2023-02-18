package service

import "lianghj.top/admin/goadminv1/models"

func GetUserAll() ([]models.SysUser, error) {
	users := []models.SysUser{}
	err := models.DB.Model(models.SysUser{}).Find(&users).Error
	return users, err
}
