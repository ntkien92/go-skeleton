---
description: Scaffold full CRUD for a given entity
---

# CRUD Scaffold Workflow

## Purpose
Automatically generate all layers required for a new CRUD resource (model, repository, service, and Echo handler) and register the routes in the application.

## Prerequisites
- The `.agents/skills/` directory must contain the following skills:
  - `generate_model`
  - `generate_repository`
  - `generate_service`
  - `generate_handler`
- The project builds with `go build ./...` and uses `go fmt` for formatting.

## Steps
1. **Collect entity definition** (manual step).  
   - Ask the user for `entity_name` (PascalCase, e.g., `Product`).
   - Ask for a list of fields in `Name:Type` format (e.g., `Title:string`).
2. **Generate model** – invoke the `generate_model` skill with the collected inputs.
3. **Generate repository** – invoke the `generate_repository` skill using the same `entity_name`.
4. **Generate service** – invoke the `generate_service` skill.
5. **Generate handler** – invoke the `generate_handler` skill.  
   - Optionally provide a custom `route_prefix` (default `/api/{{entity_snake}}s`).
6. **Register routes** – the `generate_handler` skill updates `app/handler/main_handler.go` (or the router file) with the new routes.
7. // turbo **Run code formatter**
   ```bash
   go fmt ./...
   ```
   This ensures the generated files are properly formatted.

## Completion
After the workflow finishes you will have:
- `app/model/<entity_snake>.go`
- `app/repository/<entity_snake>_repository.go`
- `app/service/<entity_snake>_service.go`
- `app/handler/<entity_snake>_handler.go`
- Updated route registration in the main handler file.

You can now run the application (`go run ./cmd/...` or your preferred entry point) and the new CRUD endpoints will be available.
