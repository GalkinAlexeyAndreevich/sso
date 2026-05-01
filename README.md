# SSO Auth Service (Учебный проект)

Небольшой учебный сервис авторизации на Go с gRPC API, SQLite-хранилищем и JWT-токенами.

## Что умеет сервис

- `Register` - регистрация пользователя по `email` и `password`
- `Login` - вход пользователя и выдача JWT-токена
- `IsAdmin` - проверка, является ли пользователь администратором

## Технологии

- Go `1.24+`
- gRPC
- SQLite (`modernc.org/sqlite`)
- JWT (`github.com/golang-jwt/jwt/v5`)
- Миграции (`golang-migrate`)
- Тесты (`testing` + `testify`)

## Требования

- Установленный Go `1.24+`
- (Опционально) [Task](https://taskfile.dev/) для удобного запуска команд

## Быстрый старт

1) Установить зависимости:

```bash
go mod download
```

2) Применить миграции:

```bash
go run cmd/migrator/main.go -storage-path ./storage/sso.db -migrations-path ./migrations -migrations-table migrations
```

или через Task:

```bash
task migrate
```

3) Запустить сервис:

```bash
go run cmd/sso/main.go -config ./config/config.yaml
```

или через Task:

```bash
task start
```

По умолчанию gRPC-сервер стартует на порту `44044`.

## Конфигурация

Основной конфиг: `config/config.yaml`.

Пример:

```yaml
env: "local"
storage_path: "./storage/sso.db"
token_ttl: 1h
grpc:
  port: 44044
  timeout: 10m
```

Сервис принимает путь к конфигу через:

- флаг `-config`
- или переменную окружения `CONFIG_PATH`

## Тесты

Запуск интеграционных тестов:

```bash
go test -v -count=1 ./tests/...
```

или через Task:

```bash
task test
```

Перед первым запуском тестов при необходимости можно накатить тестовые миграции:

```bash
task test-migrate
```

## Структура проекта

- `cmd/sso` - точка входа сервиса
- `cmd/migrator` - утилита для применения миграций
- `internal/services/auth` - бизнес-логика авторизации
- `internal/grpc/auth` - gRPC-обработчики
- `internal/storage/sqlite` - реализация хранилища на SQLite
- `migrations` - миграции базы данных
- `tests` - интеграционные тесты

## Примечания

- Проект учебный и предназначен для демонстрации базовой архитектуры auth/SSO-сервиса.
- Для production-сценариев рекомендуется добавить:
  - rate limiting / brute-force protection,
  - refresh token flow,
  - аудит и трассировку,
  - ротацию секретов и более строгую политику безопасности.
