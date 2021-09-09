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
)

func GetAllCommPic(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	ctype := c.Query("type")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetAllCommPict(pageSize, pageNum, ctype)
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


func GetCommodity(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	ctype := c.Query("type")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetComm(pageSize, pageNum, ctype)
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


func AddCommodity(c *gin.Context)  {
	var data model.Commodity
	data.Name = c.PostForm("name")
	data.Type = c.PostForm("type")

	code = model.CreateComm(&data)

	//存入商品图片
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  errmsg.ERROR,
				"message": err,
			},
		)
	}
	files := form.File["files"]
	for _, file := range files {
		var commPic model.CommodityPict
		fileExt:=strings.ToLower(path.Ext(file.Filename))
		if fileExt!=".png"&&fileExt!=".jpg"&&fileExt!=".gif"&&fileExt!=".jpeg"{
			c.JSON(
				http.StatusOK, gin.H{
					"status":  errmsg.ERROR_FILE_FORM,
					"message": errmsg.GetErrMsg(code),
				},
			)
			return
		}
		fileName := fmt.Sprintf("%s%s","upload/",file.Filename)
		commPic.Img = fileName
		commPic.Cid = int(data.ID)
		err = c.SaveUploadedFile(file, fileName)
		if err != nil{
			c.JSON(
				http.StatusOK, gin.H{
					"status":  errmsg.ERROR_UPLOAD_FAIL,
					"message": errmsg.GetErrMsg(code),
					"err" : err,
				},
			)
			return
		}
		code = model.CreateCommodityPict(&commPic)
		if code == errmsg.ERROR{
			c.JSON(
				http.StatusOK, gin.H{
					"status":  errmsg.ERROR,
					"message": "有一张图片上传数据库失败",
				},
			)
		}
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func EditCommodity(c *gin.Context)  {
	var data model.Commodity
	id, _ := strconv.Atoi(c.Param("id"))
	data.Name = c.PostForm("name")
	//data.Type = c.PostForm("type")

	code = model.EditComm(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

func DeleteCommodity(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteComm(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}