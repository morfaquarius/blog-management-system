package repository

import (
	"advanced-blog-management-system/internal/model"
	"context"
	"database/sql"
)

// После rows.Next() обязательно проверьте rows.Err(): ошибка может возникнуть
// во время чтения результата, уже после успешного выполнения SQL-запроса.
type CommentRepo struct {
	db *sql.DB
}

func NewCommentRepo(db *sql.DB) *CommentRepo {
	return &CommentRepo{db: db}
}

// TODO: Реализовать Create через INSERT ... RETURNING id с параметрами.
// Установите CreatedAt/UpdatedAt одним UTC-временем и сканируйте ID в модель;
// content хранит значение JSON-поля text.
func (r *CommentRepo) Create(ctx context.Context, comment *model.Comment) error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать GetByID с явным списком столбцов. При sql.ErrNoRows
// верните nil,nil, а остальные ошибки оберните с %w.
func (r *CommentRepo) GetByID(ctx context.Context, id int) (*model.Comment, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByPostID с параметрами postID/limit/offset и
// стабильной сортировкой created_at ASC, id ASC. Закройте rows и проверьте
// rows.Err() после обхода результата.
func (r *CommentRepo) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]*model.Comment, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetCountByPostID через QueryRowContext и параметр postID.
func (r *CommentRepo) GetCountByPostID(ctx context.Context, postID int) (int, error) {
	// TODO: реализовать
	return 0, nil
}
