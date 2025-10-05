.PHONY: build-ui install-ui-dependencies run-ui run-backend build-backend build go-lint

build-ui:
	cd ui && pnpm build

run-backend:
	go run cmd/app/app.go

run-ui:
	cd ui && pnpm dev

build-backend:
	go build -o app cmd/app/app.go

build: build-ui build-backend

go-lint:
	go fmt ./...
	go vet ./...