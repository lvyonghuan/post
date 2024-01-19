package api

import "github.com/gin-gonic/gin"

// GetPost 接收投稿
func GetPost(c *gin.Context) {
	authorName := c.PostForm("author_name") //作者名字
	if authorName == "" {
		//TODO
	}

	authorPhoneNum := c.PostForm("author_phone_num") //作者手机号
	if authorPhoneNum == "" {
		//TODO
	}

	authorEmail := c.PostForm("author_email") //作者邮箱
	if authorEmail == "" {
		//TODO
	}

	title := c.PostForm("title") //标题
	if title == "" {
		//TODO
	}

	track := c.PostForm("track") //投稿赛道
	if track == "" {
		//TODO
	}

	file, err := c.FormFile("word") //文件
	if err != nil && file == nil {
		//TODO
	}

	//存储作者信息

	//存储文章信息

	//存储文件

	//TODO
}
