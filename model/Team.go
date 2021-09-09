package model

import (
	"gorm.io/gorm"
	"renovation/utils/errmsg"
)

type Team struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Img  	string `gorm:"type:varchar(100);not null" json:"img"`
}

func CreateTeam(data *Team) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}


func GetTeam(pageSize int, pageNum int) ([]Team, int64) {
	var team []Team
	var total int64
	db.Model(&team).Count(&total)
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&team).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return team, total
}

func DeleteTeam(id int) int {
	var team Team
	err = db.Where("id = ? ", id).Delete(&team).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
