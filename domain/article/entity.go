package article

import (
	"time"

	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
)

// ArticleStatus is a type of article current status.
type ArticleStatus string

const (
	ArticleStatusDraft     ArticleStatus = "DRAFT"
	ArticleStatusPublished ArticleStatus = "PUBLISHED"
	ArticleStatusArchived  ArticleStatus = "ARCHIVED"
)

// Article is a collection of property of article.
type Article struct {
	ID             int64          `json:"id"`
	Title          string         `json:"title"`
	Subtitle       string         `json:"subtitle"`
	Content        string         `json:"content"`
	Status         ArticleStatus  `json:"status"`
	CreatedAt      time.Time      `json:"createdAt"`
	PublishedAt    *time.Time     `json:"publishedAt"`
	LastModifiedAt *time.Time     `json:"lastModifiedAt"`
	Author         entity.Account `json:"author"`
}
