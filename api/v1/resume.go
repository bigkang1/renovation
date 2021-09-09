package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
	"renovation/model"
	"renovation/utils/errmsg"
	"strconv"
)

type ResV struct {
	Name        string `json:"name"`
	Jid      string `json:"jid"`
	Telephone    string `json:"telephone"`
	Content      string `json:"content"`
	Vid        string  `json:"vid"`
	Val        string  `json:"val"`
}

func AddResume(c *gin.Context) {
	var data ResV
	_ = c.ShouldBindJSON(&data)

	if data.Vid == "" || data.Val == ""{
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_VALIDATE_NOTNULL,
			"message": errmsg.GetErrMsg(errmsg.ERROR_VALIDATE_NOTNULL),
		})
		return
	}
	// 同时在内存清理掉这个图片
	err := store.Verify(data.Vid, data.Val, true)
	if !err {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR_VALIDATE_CAPTCHA,
			"message": errmsg.GetErrMsg(errmsg.ERROR_VALIDATE_CAPTCHA),
		})
		return
	}
	var res model.ResumePost
	res.Name = data.Name
	res.Telephone = data.Telephone
	res.Content = data.Content
	res.Jid,_ = strconv.Atoi(data.Jid)

	code = model.CreateRes(&res)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//  查询岗位下的所有简历
func GetJobRes(c *gin.Context) {
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

	data, code, total := model.GetJobRes(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个
func GetResInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetResInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询列表
func GetRes(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	//name := c.Query("name")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, code, total := model.GetRes(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
	return

	/*data, code, total := model.SearchRes(name,pageSize,pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})*/
}


func DeleteRes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteRes(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}


var store = base64Captcha.DefaultMemStore

//  获取验证码
func GetCaptcha(c *gin.Context){
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	a := base64Captcha.NewCaptcha(driver, store)

	// 获取
	id, b64s, err := a.Generate()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  errmsg.ERROR,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCSE,
		"vid": id,
		"b64": b64s,
		"ceshi1":driver,
		"ceshi2":store,
		"message": errmsg.GetErrMsg(code),
	})
}
