package article_test

import (
	"context"
	"database/sql/driver"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sangianpatrick/devoria-article-service/domain/account/entity"
	"github.com/sangianpatrick/devoria-article-service/domain/article"
	"github.com/stretchr/testify/assert"
)

var (
	tableName string = "article"
)

func TestRepositorySave_Success(t *testing.T) {
	db, mock, _ := sqlmock.New()

	ctx := context.TODO()

	newArticle := article.Article{
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

	expectedCommand := fmt.Sprintf("INSERT INTO %s", tableName)
	expectedArgs := make([]driver.Value, 0)
	expectedArgs = append(
		expectedArgs,
		newArticle.Title,
		newArticle.Subtitle,
		newArticle.Content,
		newArticle.Status,
		newArticle.CreatedAt,
		newArticle.Author.ID,
	)

	mock.ExpectPrepare(expectedCommand).
		ExpectExec().
		WithArgs(expectedArgs...).
		WillReturnResult(sqlmock.NewResult(1, 1))

	articleRepostitory := article.NewArticleRepository(db, tableName)
	ID, err := articleRepostitory.Save(ctx, newArticle)

	assert.NoError(t, err, "should not be error")
	assert.Equal(t, int64(1), ID, "id should be `1`")

	if err = mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

