package model

import (
	"gorm.io/gorm"
	"renovation/utils/errmsg"
)

type CommodityPict struct {
	Commodity Commodity `gorm:"foreignkey:Cid"`
	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Img  	string `gorm:"type:varchar(100);not null" json:"img"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
}


func CreateCommodityPict(data *CommodityPict) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

//查询商品下的所有图片
func GetCommPic(id int, pageSize int, pageNum int) ([]CommodityPict, int, int64) {
	var commPicList []CommodityPict
	var total int64
	db.Model(&commPicList).Where("cid =?", id).Count(&total)
	err = db.Preload("Commodity").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&commPicList).Error
	if err != nil {
		return nil, errmsg.ERROR_JOB_NOT_EXIST, 0
	}
	return commPicList, errmsg.SUCCSE, total
}

func DeleteCommodityPict(id int) int {
	var commodityPict CommodityPict
	err = db.Where("id = ? ", id).Delete(&commodityPict).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

type CommodityPicts struct {
	Name        string `json:"name"`
	Type        string `json:"type"`

	ID      uint   `json:"id"`
	Img  	string `json:"img"`
	Cid     int    ` json:"cid"`
}

func GetAllCommPict(pageSize int, pageNum int,ctype string) ([]CommodityPicts, int64) {
	var commodityPicts []CommodityPicts
	var commodity []Commodity
	var total int64
	db.Model(&commodity).Where("type = ?",ctype).Count(&total)
	//err = db.Select("commodity.id,commodity.name,commodity.type,commodity_pict.id,commodity_pict.img").Limit(pageSize).Offset((pageNum - 1) * pageSize).Preload("Commodity").Find(&commodityp).Error
	//.Where("commodity.type = ?",ctype).Group("commodity.name")
	db.Raw("SELECT c.id as cid, c.name, c.type, p.id, p.img FROM commodity c INNER JOIN commodity_pict p ON c.id = p.cid where c.type = ? ORDER BY cid;",ctype).Scan(&commodityPicts)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return commodityPicts,total
}