package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Config содержит конфигурацию базы данных
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// TODO: Реализовать NewPostgresDB()
// Функция должна:
// 1. Сформировать DSN строку используя GetDSN(cfg)
// 2. Открыть подключение используя sql.Open("postgres", dsn)
//   - Проверить ошибку открытия подключения
//
// 3. Проверить подключение используя db.Ping()
//   - Если ошибка - вернуть fmt.Errorf("failed to ping database: %w", err)
//
// 4. Настроить пул соединений:
//   - db.SetMaxOpenConns(25) - максимум открытых соединений
//   - db.SetMaxIdleConns(25) - максимум неиспользуемых в пуле
//   - db.SetConnMaxLifetime(5 * time.Minute) - время жизни соединения
//
// 5. Залогировать "Connected to PostgreSQL database" используя log.Println()
// 6. Вернуть *sql.DB
//
// Нюансы:
// - sql.Open создает пул лениво и не гарантирует, что сервер доступен; нужен Ping.
// - Если Ping завершается ошибкой, закройте созданный db перед возвратом.
// - Не логируйте DSN целиком: в нем находится пароль.
// - Для запросов приложения используйте один общий *sql.DB.
func NewPostgresDB(cfg Config) (*sql.DB, error) {
	// TODO: реализовать
	return nil, nil
}

// TODO: Реализовать GetDSN()
// Функция формирует Data Source Name строку для подключения к PostgreSQL
// 1. Используйте fmt.Sprintf() для форматирования
// 2. Формат DSN: "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
// 3. Подставьте значения из Config в нужном порядке:
//   - cfg.Host (например, "localhost")
//   - cfg.Port (например, 5432)
//   - cfg.User (например, "postgres")
//   - cfg.Password
//   - cfg.DBName (например, "blogdb")
//   - cfg.SSLMode (например, "disable")
//
// 4. Вернуть построенную строку
// Пример результата: "host=localhost port=5432 user=postgres password=secret dbname=blogdb sslmode=disable"
// При формировании DSN корректно обработайте специальные символы в значениях
// (например, пробелы и кавычки); пароль не должен попадать в логи.
func GetDSN(cfg Config) string {
	// TODO: реализовать
	return ""
}

// TODO: Реализовать Close()
// Закрывает соединение с базой данных
// 1. Проверить что db != nil
// 2. Вызвать db.Close()
// 3. Залогировать "Database connection closed" используя log.Println()
// Текущая сигнатура не возвращает ошибку: при реализации как минимум залогируйте
// сбой Close; для более явной обработки можно изменить сигнатуру на error.
// Закрытие соединения — часть корректного завершения приложения.
func Close(db *sql.DB) {
	// TODO: реализовать
}

// TODO: Реализовать TestConnection()
// Выполняет тестовый запрос к БД для проверки подключения
// Вернуть ошибку если подключение не работает, nil если успешно
// Учитывайте nil db и используйте PingContext либо короткий запрос с контекстом,
// чтобы проверка не зависала бесконечно.
func TestConnection(db *sql.DB) error {
	// TODO: реализовать
	return nil
}
