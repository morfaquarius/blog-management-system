package main

import (
	"log"
)

func main() {
	// Реализуйте запуск поэтапно: конфигурация -> БД -> миграции -> зависимости
	// -> router -> HTTP server. При ошибке старта закройте уже созданные ресурсы.
	// TODO: Загрузить .env используя godotenv.Load()
	// Обработать ошибку: warning если файла нет, так как переменные могут быть в системе

	// TODO: Реализовать загрузку конфигурации из переменных окружения
	// Config должен содержать: ServerHost, ServerPort, DB параметры, JWT параметры

	// TODO: Найти корневую папку проекта (поддержка Docker)
	// Искать папку migrations в текущей директории и выше (до 2 уровней)

	// TODO: Подключиться к PostgreSQL БД
	// Использовать database.NewPostgresDB() и database.Migrate()

	// TODO: Инициализировать JWTManager
	// auth.NewJWTManager(jwtSecret, jwtExpiryHours)

	// TODO: Инициализировать репозитории
	// Сверьте имена конструкторов с исходниками: сейчас это NewUserRepo,
	// NewPostRepo и NewCommentRepo, а не NewUserRepository/NewPostRepository.

	// TODO: Инициализировать сервисы
	// NewUserService(userRepo, jwtManager)
	// NewPostService(postRepo, userRepo)
	// NewCommentService(commentRepo, postRepo)

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
	// Завершить сервер с таймаутом 30 секунд и закрыть БД

	log.Println("Server starting... (TODO: implement main.go)")
}

func setupRouter(authHandler interface{}, postHandler interface{}, commentHandler interface{}, loggingMW interface{}, authMW interface{}) interface{} {
	// После выбора сигнатур замените interface{} конкретными типами handler-ов
	// и *chi.Router: компилятор должен проверять wiring зависимостей.
	// TODO: Создать chi роутер

	// TODO: Зарегистрировать глобальные middleware
	// Recovery, Logger, CORS

	// TODO: Зарегистрировать публичные эндпоинты
	// POST /api/register, POST /api/login, GET /api/health
	// GET /api/posts, GET /api/posts/{id}
	// GET /api/posts/{postId}/comments
	// PostHandler также содержит GetByAuthor; зарегистрируйте его отдельно и
	// проверьте, что /posts/{id} не перехватывает /posts/author/{authorID}.

	// TODO: Зарегистрировать защищенные эндпоинты (требуют AuthMiddleware)
	// POST /api/posts
	// POST /api/posts/{postId}/comments

	// TODO: Вернуть настроенный *chi.Router

	return nil
}
