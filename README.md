# Effective Mobile Task - Go REST API

REST API для выполнения CRUD операций с подписками. Реализован на **Go + Gin + GORM + Postgres** с JWT авторизацией.

## Быстрый старт

```bash
# 1. Клонировать репозиторий
git clone https://github.com/1acrimosa/effective-mobile-task
cd effective-mobile-task

# 2. Установить зависимости
go mod tidy

# 3. Запустить сервер
go run cmd/app/main.go
```

Сервер запустится на http://localhost:8080

## Полный CRUD тест:

```bash
# 1. Создать подписку
curl -X POST http://localhost:8080/api/items \
  -H "Content-Type: application/json" \
  -d '{"price":"400","vendor":"Яндекс"}'

# 2. Получить все подписки
curl http://localhost:8080/api/items

# 3. Получить одну подписку
curl http://localhost:8080/api/items/1

# 4. Обновить подписку
curl -X PUT http://localhost:8080/api/items/1 \
  -H "Content-Type: application/json" \
  -d '{"price":"500","vendor":"Яндекс+"}'

# 5. Удалить подписку
curl -X DELETE http://localhost:8080/api/items/1
```

## Health Check:

```bash
GET http://localhost:8080/health
```

## Docker:

### Сборка и запуск
docker-compose up --build

### Остановка
docker-compose down








