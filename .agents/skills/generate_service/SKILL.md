---
name: generate_service
description: Generates a service layer that orchestrates repository calls and returns DTOs for a given entity.
---

# generate_service Skill

## Purpose
Create a service struct that implements the business‑logic interface for an entity, placed under `app/service/`.

## Inputs
- `entity_name` (string, required): e.g., `Product`.
- `package` (string, optional, default: `service`).

## Steps
1. Derive the snake_case name for file naming.
2. Define an interface `{{Entity}}ServiceInterface` with methods:
   - `GetList(ctx context.Context) ([]dto.{{Entity}}Response, []error)`
   - `GetDetail(ctx context.Context, req dto.Get{{Entity}}DetailRequest) (*dto.{{Entity}}Response, []error)`
   - `Create(ctx context.Context, req dto.Create{{Entity}}Request) (*dto.{{Entity}}Response, []error)`
   - `Update(ctx context.Context, id string, req dto.Update{{Entity}}Request) ([]error)`
   - `Delete(ctx context.Context, id string) ([]error)`
3. Implement a struct `{{Entity}}Service` that embeds the corresponding repository interface.
4. Wire the repository calls, map model objects to DTOs (using existing `dto` helpers), and propagate errors.
5. Write the file to `app/service/{{entity_snake}}_service.go`.
6. Run `go fmt ./...` to format the generated code.

## Output
- New file `app/service/<entity_snake>_service.go` containing the interface and its implementation.

## Example
```yaml
entity_name: Product
```
Will generate `app/service/product_service.go` with a service that calls `productRepository` methods and returns `dto.ProductResponse` objects.
