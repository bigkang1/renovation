package model

import (
	"renovation/utils/errmsg"
)

type Logo struct {
	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Img  	string `gorm:"type:varchar(100);not null" json:"img"`
	Type    string `gorm:"type:varchar(20);not null" json:"type"`
}

func GetLogoInfo(id int,ltype string) (Logo,int) {
	var logo Logo
	db.Where("id = ?",id).Where("type = ?",ltype).First(&logo)
	return logo,errmsg.SUCCSE
}

func CreateLogo(data *Logo) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}


func EditLogo(id int, data *Logo) int {
	var logo Logo
	var maps = make(map[string]interface{})
	maps["img"] = data.Img

	err = db.Model(&logo).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


/*func GetLogo(pageSize int, pageNum int) ([]Logo, int64) {
	var logo []Logo
	var total int64
	db.Model(&logo).Count(&total)
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&logo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return logo, total
}*/

func DeleteLogo(id int) int {
	var logo Logo
	err = db.Where("id = ? ", id).Delete(&logo).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
