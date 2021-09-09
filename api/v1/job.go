package v1

import (
	"renovation/model"
	"renovation/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//TODO 还有修改
func AddJob(c *gin.Context) {
	var data model.Job
	data.Station = c.PostForm("station")
	data.Duty = c.PostForm("duty")
	data.Requirement = c.PostForm("requirement")

	code = model.CheckJob(data.Station)
	if code == errmsg.SUCCSE {
		model.CreateJob(&data)
	}
	if code == errmsg.ERROR_JOBNAME_USED {
		code = errmsg.ERROR_JOBNAME_USED
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}


func GetJobInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetJobInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}

//  查询列表
func GetJob(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetJobs(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}



func EditJob(c *gin.Context) {
	var data model.Job
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckJob(data.Station)
	if code == errmsg.SUCCSE {
		model.EditJob(id, &data)
	}
	if code == errmsg.ERROR_JOBNAME_USED {
		c.Abort()
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}


func DeleteJob(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteJob(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
