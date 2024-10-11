include $(PWD)/.env

run:
	@go run *.go

dbo:
	@docker start pg-container >> /dev/null || systemctl start docker.service 

dbs: dbo
	@goose -dir $(PWD)/database/migrations postgres "postgres://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/ctrl" status

dbc: dbo
	@./database/migrations/create_migration.sh

dbu: dbo
	@goose -dir $(PWD)/database/migrations postgres "postgres://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/ctrl" up
