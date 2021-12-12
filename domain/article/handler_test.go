package article_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/sangianpatrick/devoria-article-service/domain/article"
	"github.com/sangianpatrick/devoria-article-service/domain/article/mocks"
	"github.com/sangianpatrick/devoria-article-service/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type responseBody struct {
	Data   interface{} `json:"data"`
	Status string      `json:"status"`
}

func TestHandlerCreate_Success(t *testing.T) {
	newCreateArticleRequest := article.CreateArticleRequest{
		Title:    "test",
		Subtitle: "test",
		Content:  "test",
	}
	articleGetArticleResponse := article.GetArticleResponse{
		ID:        1,
		Title:     newCreateArticleRequest.Title,
		Subtitle:  newCreateArticleRequest.Subtitle,
		Content:   newCreateArticleRequest.Content,
		Status:    article.ArticleStatusDraft,
		CreatedAt: time.Now(),
		AuthorID:  1,
	}

	resp := response.Success(response.StatusCreated, articleGetArticleResponse)

	validate := validator.New()

	articleUsecase := new(mocks.ArticleUsecase)
	articleUsecase.On("Create", mock.Anything, mock.AnythingOfType("article.CreateArticleRequest")).Return(resp)

	newCreateArticleRequestBuff, _ := json.Marshal(newCreateArticleRequest)

	articleHTTPHandler := article.ArticleHTTPHandler{
		Validate: validate,
		Usecase:  articleUsecase,
	}

	r := httptest.NewRequest(http.MethodPost, "/just/for/testing", bytes.NewReader(newCreateArticleRequestBuff))
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(articleHTTPHandler.Create)
	handler.ServeHTTP(recorder, r)

	rb := responseBody{}
	if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.StatusCreated, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusCreated))
	assert.NotNil(t, rb.Data, "should not be nil")

	data, ok := rb.Data.(map[string]interface{})
	if !ok {
		t.Error("response data isn't a type of 'map[string]interface{}'")
		return
	}

	assert.Equal(t, articleGetArticleResponse.Title, data["title"], fmt.Sprintf("title should be '%s'", articleGetArticleResponse.Title))

	articleUsecase.AssertExpectations(t)
}

func TestHandlerEdit_Success(t *testing.T) {
	newEditArticleRequest := article.EditArticleRequest{
		ID:       1,
		Title:    "test",
		Subtitle: "test",
		Content:  "test",
	}
	articleEditArticleResponse := newEditArticleRequest

	resp := response.Success(response.StatusOK, articleEditArticleResponse)

	validate := validator.New()

	articleUsecase := new(mocks.ArticleUsecase)
	articleUsecase.On("Edit", mock.Anything, mock.AnythingOfType("article.EditArticleRequest")).Return(resp)

	newEditArticleRequestBuff, _ := json.Marshal(newEditArticleRequest)

	articleHTTPHandler := article.ArticleHTTPHandler{
		Validate: validate,
		Usecase:  articleUsecase,
	}

	r := httptest.NewRequest(http.MethodPut, "/just/for/testing", bytes.NewReader(newEditArticleRequestBuff))
	recorder := httptest.NewRecorder()

	path := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, path)

	handler := http.HandlerFunc(articleHTTPHandler.Edit)
	handler.ServeHTTP(recorder, r)

	rb := responseBody{}
	if err := json.NewDecoder(recorder.Body).Decode(&rb); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, response.StatusOK, rb.Status, fmt.Sprintf("should be status '%s'", response.StatusOK))
	assert.NotNil(t, rb.Data, "should not be nil")

	data, ok := rb.Data.(map[string]interface{})
	if !ok {
		t.Error("response data isn't a type of 'map[string]interface{}'")
		return
	}

	assert.Equal(t, articleEditArticleResponse.Title, data["title"], fmt.Sprintf("title should be '%s'", articleEditArticleResponse.Title))

	articleUsecase.AssertExpectations(t)
}
