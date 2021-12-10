package article

import "github.com/sangianpatrick/devoria-article-service/domain/account/entity"

type CreateArticleResponse struct {
	Article  `json:"token"`
	Profile entity.Account `json:"profile"`
}
