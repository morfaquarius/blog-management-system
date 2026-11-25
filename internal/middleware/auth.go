package middleware

import (
	"advanced-blog-management-system/pkg/auth"
	"context"
	"encoding/json"
	"net/http"
)

// contextKey - кастомный тип, чтобы избежать коллизии
type contextKey string

const (
	// Эти ключи используются только после проверки JWT, чтобы handlers/services
	// получали идентичность пользователя из доверенного контекста запроса.
	UserIDKey    contextKey = "userID"
	UserEmailKey contextKey = "userEmail"
	UserNameKey  contextKey = "username"
)

// AuthMiddleware хранит общий JWTManager, созданный один раз при старте
// приложения и используемый для проверки токенов на каждом запросе.
type AuthMiddleware struct {
	jwtManager *auth.JWTManager
}

// TODO: Реализовать NewAuthMiddleware в internal/middleware/auth.go.
// Верните middleware, связанный с переданным JWTManager; не создавайте новый
// менеджер и не читайте секрет из окружения внутри HTTP-запроса.
func NewAuthMiddleware(jwtManager *auth.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
	}
}

// TODO: Реализовать RequireAuth для защищенных маршрутов.
// Принимайте только заголовок Authorization вида "Bearer <token>", проверяйте
// его через jwtManager.ValidateToken, помещайте user_id/email/username из
// проверенных claims в context и вызывайте next один раз. При отсутствующем,
// некорректном или просроченном токене верните JSON 401 и не вызывайте next.
func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	// При отказе обязательно return до вызова next. Context создавайте через
	// r.WithContext, а claims используйте только после успешной проверки подписи.
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next(w, r)
	}
}

// TODO: Реализовать OptionalAuth для публичных маршрутов, которым нужен
// необязательный пользовательский контекст. При валидном токене добавьте claims;
// при его отсутствии или ошибке валидации вызовите next без пользовательских
// значений и не превращайте публичный запрос в 401.
func (m *AuthMiddleware) OptionalAuth(next http.HandlerFunc) http.HandlerFunc {
	// Используйте только на публичных маршрутах, которым действительно нужен
	// необязательный контекст пользователя.
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next(w, r)
	}
}

// TODO: Реализовать GetUserIDFromContext: проверить наличие значения и тип int;
// возвращайте ok=false, если middleware не добавил ID или тип не совпадает.
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	// TODO: реализовать
	return 0, false
}

// TODO: Реализовать GetUserEmailFromContext с проверкой наличия значения и типа.
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	// TODO: реализовать
	return "", false
}

// TODO: Реализовать GetUsernameFromContext с проверкой наличия значения и типа.
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	// TODO: реализовать
	return "", false
}

// TODO: Реализовать extractToken: прочитайте Authorization и верните только
// непустую вторую часть корректного bearer-заголовка (без лишних частей).
// Возвращайте пустую строку для отсутствующего/некорректного заголовка; никогда
// не логируйте содержимое токена.
func extractToken(r *http.Request) string {
	// Принимайте только Bearer-токен без пустого значения и лишних частей.
	// TODO: реализовать
	return ""
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, ErrorResponse{Error: message})
}
