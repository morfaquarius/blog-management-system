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

// Create обрабатывает защищенный POST /posts/{id}/comments. Возьмите authorID
// из auth context, postID — из chi.URLParam(r, "id") с проверкой strconv.Atoi,
// декодируйте {"text":"..."} и вызовите сервис. Не принимайте post_id/author_id
// из JSON. Верните CommentResponse со статусом 201; событие в eventLogger
// добавляйте только после успешной записи, используя ID автора и комментария.
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByPost обрабатывает публичный GET /posts/{id}/comments. Проверьте
// положительный postID и пагинацию (defaults limit=10, offset=0; limit 1-100),
// вызовите сервис и верните {"items":[...],"total":N}.
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// respondWithJSON устанавливает application/json, записывает статус и кодирует
// один JSON-ответ; после записи статуса не отправляйте другой ответ.
func (h *CommentHandler) respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// TODO: реализовать
}

// respondWithError использует общую форму ответа; доменные ошибки передавайте
// в HandleServiceError для выбора корректного HTTP-статуса.
func (h *CommentHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}
