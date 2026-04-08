---
name: generate_repository
description: Generates a repository interface and GORM implementation for a given entity.
---

# generate_repository Skill

## Purpose
Create a repository layer that abstracts database operations for an entity, placed under `app/repository/`.

## Inputs
- `entity_name` (string, required): e.g., `Product`.
- `package` (string, optional, default: `repository`).

## Steps
1. Derive the snake_case table name from `entity_name`.
2. Define an interface `{{Entity}}Repository` with methods:
   - `GetList(ctx context.Context, preload []string) ([]model.{{Entity}}, []error)`
   - `GetDetail(ctx context.Context, query model.GetDetail{{Entity}}QueryParams) (*model.{{Entity}}, []error)`
   - `Create(ctx context.Context, data model.{{Entity}}) (string, []error)`
   - `Update(ctx context.Context, id string, data model.{{Entity}}) ([]error)`
   - `Delete(ctx context.Context, id string) ([]error)`
3. Implement the interface using GORM (`db *gorm.DB`).
4. Write the file to `app/repository/{{entity_snake}}_repository.go`.
5. Run `go fmt ./...`.

## Output
- New file `app/repository/<entity_snake>_repository.go` containing both the interface and its implementation.
