package auth

import (
	"errors"
)

var (
	ErrEmptyPassword    = errors.New("password cannot be empty")
	ErrPasswordTooShort = errors.New("password is too short")
)

// TODO: Реализовать HashPassword(password string) (string, error).
// Эта функция вызывается сервисом регистрации до UserRepo.Create; в БД
// сохраняется только возвращенный bcrypt-хеш. Пустой пароль должен вернуть
// ErrEmptyPassword, ошибки bcrypt нужно передать вызывающему коду с причиной.
// Используйте bcrypt.GenerateFromPassword с bcrypt.DefaultCost. Не выводите
// исходный пароль или хеш в логи и не включайте их в API-ответ.
func HashPassword(password string) (string, error) {
	// Хеш сохраняйте в users.password. Исходный пароль и хеш не логируйте и не
	// возвращайте в JSON; bcrypt сам генерирует соль и включает ее в строку хеша.
	// TODO: реализовать
	return "", nil
}

// TODO: Реализовать CheckPassword(password, hash string) bool.
// Используйте bcrypt.CompareHashAndPassword: повторный вызов HashPassword для
// введенного пароля не подходит, потому что bcrypt генерирует новую соль.
// Функция возвращает false для несовпадающего пароля и поврежденного хеша;
// Login не должен отличать эти случаи во внешнем ответе.
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
