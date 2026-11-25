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

// TODO: Добавить структуру JWTManager с полями secretKey []byte и ttl time.Duration
// Реализовать конструктор NewJWTManager(secretKey string, ttlHours int) (*JWTManager, error)
type JWTManager struct {
	secretKey []byte
	ttl       time.Duration
}

// TODO: Реализовать конструктор NewJWTManager.
// Вернуть ошибку при пустом секрете или неположительном сроке действия.
func NewJWTManager(secretKey string, ttlHours int) (*JWTManager, error) {
	// Проверьте непустой секрет и положительный TTL. Секрет должен поступать из
	// конфигурации, не храниться в исходном коде и не попадать в логи.
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать метод GenerateToken(userID int, email, username string) (string, time.Time, error)
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

// TODO: Реализовать метод ValidateToken(tokenString string) (*Claims, error)
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
