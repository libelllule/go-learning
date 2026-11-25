# Изучение языка Go (Golang)
 
Мой учебный репозиторий: код из Tour of Go, решения Exercism и алгоритмические задачи.
 
## Структура
 
| Папка | Что внутри |
| --- | --- |
| `01_…` – `16_…` | Tour of Go, раздел Basics |
| `101_…` – `113_…` | Tour of Go, раздел Flow control |
| `201_…` – `226_…` | Tour of Go, раздел More types (указатели, структуры, слайсы, мапы, функции-значения, замыкания) |
| `301_…` – `316_…` | Tour of Go, раздел Methods and interfaces (методы, интерфейсы, type assertions и type switches) |
| `Exercism/` | решения Exercism, трек Go (у каждой задачи свои тесты) |
| `leetcode/` | алгоритмические задачи, решения в пакетах-библиотеках |
 
В коде комментарии не длиннее 1–3 строк, своими словами.
 
## Прогресс
 
- [x] Tour: Basics
- [x] Tour: Flow control
- [x] Tour: More types
- [ ] Tour: Methods and interfaces (осталось: сохранить type switches, Stringers, Errors, Readers, Images)
- [ ] Tour: Generics
- [ ] Tour: Concurrency
- [ ] Тесты: свои тесты для `leetcode/twosum`
- [ ] Мини-проект: консольный менеджер задач с JSON и тестами
- [ ] Проект №1: REST-сервис на Go и PostgreSQL
- [ ] Docker и CI для проекта №1
Exercism: 14 задач решено. LeetCode: Two Sum, Valid Parentheses.
 
## Как запускать
 
```bash
go run ./01_hello          # одна программа из Tour
cd Exercism/hamming && go test   # тесты одного упражнения Exercism
go test ./leetcode/...     # тесты алгоритмических задач
gofmt -l .                 # показать неотформатированные файлы
go vet ./...               # проверка типичных ошибок
```

 
## Правила
 
- LeetCode: таймер 25–30 минут. Сначала работает, потом один проход «почище», дальше разбор и следующая задача.
- Exercism: ментор для 1–2 упражнений в неделю, остальные решаю и отправляю.
- Тесты обязательны в проектах: 3–5 табличных тестов на ключевые функции.
## Шпаргалка по командам
 
- `go run main.go`: скомпилировать и запустить файл во временной папке (без сохранения бинарника).
- `go build main.go`: собрать самостоятельный исполняемый файл (`./main`, работает без установленного Go).
- `go mod init <имя>`: создать модуль (файл `go.mod`).
- `go mod tidy`: подтянуть зависимости и убрать лишние (в `go.mod` у `golang.org/x/tour` сейчас стоит `// indirect`, после `tidy` метка исчезнет).
- `go test ./...`: запустить все тесты; `-v` подробно, `-run Имя` один тест, `-cover` покрытие, `-race` гонки.
- `go install golang.org/x/website/tour@latest`, затем `tour`: локальный Tour of Go.
 
