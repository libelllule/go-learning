# Изучение языка Go (Golang)

Мой учебный репозиторий: код из Tour of Go, решения Exercism и алгоритмические задачи.

## Структура

| Папка | Что внутри |
| --- | --- |
| `01_…` – `16_…` | Tour of Go, раздел Basics |
| `101_…` – `113_…` | Tour of Go, раздел Flow control |
| `Exercism/` | решения Exercism, трек Go (у каждой задачи свои тесты) |
| `leetcode/` | алгоритмические задачи, по одной в день (появится) |

Теория в `NOTES.md`. В коде комментарии не длиннее 1–3 строк, своими словами.

## Прогресс

- [x] Tour: Basics
- [x] Tour: Flow control
- [ ] Tour: More types (указатели, структуры, слайсы, мапы, замыкания)
- [ ] Tour: Methods and interfaces
- [ ] Tour: Generics
- [ ] Tour: Concurrency
- [ ] Мини-проект: консольный менеджер задач с JSON и тестами
- [ ] Проект №1: REST-сервис на Go и PostgreSQL

## Как запускать

```bash
go run ./01_hello          # одна программа
cd Exercism/hamming && go test   # тесты Exercism
gofmt -l .                 # показать неотформатированные файлы
go vet ./...               # проверка типичных ошибок
```

## Шпаргалка по командам

- `go run main.go`: скомпилировать и запустить файл во временной папке (без сохранения бинарника).
- `go build main.go`: собрать самостоятельный исполняемый файл (`./main`, работает без установленного Go).
- `go mod init <имя>`: создать модуль (файл `go.mod`).
- `go test ./...`: запустить все тесты.
- `go install golang.org/x/website/tour@latest`, затем `tour`: локальный Tour of Go.
