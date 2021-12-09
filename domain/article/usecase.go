package article

import (
	"context"
	"time"

	"github.com/sangianpatrick/devoria-article-service/crypto"
	"github.com/sangianpatrick/devoria-article-service/domain/account"
	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
	"github.com/sangianpatrick/devoria-article-service/exception"
	"github.com/sangianpatrick/devoria-article-service/jwt"
	"github.com/sangianpatrick/devoria-article-service/response"
	"github.com/sangianpatrick/devoria-article-service/session"
)

type ArticleUsecase interface {
	Create(ctx context.Context, params CreateArticleRequest) (resp response.Response)
}

type articleUsecaseImpl struct {
	globalIV     string
	session      session.Session
	jsonWebToken jwt.JSONWebToken
	crypto       crypto.Crypto
	location     *time.Location
	repository   ArticleRepository
	accountRepo  account.AccountRepository
}

func NewArticleUsecase(
	globalIV string,
	session session.Session,
	jsonWebToken jwt.JSONWebToken,
	crypto crypto.Crypto,
	location *time.Location,
	repository ArticleRepository,
	accountRepo account.AccountRepository,
) ArticleUsecase {
	return &articleUsecaseImpl{
		globalIV:     globalIV,
		session:      session,
		jsonWebToken: jsonWebToken,
		crypto:       crypto,
		location:     location,
		repository:   repository,
		accountRepo:  accountRepo,
	}
}

func (u *articleUsecaseImpl) Create(ctx context.Context, params CreateArticleRequest) (resp response.Response) {
	// Get detail author/account
	email := ctx.Value(entity.EmailCtx).(string)
	account, err := u.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusInvalidPayload, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	newArticle := Article{}
	newArticle.Title = params.Title
	newArticle.Subtitle = params.Subtitle
	newArticle.Content = params.Content
	newArticle.Status = ArticleStatusDraft
	newArticle.CreatedAt = time.Now().In(u.location)
	newArticle.Author = account
	ID, err := u.repository.Save(ctx, newArticle)
	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusInvalidPayload, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	newArticle.ID = ID

	return response.Success(response.StatusOK, newArticle)
}
