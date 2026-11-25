package middleware

import "net/http"

// TODO: Реализовать ToMiddleware() в internal/middleware/utils.go.
// chi.Router.Use ожидает func(http.Handler) http.Handler, тогда как RequireAuth
// сейчас имеет сигнатуру func(http.HandlerFunc) http.HandlerFunc. Адаптер нужен
// для подключения этой существующей сигнатуры в chi.
//
// Оберните handler.ServeHTTP в http.HandlerFunc, передайте его в fn и верните
// получившуюся функцию как http.Handler. Не вызывайте downstream handler при
// создании middleware — только при обработке конкретного запроса.
//
// Пример подключения в cmd/api/main.go:
// router.Use(middleware.ToMiddleware(authMW.RequireAuth))
func ToMiddleware(fn func(http.HandlerFunc) http.HandlerFunc) func(http.Handler) http.Handler {
	// Обертка должна вызвать результат fn ровно один раз, передав исходные
	// ResponseWriter и Request.
	// TODO: реализовать
	return nil
}
