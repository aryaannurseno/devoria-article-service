package article

import (
	"time"

	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
)

type CreateArticleResponse struct {
	Article `json:"token"`
	Profile entity.Account `json:"profile"`
}

type GetAll struct {
	ID             int64         `json:"id"`
	Title          string        `json:"title"`
	Subtitle       string        `json:"subtitle"`
	Content        string        `json:"content"`
	Status         ArticleStatus `json:"status"`
	CreatedAt      time.Time     `json:"createdAt"`
	PublishedAt    *time.Time    `json:"publishedAt"`
	LastModifiedAt *time.Time    `json:"lastModifiedAt"`
	AuthorID       int64         `json:"authorId"`
}
