package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"renovation/model"
	"renovation/utils/errmsg"
	"path"
	"strconv"
	"strings"
	"time"
)

func GetTeams(c *gin.Context)  {
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

	data, total := model.GetTeam(pageSize, pageNum)
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

/*func GetTeamInfo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetNoticeInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}*/

func AddTeam(c *gin.Context)  {
	var data model.Team

	f, err := c.FormFile("img")
	if err != nil {
		code = errmsg.ERROR_UPLOAD_FAIL
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": errmsg.GetErrMsg(code),
			},
		)
		return
	} else {
		fileExt:=strings.ToLower(path.Ext(f.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			code = errmsg.ERROR_FILE_FORM
			c.JSON(
				http.StatusOK, gin.H{
					"status":  code,
					"message": errmsg.GetErrMsg(code),
				},
			)
			return
		}
		fileName := fmt.Sprintf("%s%d%s","upload/",time.Now().Unix(),fileExt)
		data.Img = fileName
		err = c.SaveUploadedFile(f, fileName)
		if err != nil{
			code = errmsg.ERROR_UPLOAD_FAIL
			c.JSON(
				http.StatusOK, gin.H{
					"status":  code,
					"message": errmsg.GetErrMsg(code),
				},
			)
			return
		}
	}


	code = model.CreateTeam(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}


func DeleteTeam(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteTeam(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}