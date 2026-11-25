package model

import (
	"time"
)

// User представляет модель пользователя в системе
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Post представляет модель поста в блоге
type Post struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Comment представляет запись комментария в PostgreSQL.
// Внешнее API называет поле Content как "text".
type Comment struct {
	ID        int       `json:"id" db:"id"`
	Content   string    `json:"text" db:"content"`
	PostID    int       `json:"post_id" db:"post_id"`
	AuthorID  int       `json:"author_id" db:"author_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// UserCreateRequest представляет запрос на создание пользователя
type UserCreateRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// UserLoginRequest представляет запрос на вход пользователя
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// PostCreateRequest представляет запрос на создание поста
type PostCreateRequest struct {
	Title   string `json:"title" validate:"required,min=1,max=200"`
	Content string `json:"content" validate:"required,min=1"`
}

// CommentCreateRequest принимает text; ID поста берется из URL.
type CommentCreateRequest struct {
	Content string `json:"text" validate:"required,min=1,max=1000"`
}

// UserResponse - структура для ответа с данными пользователя (без пароля)
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// TokenResponse - структура для ответа с JWT токеном
type TokenResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	User      UserResponse `json:"user"`
}

// PostResponse - структура для ответа с данными поста
type PostResponse struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"author_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// CommentResponse - структура для ответа с данными комментария
type CommentResponse struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	AuthorID  int       `json:"author_id"`
	Content   string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// TODO: Реализовать ToResponse() в internal/model/models.go.
// Это единственный путь преобразования пользователя в публичный DTO:
// скопируйте только поля UserResponse и никогда не переносите Password.
func (u *User) ToResponse() UserResponse {
	// TODO: реализовать
	return UserResponse{}
}

// TODO: Реализовать Validate() для UserCreateRequest через validator.New().Struct(r).
// Учитывайте теги username (3-50 символов), email и пароль (не менее 6);
// проверку строк только из пробелов добавьте явно, если validator-теги ее не покрывают.
func (r *UserCreateRequest) Validate() error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать Validate() для UserLoginRequest через validator.New().Struct(r).
// Требуйте корректный email и непустой пароль; не применяйте к входу правила
// минимальной длины создания пароля, чтобы не менять существующие учетные данные.
func (r *UserLoginRequest) Validate() error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать Validate() для PostCreateRequest через validator.New().Struct(r).
// Заголовок должен быть непустым и длиной не более 200, content — непустым;
// не принимайте авторский ID из этого DTO.
func (r *PostCreateRequest) Validate() error {
	// TODO: реализовать
	return nil
}

// TODO: Реализовать Validate() для CommentCreateRequest через validator.New().Struct(r).
// JSON-имя поля — text; допустимая длина от 1 до 1000 символов. ID поста
// берется из URL и проверяется handler/service отдельно.
func (r *CommentCreateRequest) Validate() error {
	// TODO: реализовать
	return nil
}
