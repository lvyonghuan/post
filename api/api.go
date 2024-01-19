package api

import "github.com/gin-gonic/gin"

// InitRouters initializes all routes for the application.
func InitRouters() {
	r := gin.Default()

	//C端
	c := r.Group("/c")
	{
		c.PUT("/post") //投稿
		c.GET("/list") //获取投稿列表
	}

	//B端
	b := r.Group("/b")
	{
		b.GET("/get")
	}

}
