package article

// CreateArticleRequest is model for creating article.
type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

// EditArticleRequest is model for modified article.
type EditArticleRequest struct {
	ID       int64  `json:"id" validate:"required"`
	Title    string `json:"title" validate:"required"`
	Subtitle string `json:"subtitle" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

type EditStatusArticleRequest struct {
	ID     int64         `json:"id" validate:"required"`
	Status ArticleStatus `json:"status" validate:"required"`
}

type GetOneArticleRequest struct {
	ID     int64         `json:"id" validate:"required"`
}