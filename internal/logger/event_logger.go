package logger

// EventLogger должен буферизовать события в канале и записывать их в файл
// одним worker-ом. Файл по умолчанию называется logs.txt.
type EventLogger struct {
	// TODO: Добавить eventsChan chan string, doneChan chan struct{},
	// file *os.File, sync.RWMutex и stopped bool. Защитите запуск и остановку
	// через sync.Once, чтобы Start/Stop были безопасны при повторном вызове.
}

// TODO: Реализовать NewEventLogger()
// Открыть файл (создать, если не существует, и дописывать в конец), создать
// родительский каталог при необходимости, eventsChan емкостью 100 и канал
// завершения. Для файла используй O_CREATE|O_APPEND|O_WRONLY и права 0644.
// Вернуть ошибку, если файл открыть или создать не удалось.
func NewEventLogger(filePath string) (*EventLogger, error) {
	// TODO: реализовать
	return nil, nil
}

// LogEvent неблокирующе отправляет событие в очередь; если очередь заполнена,
// выведи предупреждение в консоль, не задерживая HTTP-обработчик.
func (el *EventLogger) LogEvent(event string) {
	// Возьмите RLock, проверьте stopped, затем сделайте select с отправкой в
	// eventsChan и default; отпустите lock во всех ветках. Stop под Lock ставит
	// stopped=true и закрывает канал, так что send в закрытый канал невозможен.
	// Формат события: "user <userID> created post <postID>" либо
	// "user <userID> created comment <commentID>". Вызывайте этот метод только
	// после успешного сохранения записи.
	// TODO: реализовать
}

// TODO: Реализовать Start()
// Запустить ровно одну worker-горутины при старте приложения.
func (el *EventLogger) Start() {
	// Вызывайте до ListenAndServe. Запустите только одну горутину; Stop
	// вызывается только после Start.
	// TODO: реализовать
}

// TODO: Реализовать worker()
// Последовательно читать очередь; перед записью каждого события подождать
// time.Second. Дописывать строки вида
// "[2006-01-02 15:04:05] user 5 created post 12". При закрытии очереди
// обработать все принятые события, закрыть файл и уведомить Stop().
// При ошибке записи или закрытия файла выведите ошибку в консоль.
func (el *EventLogger) worker() {
	// TODO: реализовать
}

// TODO: Реализовать Stop()
// Идемпотентно прекратить прием событий, закрыть очередь, дождаться обработки
// накопленных событий и закрытия файла. Вызывать при graceful shutdown после
// остановки HTTP-сервера.
func (el *EventLogger) Stop() {
	// Продумайте синхронизацию LogEvent/Stop и защиту от повторного close.
	// TODO: реализовать
}
