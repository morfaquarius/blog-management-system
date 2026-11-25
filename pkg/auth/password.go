package auth

import (
	"errors"
)

var (
	ErrEmptyPassword    = errors.New("password cannot be empty")
	ErrPasswordTooShort = errors.New("password is too short")
)

// TODO: Реализовать HashPassword(password string) (string, error)
// - Проверить что пароль не пустой, вернуть ErrEmptyPassword если пусто
// - Использовать bcrypt.GenerateFromPassword с bcrypt.DefaultCost
// - Вернуть хешированный пароль как string или ошибку
func HashPassword(password string) (string, error) {
	// Хеш сохраняйте в users.password. Исходный пароль и хеш не логируйте и не
	// возвращайте в JSON; bcrypt сам генерирует соль и включает ее в строку хеша.
	// TODO: реализовать
	return "", nil
}

// TODO: Реализовать CheckPassword(password, hash string) bool
// - Использовать bcrypt.CompareHashAndPassword для проверки пароля и хеша
// - Вернуть true если пароль совпадает, false если не совпадает или ошибка
func CheckPassword(password, hash string) bool {
	// При входе сравнивайте пароль с хешем, а не хешируйте пароль заново: соль
	// делает результаты двух вызовов GenerateFromPassword различными.
	// TODO: реализовать
	return false
}

// TODO: Реализовать ValidatePasswordStrength(password string) error (опционально)
// - Проверить минимальную длину (8 символов)
// - Проверить наличие различных типов символов (заглавные, строчные, цифры, спецсимволы)
// - Вернуть ошибку если пароль не соответствует требованиям
func ValidatePasswordStrength(password string) error {
	// TODO: реализовать (опционально)
	return nil
}
