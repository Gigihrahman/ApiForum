DATABASE_URL ?=

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database "$(DATABASE_URL)" -path scripts/migrations up

migrate-down:
	@ migrate -database "$(DATABASE_URL)" -path scripts/migrations down

migrate-force:
	@ migrate -database "$(DATABASE_URL)" -path scripts/migrations force $(version)