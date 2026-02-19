# Database DSN from environment
DATABASE_DSN ?= ${DATABASE_DSN}

# Run the server
run:
	go run ./cmd/api

# Build binary
build:
	go build -o bin/api ./cmd/api

# Run migrations up
migrate-up:
	migrate -path ./migrations -database "${DATABASE_DSN}" up

# Run migrations down
migrate-down:
	migrate -path ./migrations -database "${DATABASE_DSN}" down

# Create new migration file
migrate-create:
	migrate create -seq -ext=.sql -dir=./migrations name=$(name)

# Clean binary
clean:
	rm -rf bin
