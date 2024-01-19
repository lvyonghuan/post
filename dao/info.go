package dao

import (
	"errors"
	"post/model"

	"gorm.io/gorm"
)

// AuthorInfoGet 通过作者邮箱、作者电话验证作者信息是否存在并且返回作者信息
func AuthorInfoGet(author *model.Author) (err error) {
	result := db.Where("author_email = ? AND author_phone_num = ?", author.AuthorEmail, author.AuthorPhoneNum).Find(author)
	if result.Error != nil {
		//如果没有作者信息，创建作者信息
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//创建作者信息
			err = db.Create(author).Error
		} else {
			return errors.New("查找作者信息失败:" + result.Error.Error())
		}
	}

	return err
}

// UpdateAuthorInfo 更新作者信息
func UpdateAuthorInfo(author *model.Author) (err error) {
	result := db.Where("author_email = ? AND author_phone_num = ?", author.AuthorEmail, author.AuthorPhoneNum).Updates(author)
	if result.Error != nil {
		return errors.New("更新作者信息失败:" + result.Error.Error())
	}
	return err
}

// ArticleInfoStore 存储文章信息
func ArticleInfoStore(article *model.Article) (err error) {
	result := db.Create(article)
	if result.Error != nil {
		return errors.New("存储文章信息失败:" + result.Error.Error())
	}
	return err
}

// GetArticleInfoByTitle 通过文章标题获取文章信息
func GetArticleInfoByTitle(article *model.Article) (err error) {
	result := db.Where("title = ?", article.Title).Find(article)
	if result.Error != nil {
		return errors.New("查找文章信息失败:" + result.Error.Error())
	}
	return err
}

// GetArticleInfoByID 通过文章ID获取文章信息
func GetArticleInfoByID(article *model.Article) (err error) {
	result := db.Where("id = ?", article.ID).Find(article)
	if result.Error != nil {
		return errors.New("查找文章信息失败:" + result.Error.Error())
	}
	return err
}
