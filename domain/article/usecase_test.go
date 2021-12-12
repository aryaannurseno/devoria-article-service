package article_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	cryptoMocks "github.com/sangianpatrick/devoria-article-service/crypto/mocks"
	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
	accountMocks "github.com/sangianpatrick/devoria-article-service/domain/account/mocks"
	"github.com/sangianpatrick/devoria-article-service/domain/article"
	articleMocks "github.com/sangianpatrick/devoria-article-service/domain/article/mocks"
	jsonWebTokenMocks "github.com/sangianpatrick/devoria-article-service/jwt/mocks"
	sessionMocks "github.com/sangianpatrick/devoria-article-service/session/mocks"
)

var (
	location *time.Location
)

func TestMain(m *testing.M) {
	location, _ = time.LoadLocation("Asia/Jakarta")

	m.Run()
}

func TestUsecaseCreate_Success(t *testing.T) {
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	accountRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(entity.Account{}, nil)
	articleRepo := new(articleMocks.ArticleRepository)

	articleRepo.On("Save", mock.Anything, mock.AnythingOfType("article.Article")).Return(int64(1), nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	params := article.CreateArticleRequest{
		Title:    "test",
		Subtitle: "test",
		Content:  "test",
	}
	resp := u.Create(ctx, params)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseEdit_Success(t *testing.T) {
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	accountRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(entity.Account{}, nil)

	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("Update",
		mock.Anything,
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("article.Article")).Return(nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	params := article.EditArticleRequest{
		ID:       1,
		Title:    "test",
		Subtitle: "test",
		Content:  "test",
	}
	resp := u.Edit(ctx, params)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseGetAllPublic_Success(t *testing.T) {
	var articles = []article.Article{
		{
			ID:        1,
			Title:     "test",
			Subtitle:  "test",
			Content:   "test",
			Status:    article.ArticleStatusDraft,
			CreatedAt: time.Now().In(location),
			Author: entity.Account{
				ID: 1,
			},
		},
	}
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("FindMany",
		mock.Anything).Return(articles, nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	resp := u.GetAllPublic(ctx)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseGetAllPrivate_Success(t *testing.T) {
	var articles = []article.Article{
		{
			ID:        1,
			Title:     "test",
			Subtitle:  "test",
			Content:   "test",
			Status:    article.ArticleStatusDraft,
			CreatedAt: time.Now().In(location),
			Author: entity.Account{
				ID: 1,
			},
		},
	}
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	accountRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(entity.Account{}, nil)

	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("FindManySpecificProfile",
		mock.Anything, mock.AnythingOfType("int64")).Return(articles, nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	resp := u.GetAllPrivate(ctx)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseEditStatus_SuccessPublished(t *testing.T) {
	var dataArticle = article.Article{
		ID:        1,
		Title:     "test",
		Subtitle:  "test",
		Content:   "test",
		Status:    article.ArticleStatusDraft,
		CreatedAt: time.Now().In(location),
		Author: entity.Account{
			ID: 1,
		},
	}

	var dataAcc = entity.Account{
		ID:        1,
		Email:     "mail@mail.c",
		FirstName: "Pablo",
		LastName:  "Picasso",
		CreatedAt: time.Now().In(location),
	}
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	accountRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(dataAcc, nil)

	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("FindByID",
		mock.Anything, mock.AnythingOfType("int64")).Return(dataArticle, nil)

	articleRepo.On("UpdateStatus",
		mock.Anything,
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("article.Article")).Return(nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	params := article.EditStatusArticleRequest{
		ID:     1,
		Status: "PUBLISHED",
	}

	resp := u.EditStatus(ctx, params)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseEditStatus_SuccessArchived(t *testing.T) {
	var dataArticle = article.Article{
		ID:        1,
		Title:     "test",
		Subtitle:  "test",
		Content:   "test",
		Status:    article.ArticleStatusPublished,
		CreatedAt: time.Now().In(location),
		Author: entity.Account{
			ID: 1,
		},
	}

	var dataAcc = entity.Account{
		ID:        1,
		Email:     "mail@mail.c",
		FirstName: "Pablo",
		LastName:  "Picasso",
		CreatedAt: time.Now().In(location),
	}
	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	accountRepo.On("FindByEmail", mock.Anything, mock.AnythingOfType("string")).Return(dataAcc, nil)

	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("FindByID",
		mock.Anything, mock.AnythingOfType("int64")).Return(dataArticle, nil)

	articleRepo.On("UpdateStatus",
		mock.Anything,
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("int64"),
		mock.AnythingOfType("article.Article")).Return(nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	params := article.EditStatusArticleRequest{
		ID:     1,
		Status: "ARCHIVED",
	}

	resp := u.EditStatus(ctx, params)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}

func TestUsecaseGetOne_Success(t *testing.T) {

	sess := new(sessionMocks.Session)
	jsonWebToken := new(jsonWebTokenMocks.JSONWebToken)
	crypto := new(cryptoMocks.Crypto)
	accountRepo := new(accountMocks.AccountRepository)
	articleRepo := new(articleMocks.ArticleRepository)
	articleRepo.On("FindByID",
		mock.Anything, mock.AnythingOfType("int64")).Return(article.Article{}, nil)

	u := article.NewArticleUsecase("globalIVTest", sess, jsonWebToken, crypto, location, articleRepo, accountRepo)
	ctx := context.WithValue(context.Background(), entity.EmailCtx, "email@gmail.co")

	params := article.GetOneArticleRequest{
		ID: 1,
	}

	resp := u.GetOne(ctx, params)
	assert.NoError(t, resp.Err())

	sess.AssertExpectations(t)
	jsonWebToken.AssertExpectations(t)
	crypto.AssertExpectations(t)
	accountRepo.AssertExpectations(t)
	articleRepo.AssertExpectations(t)

}
