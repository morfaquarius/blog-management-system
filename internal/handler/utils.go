package handler

import (
	"net/http"
)

// ErrorResponse представляет структуру ошибки в ответе
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// TODO: Реализовать WriteError как общий отправитель JSON-ошибок handler-пакета.
// До WriteHeader установите Content-Type: application/json; Error заполните
// через http.StatusText(statusCode), Message — безопасным сообщением для API.
// Не пишите тело внутренней ошибки БД/stack trace и не вызывайте после уже
// отправленного ответа.
func WriteError(w http.ResponseWriter, message string, statusCode int) {
	// TODO: реализовать
}

// TODO: Реализовать HandleServiceError: сопоставить доменные ошибки сервиса и
// validator.ValidationErrors с API-статусами, затем вызвать WriteError ровно раз.
// Используйте errors.Is/errors.As для обернутых ошибок, а не сравнение текста.
// Для каждого типа ошибки из apperrors вернуть нужный HTTP статус:
//   - ErrUserAlreadyExists -> 409 Conflict
//   - ErrInvalidCredentials -> 401 Unauthorized
//    - ErrUserNotFound -> 404 Not Found
//    - ErrPostNotFound -> 404 Not Found
//    - ErrCommentNotFound -> 404 Not Found
//    - ErrInvalidPostID -> 400 Bad Request
//    - ErrForbidden -> 403 Forbidden
//    - ErrUnauthorized -> 401 Unauthorized
//
// Для validation errors верните 400 и понятное безопасное описание; для
// неизвестных ошибок верните общий 500. Логирование подробной причины — задача
// серверного слоя; внешний ответ не должен раскрывать SQL, DSN или stack trace.
func HandleServiceError(w http.ResponseWriter, err error) {
	// TODO: реализовать
}
