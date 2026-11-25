package repository

import (
	"advanced-blog-management-system/internal/model"
	"context"
	"database/sql"
)

// Методы репозитория отвечают за SQL и преобразование строк в модели, но не
// формируют HTTP-ответы. Оборачивайте ошибки через fmt.Errorf с %w, чтобы
// вызывающий сервис мог определить причину через errors.Is/errors.As.
type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// TODO: Реализовать Create() через INSERT ... RETURNING id.
// Передайте значения параметрами PostgreSQL, установите CreatedAt/UpdatedAt
// единым временем UTC, сохраните bcrypt-хеш в password и присвойте полученный
// id переданной модели. Уникальные конфликты должны быть распознаваемы сервисом.
func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать GetByID явным списком столбцов, включая password для
// внутренней проверки входа. На sql.ErrNoRows верните nil, nil; прочие ошибки
// оберните через %w. Handler не должен сериализовать эту внутреннюю модель.
func (r *UserRepo) GetByID(ctx context.Context, id int) (*model.User, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByEmail параметризованным запросом с тем же scan order,
// что и GetByID; различайте отсутствие записи и ошибку БД.
func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByUsername параметризованным запросом; не конкатенируйте
// значение username в SQL и соблюдайте nil,nil-контракт для отсутствия строки.
func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать ExistsByEmail через SELECT EXISTS и QueryRowContext;
// возвращайте ошибку выполнения отдельно от значения bool.
func (r *UserRepo) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	// TODO: реализовать
	return false, nil
}

// TODO: Реализовать ExistsByUsername аналогично ExistsByEmail.
func (r *UserRepo) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	// TODO: реализовать
	return false, nil
}

// TODO: Реализовать Update через ExecContext, обновив UpdatedAt и используя
// параметры. Проверьте RowsAffected: ноль означает, что пользователя нет;
// конфликт уникальности email/username должен остаться различимым.
func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать Delete параметризованным ExecContext и проверить
// RowsAffected; верните соответствующую ошибку, если пользователь не найден.
func (r *UserRepo) Delete(ctx context.Context, id int) error {
	// TODO: реализовать
	return nil
}
