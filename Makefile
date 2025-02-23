# Makefile для создания миграций

# Переменные
DB_DSN := postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат одной миграции
migrate-down:
	$(MIGRATE) down 1

# Откат всех миграций
migrate-down-all:
	$(MIGRATE) down

# Запуск сервера
run:
	go run cmd/app/main.go

# Генерация кода OpenAPI
gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

lint:
	golangci-lint run --out-format=colored-line-number