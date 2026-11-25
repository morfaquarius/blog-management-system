package main

import (
	"log"

	"advanced-blog-management-system/internal/handler"
	"advanced-blog-management-system/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Реализуйте запуск поэтапно: конфигурация -> БД -> миграции -> зависимости
	// -> router -> HTTP server. При ошибке старта закройте уже созданные ресурсы.
	// TODO: Загрузить .env используя godotenv.Load()
	// Обработать ошибку: warning если файла нет, так как переменные могут быть в системе

	// TODO: Реализовать загрузку конфигурации из переменных окружения
	// Config должен содержать: ServerHost, ServerPort, DB параметры, JWT параметры

	// TODO: Прочитать MIGRATIONS_DIR (локально migrations, в контейнере
	// /app/migrations); не зависеть от текущей рабочей директории.

	// TODO: Подключиться к PostgreSQL БД
	// Использовать database.NewPostgresDB() и database.Migrate(db, migrationsDir)

	// TODO: Инициализировать JWTManager
	// auth.NewJWTManager(jwtSecret, jwtExpiryHours) возвращает (*JWTManager, error)

	// TODO: Инициализировать репозитории
	// Сверьте имена конструкторов с исходниками: сейчас это NewUserRepo,
	// NewPostRepo и NewCommentRepo, а не NewUserRepository/NewPostRepository.

	// TODO: Инициализировать сервисы
	// NewUserService(userRepo, jwtManager)
	// NewPostService(postRepo, userRepo)
	// NewCommentService(commentRepo, postRepo)

	// TODO: Создать EventLogger с EVENT_LOG_PATH (по умолчанию
	// ./logs/events.log), вызвать Start() до запуска HTTP-сервера и обработать
	// ошибку конструктора. Каждое событие worker пишет и в консоль, и в файл.
	// TODO: Инициализировать handlers
	// NewAuthHandler(userService), NewPostHandler(postService, eventLogger), NewCommentHandler(commentService, eventLogger)

	// TODO: Инициализировать middleware
	// LoggingMiddleware, AuthMiddleware

	// TODO: Настроить HTTP роутер через setupRouter()

	// TODO: Создать и запустить HTTP сервер
	// Использовать goroutine с ListenAndServe()

	// TODO: Установить обработчик сигналов (SIGINT, SIGTERM)
	// Блокировать main пока не получен сигнал

	// TODO: Graceful shutdown
	// Завершить сервер с таймаутом 30 секунд, затем вызвать eventLogger.Stop()
	// (он должен дописать буфер событий и закрыть файл), затем закрыть БД.

	log.Println("Server starting... (TODO: implement main.go)")
}

func setupRouter(authHandler *handler.AuthHandler, postHandler *handler.PostHandler, commentHandler *handler.CommentHandler, loggingMW *middleware.LoggingMiddleware, authMW *middleware.AuthMiddleware) *chi.Mux {
	// TODO: Создать chi роутер

	// TODO: Зарегистрировать глобальные middleware
	// Recovery, Logger, CORS

	// TODO: Зарегистрировать обязательные публичные endpoint:
	// POST /register, POST /login, GET /health (JSON {"status":"ok"}),
	// GET /posts, GET /posts/{id}, GET /posts/{id}/comments.
	// GET /health должен вернуть Content-Type application/json, статус 200 и
	// JSON-объект {"status":"ok"}.
	// Маршрут /posts/author/{authorID} в PostHandler — необязательное расширение.

	// TODO: Зарегистрировать защищенные эндпоинты (требуют AuthMiddleware)
	// POST /posts
	// POST /posts/{id}/comments
	// После успешного POST обработчик передает событие в eventLogger.

	// TODO: Вернуть настроенный *chi.Router

	return nil
}
