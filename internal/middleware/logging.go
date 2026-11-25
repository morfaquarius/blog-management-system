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

// TODO: Реализовать Logger как middleware для каждого HTTP-запроса.
// Оберните ResponseWriter, запомните фактический статус (по умолчанию 200),
// измерьте время обработки и запишите IP/метод/путь/статус/duration через
// переданный logger. Не записывайте заголовок Authorization или тело запроса.
func (m *LoggingMiddleware) Logger(next http.Handler) http.Handler {
	// Обертка ResponseWriter должна фиксировать статус при WriteHeader; если
	// handler ничего не записал, HTTP подразумевает статус 200.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// TODO: Реализовать Recovery как middleware на внешнем уровне роутера.
// defer/recover должен перехватить panic, записать причину в серверный лог,
// не раскрывая stack trace клиенту, и вернуть HTTP 500, если ответ еще
// не начал отправляться. Не пытайтесь писать второй статус после WriteHeader.
func (m *LoggingMiddleware) Recovery(next http.Handler) http.Handler {
	// Не раскрывайте клиенту детали panic. Если заголовки уже отправлены,
	// нельзя безопасно записать новый ответ со статусом 500.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// TODO: Реализовать CORS middleware и согласовать разрешенные origin с
// окружением/требованиями API. Установите нужные CORS-заголовки до next;
// OPTIONS preflight обработайте без вызова handler и завершите подходящим
// успешным статусом. Не отражайте произвольный Origin без проверки.
func (m *LoggingMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: реализовать
		next.ServeHTTP(w, r)
	})
}

// responseWriter фиксирует первый статус, записанный downstream handler.
// Учитывайте, что если Write вызван без WriteHeader, net/http отправляет 200.
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
