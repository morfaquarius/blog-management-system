package handler

import (
	"advanced-blog-management-system/internal/logger"
	"advanced-blog-management-system/internal/service"
	"net/http"
)

type CommentHandler struct {
	commentService service.CommentServiceInterface
	eventLogger    *logger.EventLogger
}

func NewCommentHandler(commentService service.CommentServiceInterface, eventLogger *logger.EventLogger) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
		eventLogger:    eventLogger,
	}
}

// Create обрабатывает POST /posts/{id}/comments
// TODO: Получить userID из контекста, извлечь postID, распарсить JSON, валидировать,
// вызвать commentService.Create(). Вернуть CommentResponse (201) или ошибку
// ID бери из chi.URLParam(r, "id") и проверяй через strconv.Atoi; тело содержит
// {"text":"..."}, а post_id из JSON принимать нельзя.
// После успешного создания передайте в eventLogger ID автора и комментария.
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByPost обрабатывает GET /posts/{id}/comments?limit=10&offset=0
// TODO: Извлечь postID, получить limit и offset, валидировать,
// вызвать commentService.GetByPost(). Вернуть список комментариев и общее количество
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// respondWithJSON отправляет JSON ответ
// TODO: Установить Content-Type, WriteHeader, закодировать JSON
func (h *CommentHandler) respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// TODO: реализовать
}

// respondWithError отправляет JSON ошибку
// TODO: Создать ErrorResponse и отправить используя respondWithJSON()
func (h *CommentHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}
