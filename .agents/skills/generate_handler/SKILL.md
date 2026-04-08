---
name: generate_handler
description: Generates Echo HTTP handler files that expose CRUD endpoints for a given entity.
---

# generate_handler Skill

## Purpose
Create a handler struct with Echo route functions for an entity, placed under `app/handler/` and register the routes in the main handler file.

## Inputs
- `entity_name` (string, required): e.g., `Product`.
- `route_prefix` (string, optional, default: `/api/{{entity_snake}}s`).
- `package` (string, optional, default: `handler`).

## Steps
1. Convert `entity_name` to PascalCase (`Entity`) and snake_case (`entity_snake`).
2. Define a struct `{{Entity}}Handler` with a field `{{entity}}Service interfaces.{{Entity}}ServiceInterface`.
3. Implement constructor `New{{Entity}}Handler(service interfaces.{{Entity}}ServiceInterface) interfaces.{{Entity}}HandlerInterface`.
4. Create Echo handler functions:
   - `GetList() echo.HandlerFunc`
   - `GetDetail() echo.HandlerFunc`
   - `Create() echo.HandlerFunc`
   - `Update() echo.HandlerFunc`
   - `Delete() echo.HandlerFunc`
   Each function should:
   * Initialise an `ApiResponse`.
   * Bind request DTOs.
   * Call the corresponding service method.
   * Return JSON response or error via the shared `errorResponse` helper.
5. Append route registrations to `app/handler/main_handler.go` (or a dedicated router file) using the provided `route_prefix`.
6. Write the handler file to `app/handler/{{entity_snake}}_handler.go`.
7. Run `go fmt ./...`.

## Output
- New file `app/handler/<entity_snake>_handler.go` containing the handler implementation.
- Updated route registration in the main handler file.

## Example
```yaml
entity_name: Product
route_prefix: /api/products
```
Will generate `app/handler/product_handler.go` with CRUD Echo handlers and register them under `/api/products`.
