package model

import (
	"gorm.io/gorm"
	"renovation/utils/errmsg"
)

type Commodity struct {
	ID          uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Type        string `gorm:"type:varchar(20);not null" json:"type"`
	//Content      string `gorm:"type:longtext" json:"content"`
}

func CreateComm(data *Commodity) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

func GetComm(pageSize int, pageNum int, ctype string) ([]Commodity, int64) {
	var commodity []Commodity
	var total int64
	db.Model(&commodity).Where("Type = ?",ctype).Count(&total)
	err = db.Where("Type = ?",ctype).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&commodity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return commodity, total
}

/*func GetAllCommPic(pageSize int, pageNum int,ctype string) ([]Commodity, int64) {
	var commodity []Commodity
	var total int64
	db.Model(&commodity).Where("type = ?",ctype).Count(&total)
	err = db.Select("commodity.id,commodity.name,commodity.type,commodity_pict.id,commodity_pict.img").Where("commodity.type = ?",ctype).Where("commodity_pict.cid = commodity.id").Group("commodity.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("CommodityPict").Find(&commodity).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return commodity,total
}*/

func EditComm(id int, data *Commodity) int {
	var commodity Commodity
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	//maps["type"] = data.Type

	err = db.Model(&commodity).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


func DeleteComm(id int) int {
	var commodity Commodity
	err = db.Where("id = ? ", id).Delete(&commodity).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
