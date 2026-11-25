package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

// Claims представляет данные, хранимые в JWT токене
type Claims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTManager содержит ключ подписи и срок жизни токенов. Создавайте один
// экземпляр из конфигурации при старте и передавайте его сервису пользователей
// и AuthMiddleware; не создавайте новый ключ на каждый запрос.
type JWTManager struct {
	secretKey []byte
	ttl       time.Duration
}

// TODO: Реализовать NewJWTManager в pkg/auth/jwt.go.
// Отклоняйте пустой/пробельный secretKey и ttlHours <= 0; срок в часах
// преобразуйте в time.Duration. Ошибка должна сообщать о неверной конфигурации,
// не включая сам ключ.
func NewJWTManager(secretKey string, ttlHours int) (*JWTManager, error) {
	// Проверьте непустой секрет и положительный TTL. Секрет должен поступать из
	// конфигурации, не храниться в исходном коде и не попадать в логи.
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GenerateToken для вызова из UserService.Register и Login.
// Принимает идентификатор и публичные данные пользователя; не добавляйте пароль,
// его хеш или другие секреты в claims.
// - Создать Claims со сроком действия = now + m.ttl
// - Установить IssuedAt, NotBefore, ExpiresAt, Issuer, Subject в Claims
// - Подписать токен методом HS256 секретным ключом
// - Вернуть (tokenString, expiresAt, nil) или ("", time.Time{}, error)
func (m *JWTManager) GenerateToken(userID int, email, username string) (string, time.Time, error) {
	// Время задавайте в UTC. Не помещайте в claims пароль, хеш пароля или другие
	// секреты: содержимое JWT закодировано, но не зашифровано.
	// TODO: реализовать
	return "", time.Time{}, nil
}

// TODO: Реализовать ValidateToken для вызова AuthMiddleware.RequireAuth.
// Метод возвращает claims только после успешной проверки подписи, алгоритма и
// срока действия; ошибки сопоставляются с ErrInvalidToken/ErrExpiredToken.
// - Распарсить токен с проверкой подписи через jwt.ParseWithClaims
// - Извлечь Claims из токена
// - Проверить что token.Valid == true и claims != nil
// - Проверить что время истечения еще не наступило (ExpiresAt.Time > now)
// - Вернуть claims или ошибку (ErrInvalidToken, ErrExpiredToken)
func (m *JWTManager) ValidateToken(tokenString string) (*Claims, error) {
	// Явно ограничьте допустимый алгоритм подписи HS256: нельзя выбирать алгоритм
	// проверки на основании непроверенного заголовка токена. Claims используйте
	// только после успешной проверки подписи и срока действия.
	// TODO: реализовать
	return nil, nil
}
