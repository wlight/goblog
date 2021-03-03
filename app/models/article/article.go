package article

import (
	"goblog/app/models"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

// Article 文章模型
type Article struct {
	models.BaseModel
	Title string `valid:"title"`
	Body  string `valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User user.User
}

// Link 方法用来生成文章链接
func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(a.ID))
}

// 创建日期
func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}