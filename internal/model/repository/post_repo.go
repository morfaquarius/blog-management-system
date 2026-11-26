package repository

import (
	"advanced-blog-management-system/internal/model"
	"context"
	"database/sql"
)

// Для списков используйте стабильную сортировку (created_at и при необходимости
// id), иначе записи с одинаковым временем могут перемещаться между страницами.
type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

// TODO: Реализовать Create через INSERT ... RETURNING id с параметрами.
// Установите CreatedAt в UTC и заполните ID из Scan; AuthorID задается сервисом
// по проверенному JWT-контексту, а не клиентом.
func (r *PostRepo) Create(ctx context.Context, post *model.Post) error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать GetByID с явным списком столбцов и параметром $1.
// При sql.ErrNoRows верните nil,nil; прочую ошибку БД оберните через %w.
func (r *PostRepo) GetByID(ctx context.Context, id int) (*model.Post, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetAll параметрами limit/offset и стабильной сортировкой
// (created_at DESC, затем id DESC). Закройте rows и проверьте rows.Err().
func (r *PostRepo) GetAll(ctx context.Context, limit, offset int) ([]*model.Post, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetTotalCount через QueryRowContext и вернуть ошибку
// сканирования/БД с контекстом.
func (r *PostRepo) GetTotalCount(ctx context.Context) (int, error) {
	// TODO: реализовать
	return 0, nil
}

// TODO: Реализовать Exists через SELECT EXISTS с параметром ID.
func (r *PostRepo) Exists(ctx context.Context, id int) (bool, error) {
	// TODO: реализовать
	return false, nil
}

// TODO: Реализовать GetByAuthorID с параметрами authorID/limit/offset,
// стабильной сортировкой и проверкой ошибок итерации rows.
func (r *PostRepo) GetByAuthorID(ctx context.Context, authorID int, limit, offset int) ([]*model.Post, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetTotalCountByAuthorID через параметризованный COUNT.
func (r *PostRepo) GetTotalCountByAuthorID(ctx context.Context, authorID int) (int, error) {
	// TODO: реализовать
	return 0, nil
}
