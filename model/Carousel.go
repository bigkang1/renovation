package model

import (
	"gorm.io/gorm"
	"renovation/utils/errmsg"
)

type Carousel struct {
	Commodity Commodity `gorm:"foreignkey:Cid"`
	ID      uint   `gorm:"primary_key;auto_increment" json:"id"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
}


func CreateCarousel(data *Carousel) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}


/*func GetCarouselInfo(id int) (Carousel,int) {
	var carousel Carousel
	db.Where("id = ?",id).First(&carousel)
	return carousel,errmsg.SUCCSE
}*/

type Carousels struct {
	Id      uint   `json:"id"`
	Cid     int    `json:"cid"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

func GetCarousels(pageSize int, pageNum int) ([]Carousels) {
	var carousels []Carousels
	//var carousel []Carousel
	//var total int64

	//db.Model(&carousel).Count(&total)
	//err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&carousel).Error
	err = db.Raw("SELECT c.id as cid, c.name, c.type ,l.id FROM commodity c INNER JOIN carousel l ON c.id = l.cid").Scan(&carousels).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return carousels
}

/*//  EditNotice 编辑公告信息
func EditNotice(id int, data *carousel) int {
	var notice carousel
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["content"] = data.Content

	err = db.Model(&notice).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}*/


func DeleteCarousel(id int) int {
	var carousel Carousel
	err = db.Where("id = ? ", id).Delete(&carousel).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
