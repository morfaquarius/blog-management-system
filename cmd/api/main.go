package main

import (
	"log"

	"advanced-blog-management-system/internal/handler"
	"advanced-blog-management-system/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Реализуйте последовательность старта: .env/env config -> DB -> migrations
	// -> JWT/repositories/services -> EventLogger/handlers/middleware -> router
	// -> HTTP server. На каждом ошибочном шаге освободите ранее созданные ресурсы.
	// TODO: Загрузить .env используя godotenv.Load()
	// Обработать ошибку: warning если файла нет, так как переменные могут быть в системе

	// TODO: Разобрать переменные окружения в конфигурацию приложения.
	// Нужны SERVER_HOST/PORT, DB_HOST/PORT/USER/PASSWORD/NAME/SSLMODE,
	// JWT_SECRET/EXPIRY_HOURS, MIGRATIONS_DIR и EVENT_LOG_PATH; проверяйте ошибки
	// парсинга чисел и не печатайте секреты.

	// TODO: Прочитать MIGRATIONS_DIR (обычно migrations локально и
	// /app/migrations в контейнере) и передать его database.Migrate; не строить
	// путь молча от текущей рабочей директории.

	// TODO: Подключиться к PostgreSQL БД
	// Использовать database.NewPostgresDB() и database.Migrate(db, migrationsDir)

	// TODO: Инициализировать JWTManager
	// auth.NewJWTManager(jwtSecret, jwtExpiryHours) возвращает (*JWTManager, error)

	// TODO: Передать один общий *sql.DB конструкторам NewUserRepo, NewPostRepo
	// и NewCommentRepo из internal/repository. Их реальные имена важны: это не
	// NewUserRepository/NewPostRepository.

	// TODO: Инициализировать сервисы
	// NewUserService(userRepo, jwtManager)
	// NewPostService(postRepo, userRepo)
	// NewCommentService(commentRepo, postRepo)

	// TODO: Создать EventLogger с EVENT_LOG_PATH (по умолчанию
	// ./logs/events.log), обработать ошибку конструктора и вызвать Start()
	// до приема HTTP-запросов. Worker пишет каждое событие в консоль и файл.
	// TODO: Инициализировать handlers
	// NewAuthHandler(userService), NewPostHandler(postService, eventLogger), NewCommentHandler(commentService, eventLogger)

	// TODO: Создать LoggingMiddleware и AuthMiddleware с общими logger/JWTManager;
	// не создавать экземпляр middleware на каждый запрос.

	// TODO: Настроить HTTP роутер через setupRouter()

	// TODO: Создать http.Server с адресом SERVER_HOST:SERVER_PORT и таймаутами
	// чтения/записи/заголовков. Запустите ListenAndServe в goroutine и отдельно
	// обработайте http.ErrServerClosed как штатное завершение.

	// TODO: Установить обработчик SIGINT/SIGTERM через signal.NotifyContext и
	// удерживать main до сигнала или фатальной ошибки сервера.

	// TODO: Graceful shutdown: Server.Shutdown с timeout 30 секунд, затем
	// EventLogger.Stop (дописать очередь и закрыть файл), затем database.Close.
	// При истечении timeout явно обработайте ошибку и не оставляйте ресурсы.

	log.Println("Server starting... (TODO: implement main.go)")
}

func setupRouter(authHandler *handler.AuthHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler, loggingMW *middleware.LoggingMiddleware, authMW *middleware.AuthMiddleware) *chi.Mux {
	// TODO: Создать chi роутер и вернуть его после регистрации всех endpoint;
	// setupRouter получает готовые зависимости и не создает новые.

	// TODO: Зарегистрировать Recovery, request Logger и CORS как общие
	// middleware. Выберите порядок так, чтобы panic внутри downstream handler
	// перехватывался Recovery, а logger видел итоговый статус.

	// TODO: Зарегистрировать публичные endpoint:
	// POST /register, POST /login, GET /health, GET /posts,
	// GET /posts/{id}, GET /posts/{id}/comments.
	// GET /health должен вернуть Content-Type application/json, статус 200 и
	// JSON-объект {"status":"ok"}.
	// Маршрут /posts/author/{authorID} в PostHandler — необязательное расширение.

	// TODO: В отдельной защищенной группе зарегистрировать POST /posts и
	// POST /posts/{id}/comments с AuthMiddleware. Убедитесь, что middleware не
	// применяется к публичным маршрутам; handler сам пишет событие после успеха.

	// TODO: Вернуть настроенный *chi.Router

	return nil
}
