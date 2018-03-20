package iot_learn

//https://github.com/gin-gonic/gin
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func First(c *gin.Context){
	one := c.Query("one")
	fmt.Println("test get method")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "sucess",
		"data": one,
	})
}

func Two(c *gin.Context){
	two := c.PostForm("two")
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "sucess",
		"data": two,
	})
}

func Html(c *gin.Context)  {
    code := c.Query("code")
    url := "http://www.baidu.com"
	c.HTML(http.StatusOK, "index1.tmpl", gin.H{
		"title": "测试页面",
		"url":url,
		"code":  code,
	})
}

func Redirect(c *gin.Context)  {
	c.Redirect(http.StatusOK, "http://shanshanpt.github.io/")
}
