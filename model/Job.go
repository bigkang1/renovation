package model

import (
	"gorm.io/gorm"
	"renovation/utils/errmsg"
)

//岗位招聘模型
type Job struct {
	ID           uint   `gorm:"primary_key;auto_increment" json:"id"`
	Station      string `gorm:"type:varchar(100);not null" json:"station"`
	Duty         string `gorm:"type:longtext" json:"duty"`
	Requirement      string `gorm:"type:longtext" json:"requirement"`
}


func CheckJob(station string) (code int) {
	var job Job
	db.Select("id").Where("station = ?", station).First(&job)
	if job.ID > 0 {
		return errmsg.ERROR_JOBNAME_USED
	}
	return errmsg.SUCCSE
}

func CreateJob(data *Job) int {
	err := db.Create(&data).Error
	if err != nil {
		println(err)
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

func GetJobInfo(id int) (Job,int) {
	var job Job
	db.Where("id = ?",id).First(&job)
	return job,errmsg.SUCCSE
}


func GetJobs(pageSize int, pageNum int) ([]Job, int64) {
	var job []Job
	var total int64
	err = db.Find(&job).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&job).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return job, total
}


func EditJob(id int, data *Job) int {
	var job Job
	var maps = make(map[string]interface{})
	maps["station"] = data.Station
	maps["duty"] = data.Duty
	maps["requirement"] = data.Requirement

	err = db.Model(&job).Where("id = ? ", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}


func DeleteJob(id int) int {
	var job Job
	err = db.Where("id = ? ", id).Delete(&job).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
