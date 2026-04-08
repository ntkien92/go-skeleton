---
name: generate_model
description: Generates a new GORM model struct file for a given entity.
---

# generate_model Skill

## Purpose
Create a Go struct representing a database entity, placed under `app/model/`.

## Inputs
- `entity_name` (string, required): Name of the entity, e.g., `Product`.
- `fields` (list of strings, required): Each field in the form `Name:Type` (e.g., `Title:string`).
- `package` (string, optional, default: `model`).

## Steps
1. Convert each `Name:Type` into a Go struct field with appropriate GORM tags.
2. Embed `model.ModelUUID` to provide `ID`, `CreatedAt`, `UpdatedAt` fields.
3. Assemble the struct source code.
4. Write the file to `app/model/<entity_snake>.go`.
5. Run `go fmt ./...` to format the generated code.

## Output
- New file `app/model/<entity_snake>.go` containing the struct definition.

## Example
```yaml
entity_name: Product
fields:
  - Name:string
  - Price:float64
  - Description:string
```
Will generate `app/model/product.go` with:
```go
package model

type Product struct {
    ModelUUID
    Name        string
    Price       float64
    Description string
}
```
