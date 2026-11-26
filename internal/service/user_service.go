package service

import (
	"advanced-blog-management-system/internal/model"
	"advanced-blog-management-system/internal/model/repository"
	"advanced-blog-management-system/pkg/auth"
	"context"
)

type UserService struct {
	userRepo   repository.UserRepository
	jwtManager *auth.JWTManager
}

// Сервис координирует правила регистрации и входа; SQL остается в репозитории,
// а преобразование ошибок в HTTP-коды — в handler.
func NewUserService(userRepo repository.UserRepository, jwtManager *auth.JWTManager) *UserService {
	return &UserService{
		userRepo:   userRepo,
		jwtManager: jwtManager,
	}
}

// TODO: Реализовать Register() в internal/service/user_service.go.
// Порядок: проверить req и его Validate(); проверить занятость email/username;
// захешировать пароль через auth.HashPassword; создать model.User через
// userRepo.Create; выпустить токен через jwtManager и вернуть TokenResponse,
// где user преобразован в UserResponse. При любой ошибке остановить сценарий.
// Ошибка уникального ограничения БД также должна стать ErrUserAlreadyExists:
// предварительная Exists-проверка не защищает от параллельных регистраций.
func (s *UserService) Register(ctx context.Context, req *model.UserCreateRequest) (*model.TokenResponse, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать Login() в internal/service/user_service.go.
// Провалидируйте req, найдите пользователя по email и сравните пароль с
// сохраненным хешем через auth.CheckPassword; при успехе выпустите JWT и
// верните TokenResponse без пароля. Для неизвестного email и неверного пароля
// возвращайте одну ErrInvalidCredentials, чтобы не раскрывать наличие аккаунта.
func (s *UserService) Login(ctx context.Context, req *model.UserLoginRequest) (*model.TokenResponse, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByID(): получить запись из репозитория и преобразовать
// отсутствующий результат в ErrUserNotFound; прочие ошибки БД обернуть и вернуть.
func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetByEmail() с тем же контрактом ошибок, что и GetByID().
// Не возвращайте password-хеш в HTTP-ответ; handler должен использовать DTO.
func (s *UserService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	// TODO: реализовать
	return nil, nil
}
