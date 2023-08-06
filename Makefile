.PHONY:
migrate-up:
	migrate -source file://./schema/migrations -database postgres://postgres:password@localhost:5432/test_task_db?sslmode=disable up

.PHONY:
migrate-down:
	migrate -source file://./schema/migrations -database postgres://postgres:password@localhost:5432/test_task_db?sslmode=disable down


.PHONY:
create-migration:
	migrate create -ext sql -dir schema/migrations -seq ${name}