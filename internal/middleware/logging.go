package middleware

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
	logger *log.Logger
}

func NewLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

// TODO: Реализовать метод Logger(next http.Handler) http.Handler
// Логировать в консоль каждый запрос: IP-адрес, метод, путь, статус-код и время
// выполнения. Не логировать Authorization или тело запроса с паролем.
func (m *LoggingMiddleware) Logger(next http.Handler) http.Handler {
	// Обертка ResponseWriter должна фиксировать статус при WriteHeader; если
	// handler ничего не записал, HTTP подразумевает статус 200.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// TODO: Реализовать метод Recovery(next http.Handler) http.Handler
// Перехватить панику используя defer и recover()
// Залогировать ошибку и вернуть 500 Internal Server Error
func (m *LoggingMiddleware) Recovery(next http.Handler) http.Handler {
	// Не раскрывайте клиенту детали panic. Если заголовки уже отправлены,
	// нельзя безопасно записать новый ответ со статусом 500.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// TODO: Реализовать метод CORS(next http.Handler) http.Handler
// Добавить CORS заголовки (Allow-Origin, Allow-Methods, Allow-Headers)
// Обработать OPTIONS запрос
func (m *LoggingMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// responseWriter обертка для ResponseWriter чтобы перехватить статус код
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
