package apperrors

// TODO: Определить экспортируемые ошибки приложения в этом файле пакета
// apperrors. Их возвращают service-методы, а handler/utils.go преобразует их
// в стабильные HTTP-коды. Используйте errors.New() для каждой ошибки.
// Необходимы следующие ошибки:
// - ErrUnauthorized - пользователь не аутентифицирован
// - ErrForbidden - пользователь не имеет прав доступа
// - ErrUserNotFound - пользователь не найден в БД
// - ErrUserAlreadyExists - пользователь с таким email/username уже существует
// - ErrInvalidCredentials - неверный email или пароль при входе
// - ErrPostNotFound - пост не найден в БД
// - ErrCommentNotFound - комментарий не найден в БД
// - ErrInvalidPostID - передан невалидный ID поста
//
// Объявите ошибки package-level переменными, чтобы errors.Is работал и после
// оборачивания через fmt.Errorf("...: %w", err). Не создавайте новую ошибку с
// тем же текстом внутри каждого метода.
//
// Эти переменные будут использоваться в:
// - Services: возвращать эти ошибки для обозначения бизнес-логики
// - Handlers: проверять ошибки через errors.Is() и возвращать нужные HTTP коды
// Оборачивайте их через fmt.Errorf("...: %w", err), иначе errors.Is не сможет
// найти исходную доменную ошибку. Не используйте текст ошибки как API-контракт.
//
// Пример создания:
// var (
//     ErrExample = errors.New("example error message")
//     ErrOther = errors.New("other error message")
// )
