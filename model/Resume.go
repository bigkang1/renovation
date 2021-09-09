package model

import (
	"renovation/utils/errmsg"
	"gorm.io/gorm"
)

//简历模型
type ResumePost struct {
	Job Job `gorm:"foreignkey:Jid"`
	gorm.Model
	Name        string `gorm:"type:varchar(100);not null" json:"name"`
	Jid          int    `gorm:"type:int;not null" json:"jid"`
	Telephone    string `gorm:"type:varchar(200);not null" json:"telephone"`
	Content      string `gorm:"type:longtext" json:"content"`
}


func CreateRes(data *ResumePost) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}


func GetResInfo(id int) (ResumePost, int) {
	var resume ResumePost
	err = db.Where("id = ?", id).Preload("Job").First(&resume).Error
	db.Model(&resume).Where("id = ?", id)
	if err != nil {
		return resume, errmsg.ERROR_RES_NOT_EXIST
	}
	return resume, errmsg.SUCCSE
}

//查询所有简历
func GetRes( pageSize int, pageNum int) ([]ResumePost, int, int64) {
	var resumeList []ResumePost
	var err error
	var total int64

	err = db.Select("resume.id, name, created_at, telephone, content, category.station").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Job").Find(&resumeList).Error
	// 单独计数
	db.Model(&resumeList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return resumeList, errmsg.SUCCSE, total

}

/*func SearchRes(title string, pageSize int, pageNum int) ([]Resume, int, int64) {
	var resumeList []Resume
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, , job.station").Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%",
	).Find(&resumeList).Count(&total).Error
	// 单独计数
	//db.Model(&articleList).Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return resumeList, errmsg.SUCCSE, total
}*/

/*func EditRes(id int, data *Resume) int {
	var resume Resume
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&resume).Where("id = ? ", id).Updates(&maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}*/

func DeleteRes(id int) int {
	var resume ResumePost
	err = db.Where("id = ? ", id).Delete(&resume).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//查询岗位下的所有简历
func GetJobRes(id int, pageSize int, pageNum int) ([]ResumePost, int, int64) {
	var resumeList []ResumePost
	var total int64

	err = db.Preload("Job").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"jid =?", id).Find(&resumeList).Error
	db.Model(&resumeList).Where("jid =?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_RES_NOT_EXIST, 0
	}
	return resumeList, errmsg.SUCCSE, total
}
