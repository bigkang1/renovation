package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"renovation/model"
	"renovation/utils/errmsg"
	"strconv"
)

func GetCarousels(c *gin.Context)  {
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

	data := model.GetCarousels(pageSize, pageNum)
	code = errmsg.SUCCSE
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			//"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

/*func GetCarouselInfo(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCarouselInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}*/

func AddCarousel(c *gin.Context)  {
	var data model.Carousel
	data.Cid,_ = strconv.Atoi(c.PostForm("cid"))

	code = model.CreateCarousel(&data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

/*func EditCarousel(c *gin.Context)  {
	var data model.Notice
	id, _ := strconv.Atoi(c.Param("id"))
	data.Content = c.PostForm("content")
	data.Title = c.PostForm("title")

	code = model.EditNotice(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}*/

func DeleteCarousel(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCarousel(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}