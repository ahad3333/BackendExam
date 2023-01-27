migration-up:
	migrate -path ./migration/postgres/ -database 'postgres://postgres:0003@localhost:3003/catalog?sslmode=disable' up

migration-down:
	migrate -path ./migration/postgres/ -database 'postgres://postgres:0003@localhost:3003/catalog?sslmode=disable' down

swag:
	export PATH=$(go env GOPATH)/bin:$PATH
	
swag-gen:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go 

migrate:
	migrate create -ext sql -dir ./migrations/postgres -seq -digits 2 create_table


