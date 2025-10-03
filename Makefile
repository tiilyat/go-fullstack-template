.PHONY: build-frontend install-frontend-dependencies run-frontend run-backend build-backend build go-lint

build-frontend:
	cd frontend && pnpm build

run-backend:
	go run cmd/app/app.go

run-frontend:
	cd frontend && pnpm dev

build-backend:
	go build -o app cmd/app/app.go

build: build-frontend build-backend

go-lint:
	go fmt ./...
	go vet ./...