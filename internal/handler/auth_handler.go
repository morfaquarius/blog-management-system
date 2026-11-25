package handler

import (
	"advanced-blog-management-system/internal/service"
	"net/http"
)

type AuthHandler struct {
	userService service.UserServiceInterface
}

// NewAuthHandler создает handler с интерфейсом сервиса; SQL и правила
// регистрации/входа должны оставаться за пределами HTTP-слоя.
func NewAuthHandler(userService service.UserServiceInterface) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Register обрабатывает POST /register: декодирует UserCreateRequest, вызывает
// сервис и возвращает TokenResponse со статусом 201. При неверном методе,
// JSON или DTO верните ошибку и сразу завершите handler; не передавайте сервису
// частично декодированный запрос. В ответе поле user — только UserResponse,
// никогда model.User с Password.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// Login обрабатывает POST /login: декодирует и валидирует UserLoginRequest,
// вызывает сервис и при успехе возвращает TokenResponse со статусом 200.
// Неверный email и неверный пароль должны иметь одинаковый внешний ответ.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetProfile — дополнительный GET endpoint профиля: если он будет подключен,
// userID следует взять из контекста проверенного middleware, а не из URL/query.
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать получение профиля (опционально)
}

// respondWithJSON устанавливает application/json, записывает статус ровно один
// раз и кодирует тело. Если кодирование после отправки заголовков завершится
// ошибкой, не пытайтесь отправить поверх него второй HTTP-ответ.
func (h *AuthHandler) respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// TODO: реализовать
}

// respondWithError формирует общий ErrorResponse и передает его в
// respondWithJSON; доменные ошибки предварительно преобразуются в utils.go.
func (h *AuthHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}
