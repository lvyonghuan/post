package model

// Article 文章
type Article struct {
	ID uint `json:"id"`

	Title    string `json:"title"`                  // 标题
	Author   string `json:"author"`                 // 作者名字
	AuthorID uint   `json:"author_id" gorm:"index"` // 作者ID
	Track    string `json:"track"`                  // 投稿赛道

	CreatedAt int64 `json:"created_at"`
}

// Author 作者信息
type Author struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	AuthorName     string `json:"author_name"`                         // 作者名字
	AuthorPhoneNum string `json:"author_phone_num" gorm:"uniqueIndex"` // 作者手机号 唯一
	AuthorEmail    string `json:"author_email" gorm:"uniqueIndex"`     // 作者邮箱 唯一

	PostNum uint `json:"post_num"` // 投稿数
	//TODO:议定字段
	TracksPostNum []uint `json:"tracks_post_num"` // 各个投稿赛道数
	ArticleIDs    []uint `json:"article_ids"`     // 投稿ID
}
