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

// Create обрабатывает защищенный POST /posts: возьмите authorID из auth context,
// декодируйте/провалидируйте PostCreateRequest и вызовите postService.Create.
// Верните PostResponse со статусом 201. Не принимайте author_id из тела JSON.
// После успешного сохранения отправьте в eventLogger событие с ID пользователя
// и ID созданного поста; при ошибке БД не регистрируйте событие.
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByID обрабатывает публичный GET /posts/{id}: извлеките положительный ID
// через chi.URLParam и strconv.Atoi, вызовите сервис и верните PostResponse
// (200) либо общую JSON-ошибку. Этот маршрут не требует JWT.
func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetAll обрабатывает публичный GET /posts?limit=10&offset=0. При отсутствии
// параметров используйте limit=10 и offset=0; при наличии разбирайте числа,
// проверяйте limit 1-100 и offset >= 0. Не заменяйте ошибочный ввод default.
// Ответ содержит {"items":[...],"total":N}.
func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// GetByAuthor обслуживает необязательный GET
// /posts/author/{authorID}?limit=10&offset=0. Реализуйте и регистрируйте его
// только после обязательных маршрутов; проверьте, что chi не трактует
// "author" как параметр {id} маршрута /posts/{id}.
func (h *PostHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	// TODO: реализовать
}

// respondWithJSON устанавливает application/json, записывает статус и кодирует
// один JSON-ответ. Не вызывайте его после того, как уже записан другой ответ.
func (h *PostHandler) respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	// TODO: реализовать
}

// respondWithError использует общую форму ответа; ошибки сервиса сначала
// сопоставляйте с HTTP-статусом через HandleServiceError.
func (h *PostHandler) respondWithError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}
