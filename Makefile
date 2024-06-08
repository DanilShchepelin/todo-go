start_live:
	make shared_start
	make app_start_live

# Запуск зависимостей для приложения (бд и прочие сервисы)
shared_start:
	docker compose -f docker-compose.shared.yml up -d

# Выключить контейнер и оставить активным процессом в консоли
app_start_live:
	docker compose up

# Выключение приложения
down:
	make app_down
	make shared_down

# Выключение зависимостей
shared_down:
	docker compose -f docker-compose.shared.yml down

# Выключить контейнер
app_down:
	docker compose down

# Пересобрать контейнер
app_build:
	sudo chmod -R 777 ./data && docker compose build