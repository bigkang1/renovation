package v1

import (
	"fmt"
	"path"
	"renovation/model"
	"renovation/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)


func AddCommodityPict(c *gin.Context) {
	var data model.CommodityPict
	data.Cid,_ = strconv.Atoi(c.PostForm("cid"))

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

	code = model.CreateCommodityPict(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//查询商品下的所有图片
func GetCommPic(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetCommPic(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteCommodityPict(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCommodityPict(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}
