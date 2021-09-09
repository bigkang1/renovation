package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"renovation/model"
	"renovation/utils/errmsg"
	"strconv"
	"strings"
	"time"
)

/*func GetLogos(c *gin.Context)  {
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
}*/

func GetLogoInfo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	ltype := c.Query("type")

	data, code := model.GetLogoInfo(id,ltype)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func AddLogo(c *gin.Context)  {
	var data model.Logo
	data.Type = c.PostForm("type")

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

	code = model.CreateLogo(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}


func EditLogo(c *gin.Context)  {
	var data model.Logo
	id, _ := strconv.Atoi(c.Param("id"))

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

	code = model.EditLogo(id,&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteLogo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteLogo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}