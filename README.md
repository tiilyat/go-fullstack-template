# Go Fullstack Web Application

A minimal fullstack web application template with Go backend and Vue 3 frontend, using Go's embed directive to bundle the frontend directly into the binary.

## Prerequisites

- Go 1.24.0+
- Node.js ^20.19.0 || >=22.12.0
- pnpm

## Project Structure

- `cmd/` - Application entry points
- `frontend/` - Vue 3 frontend with Vite
- `frontend/embed.go` - Frontend embedding logic

## Quick Start

Build and run:
```bash
make build
./app
```

The server will start on `http://localhost:8080`

## Development

Run backend:
```bash
make run-backend
```

Run frontend with hot reload:
```bash
make run-frontend
```

## Available Make Commands

- `make build-frontend` - Build frontend assets
- `make run-frontend` - Run frontend dev server with hot reload
- `make run-backend` - Run Go backend server
- `make build-backend` - Build Go binary
- `make build` - Build both frontend and backend

## Architecture

- Backend serves API endpoints under `/api/*`
- Frontend is served from `/` with SPA fallback routing
- All frontend assets are embedded into the Go binary using `//go:embed`

## License

MIT
