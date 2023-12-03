DB_URL=postgres://postgres:project-management-password@localhost:5432/project-management?sslmode=disable


dev:
	air
.PHONY: dev

sqlc-gen:
	sqlc generate
.PHONY: sqlc-gen

migration-up:
	migrate -path ./migration/ -database "$(DB_URL)" -verbose up
.PHONY: migration-up

migration-up1:
	migrate -path ./migration/ -database "$(DB_URL)" -verbose up 1
.PHONY: migration-up1

migration-down:
	migrate -path ./migration/ -database "$(DB_URL)" -verbose down
.PHONY: migration-down

migration-down1:
	migrate -path ./migration/ -database "$(DB_URL)" -verbose down 1
.PHONY: migration-down1



