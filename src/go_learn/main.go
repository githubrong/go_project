package iot_learn

import "github.com/gin-gonic/gin"

/*
go get gopkg.in/mgo.v2
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
 */

func Main() {

	router := gin.Default()
	router.GET("/learn/first",First)
	router.POST("/learn/two",Two)
    router.GET("/learn/redirect",Redirect)
	// 加载index.html模板:
	router.LoadHTMLFiles("C:/go_workspace/smart_house/api/src/templates/index1.tmpl")
	router.GET("/learn/html",Html)
	router.Run(":8090")
}
