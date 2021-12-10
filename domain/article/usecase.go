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
	Edit(ctx context.Context, params EditArticleRequest) (resp response.Response)
	GetAllPublic(ctx context.Context) (resp response.Response)
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

	return response.Success(response.StatusCreated, newArticle)
}

func (u *articleUsecaseImpl) Edit(ctx context.Context, params EditArticleRequest) (resp response.Response) {
	// Get detail author/account
	email := ctx.Value(entity.EmailCtx).(string)
	account, err := u.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusInvalidPayload, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	lastModifiedAt := time.Now().In(u.location)

	newArticle := Article{}
	newArticle.Title = params.Title
	newArticle.Subtitle = params.Subtitle
	newArticle.Content = params.Content
	newArticle.LastModifiedAt = &lastModifiedAt
	err = u.repository.Update(ctx, params.ID, account.ID, newArticle)
	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusForbiddend, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}

	return response.Success(response.StatusOK, params.ID)
}

func (u *articleUsecaseImpl) GetAllPublic(ctx context.Context) (resp response.Response) {
	articles, err := u.repository.FindMany(ctx)
	if err != nil {
		if err == exception.ErrNotFound {
			return response.Error(response.StatusNotFound, nil, exception.ErrBadRequest)
		}
		return response.Error(response.StatusUnexpectedError, nil, exception.ErrInternalServer)
	}
	var arr []GetAll

	for _, element := range articles {
		m := GetAll{}
		m.ID = element.ID
		m.Title = element.Title
		m.Content = element.Content
		m.Status = element.Status
		m.CreatedAt = element.CreatedAt
		m.LastModifiedAt = element.LastModifiedAt
		m.AuthorID = element.Author.ID

		arr = append(arr, m)
	}

	return response.Success(response.StatusOK, arr)
}
