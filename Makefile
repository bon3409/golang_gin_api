# 取得.env 的變數
# 如果在 docker container 中要使用的話，下面這段要註解，才可以取得 docker env variable
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# 執行 make 指令如果有帶參數，不會出錯
# 參考資料：https://stackoverflow.com/a/6273809
%:
	@:

.PHONY: docker-run
docker-run:
	docker compose build
	docker compose up

.PHONY: docker-stop
docker-stop:
	docker compose down

# 建立 migrate table
# make migrate-create <file name>
.PHONY: migrate-create
migrate-create:
	migrate create -ext sql -dir ./db/migrate -seq $(filter-out $@,$(MAKECMDGOALS))

# 執行 migrate up
.PHONY: migrate-up
migrate-up:
	migrate -verbose -path db/migrate -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}" up

# 執行 migrate down
.PHONY: migrate-down
migrate-down:
	migrate -verbose -path db/migrate -database "mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:${DATABASE_PORT})/${DATABASE_NAME}" down
