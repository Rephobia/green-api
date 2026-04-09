# GREEN-API Test Task

Тестовое задание (GREEN-API).

Проект представляет собой веб-приложение, которое позволяет взаимодействовать с методами GREEN-API через веб-интерфейс:
- `getSettings`
- `getStateInstance`
- `sendMessage`
- `sendFileByUrl`

## Основные возможности

- **Backend на Go** — чистая архитектура, разделение на слои (`internal/handler`, `internal/service`, `internal/dto`, `internal/middleware`).
- **Простой SPA-фронтенд** — одностраничное приложение (HTML + Tailwind CSS + Vanilla JS), полностью embedded в бинарник.
- **Единый бинарный файл** — фронтенд раздаётся самим Go-приложением (через `embed.FS`).

### Реализация

- **Graceful Shutdown** — корректное завершение работы сервера с ожиданием завершения текущих запросов.
- **Middleware**:
  - Logging (логирование всех входящих запросов с методом, путём и временем выполнения)
  - Recovery (защита от паник, возврат 500 ошибки с сообщением)
  - Validate (универсальная валидация запросов)
- **Валидация реквестов** с использованием библиотеки `github.com/go-playground/validator/v10`
- **Конфигурация** через библиотеку `github.com/ilyakaznacheev/cleanenv` (поддержка `.env` + YAML)
- **Чистая обработка ошибок** с единым форматом ответов
- **Структура проекта** соответствует принципам Clean Architecture / Layered Architecture
- **Task runner** - Taskfile `https://github.com/go-task/task`

## Как запустить локально

```bash
# Клонировать репозиторий
git clone https://github.com/Rephobia/green-api.git
cd green-api

# Задать .env

cp .env.example .env

# Установить зависимости
go mod tidy

# Запустить
task serve
