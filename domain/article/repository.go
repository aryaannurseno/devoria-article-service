package article

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/sangianpatrick/devoria-article-service/exception"
)

type ArticleRepository interface {
	Save(ctx context.Context, article Article) (ID int64, err error)
	Update(ctx context.Context, ID int64, authorId int64, updatedArticle Article) (err error)
	FindByID(ctx context.Context, ID int64) (article Article, err error)
	FindMany(ctx context.Context) (bunchOfArticles []Article, err error)
	FindManySpecificProfile(ctx context.Context, authorId int64) (bunchOfArticles []Article, err error)
	UpdateStatus(ctx context.Context, ID int64, authorId int64, updatedArticle Article) (err error)
}

type articleRepositoryImpl struct {
	db        *sql.DB
	tableName string
}

func NewArticleRepository(db *sql.DB, tableName string) ArticleRepository {
	return &articleRepositoryImpl{
		db:        db,
		tableName: tableName,
	}
}

func (r *articleRepositoryImpl) Save(ctx context.Context, article Article) (ID int64, err error) {
	command := fmt.Sprintf("INSERT INTO %s (title, subtitle, content, status, createdAt, authorId) VALUES (?, ?, ?, ?, ?, ?)", r.tableName)
	stmt, err := r.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		article.Title,
		article.Subtitle,
		article.Content,
		article.Status,
		article.CreatedAt,
		article.Author.ID,
	)

	if err != nil {
		log.Println(err)
		return
	}

	ID, _ = result.LastInsertId()

	return
}

func (r *articleRepositoryImpl) Update(ctx context.Context, ID int64, authorId int64, updatedArticle Article) (err error) {
	command := fmt.Sprintf(`UPDATE %s SET title = ?, subtitle = ?, content = ?, lastModifiedAt = ? WHERE id = ? AND authorId = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		updatedArticle.Title,
		updatedArticle.Subtitle,
		updatedArticle.Content,
		*updatedArticle.LastModifiedAt,
		ID,
		authorId,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		err = exception.ErrNotFound
		return
	}

	return
}
func (r *articleRepositoryImpl) FindByID(ctx context.Context, ID int64) (article Article, err error) {
	query := fmt.Sprintf(`SELECT id, title, subtitle, content, status, createdAt, publishedAt, lastModifiedAt, authorId FROM %s WHERE id = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, ID)

	var publishedAt sql.NullTime
	var lastModifiedAt sql.NullTime

	err = row.Scan(
		&article.ID,
		&article.Title,
		&article.Subtitle,
		&article.Content,
		&article.Status,
		&article.CreatedAt,
		&publishedAt,
		&lastModifiedAt,
		&article.Author.ID,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrNotFound
		return
	}

	if publishedAt.Valid {
		article.PublishedAt = &publishedAt.Time
	}

	if lastModifiedAt.Valid {
		article.LastModifiedAt = &lastModifiedAt.Time
	}

	return
}
func (r *articleRepositoryImpl) FindMany(ctx context.Context) (bunchOfArticles []Article, err error) {
	query := fmt.Sprintf(`SELECT id, title, subtitle, content, status, createdAt, publishedAt, lastModifiedAt, authorId FROM %s`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		fmt.Println(err.Error())
		err = exception.ErrInternalServer
		return
	}

	defer rows.Close()

	for rows.Next() {
		article := Article{}
		var publishedAt sql.NullTime
		var lastModifiedAt sql.NullTime

		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Subtitle,
			&article.Content,
			&article.Status,
			&article.CreatedAt,
			&publishedAt,
			&lastModifiedAt,
			&article.Author.ID,
		)

		if err != nil {
			log.Println(err)
			err = exception.ErrNotFound
			return
		}

		if publishedAt.Valid {
			article.PublishedAt = &publishedAt.Time
		}

		if lastModifiedAt.Valid {
			article.LastModifiedAt = &lastModifiedAt.Time
		}

		bunchOfArticles = append(bunchOfArticles, article)
	}

	return
}
func (r *articleRepositoryImpl) FindManySpecificProfile(ctx context.Context, authorId int64) (bunchOfArticles []Article, err error) {
	query := fmt.Sprintf(`SELECT id, title, subtitle, content, createdAt, publishedAt, lastModifiedAt, authorId FROM %s WHERE authorId = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, authorId)
	if err != nil {
		fmt.Println(err.Error())
		err = exception.ErrInternalServer
		return
	}

	defer rows.Close()

	for rows.Next() {
		article := Article{}
		var publishedAt sql.NullTime
		var lastModifiedAt sql.NullTime

		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Subtitle,
			&article.Content,
			&article.CreatedAt,
			&publishedAt,
			&lastModifiedAt,
			&article.Author.ID,
		)

		if err != nil {
			log.Println(err)
			err = exception.ErrNotFound
			return
		}

		if publishedAt.Valid {
			article.PublishedAt = &publishedAt.Time
		}

		if lastModifiedAt.Valid {
			article.LastModifiedAt = &lastModifiedAt.Time
		}

		bunchOfArticles = append(bunchOfArticles, article)
	}

	return
}
func (r *articleRepositoryImpl) UpdateStatus(ctx context.Context, ID int64, authorId int64, updatedArticle Article) (err error) {
	command := fmt.Sprintf(`UPDATE %s SET status = ?, publishedAt = ? WHERE id = ? AND authorId = ?`, r.tableName)
	stmt, err := r.db.PrepareContext(ctx, command)
	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(
		ctx,
		updatedArticle.Status,
		updatedArticle.PublishedAt,
		ID,
		authorId,
	)

	if err != nil {
		log.Println(err)
		err = exception.ErrInternalServer
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected < 1 {
		err = exception.ErrNotFound
		return
	}

	return
}
