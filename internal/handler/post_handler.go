package handler

import (
	"advanced-blog-management-system/internal/logger"
	"advanced-blog-management-system/internal/service"
	"net/http"
)

type PostHandler struct {
	postService service.PostServiceInterface
	eventLogger *logger.EventLogger
}

// NewPostHandler создает новый экземпляр PostHandler
func NewPostHandler(postService service.PostServiceInterface, eventLogger *logger.EventLogger) *PostHandler {
	return &PostHandler{
		postService: postService,
		eventLogger: eventLogger,
	}
}

// Create обрабатывает POST /posts
// TODO: Получить userID из контекста, распарсить JSON, валидировать, вызвать postService.Create()
// Вернуть PostResponse со статусом 201 или ошибку
// Не принимайте author_id из тела JSON: автор определяется токеном в middleware.
// После успешного сохранения отправьте в eventLogger событие с ID пользователя
// и ID созданного поста.
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByID обрабатывает GET /posts/{id} — публичный endpoint.
// TODO: Извлечь положительный ID через chi.URLParam(r, "id"),
// вызвать postService.GetByID()
// и вернуть PostResponse (200) или JSON-ошибку.
func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetAll обрабатывает GET /posts?limit=10&offset=0
// TODO: Получить limit и offset из query, валидировать (limit 1-100, default 10),
// вызвать postService.GetAll(). Вернуть список постов и общее количество
// Обрабатывайте ошибку strconv.Atoi отдельно от диапазона; не превращайте
// некорректное значение параметра в молчаливый default.
func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByAuthor — дополнительный маршрут GET /posts/author/{authorID}?limit=10&offset=0.
// Это необязательное расширение; сначала реализуйте endpoint из задания.
// TODO: Извлечь authorID из URL, получить limit и offset из query,
// вызвать postService.GetByAuthor(). Вернуть список постов автора и общее количество.
func (h *PostHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// respondWithJSON отправляет JSON ответ
// TODO: Установить Content-Type, WriteHeader, закодировать JSON
func (h *PostHandler) respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// TODO: реализовать
}

// respondWithError отправляет JSON ошибку
// TODO: Создать ErrorResponse и отправить используя respondWithJSON()
func (h *PostHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}
