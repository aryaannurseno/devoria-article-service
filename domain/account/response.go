package account

import "github.com/sangianpatrick/devoria-article-service/domain/account/entity"

type AccountAuthenticationResponse struct {
	Token   string  `json:"token"`
	Profile entity.Account `json:"profile"`
}
