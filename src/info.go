package src

import (
	"post/dao"
	"post/model"
)

// GetAuthorInfo 获取作者信息
func GetAuthorInfo(authorName, authorPhoneNum, authorEmail string) (model.Author, error) {
	author := model.Author{
		AuthorName:     authorName,
		AuthorPhoneNum: authorPhoneNum,
		AuthorEmail:    authorEmail,
	}

	//在数据库中查找作者信息
	err := dao.AuthorInfoGet(&author)
	if err != nil {
		return author, err
	}

	return author, nil
}

// StoreArticleInfo 存储文章信息
func StoreArticleInfo(author *model.Author, title, track string) error {
	article := model.Article{
		AuthorID: author.ID,
		Author:   author.AuthorName,
		Title:    title,
		Track:    track,
	}

	//存储文章信息
	err := dao.ArticleInfoStore(&article)
	if err != nil {
		return err
	}

	author.PostNum++ //更新作者总投稿数
	//TODO:更新赛道投稿数
	//更新作者信息
	err = dao.UpdateAuthorInfo(author)

	return err
}
