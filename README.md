# Go Fundamentals

My Go learning repo — tracking progress from zero to production.

## Progress

| # | Chapter | Status | Updated |
|---|---------|--------|---------|
| 01 | Types & Variables | ✅ Done | 2025-04-24 |
| 02 | Control Flow | ✅ Done | 2025-04-24 |
| 03 | Structs & OOP | 🔄 In Progress | 2025-04-24 |
| 04 | Concurrency | ⏳ Pending | - |
| 05 | Design Patterns | ⏳ Pending | - |
| 06 | Error Handling | ⏳ Pending | - |
| 07 | Real Projects | ⏳ Pending | - |

Legend: ✅ Done | 🔄 In Progress | ⏳ Pending

## What I learned today
> Session: 2025-04-24
- zero values — every type has a safe default
- interfaces satisfied implicitly — no implements keyword
- WaitGroup must be passed as pointer

## Key gotchas

| Mistake | Fix |
|---------|-----|
| Copied WaitGroup by value | Pass *WaitGroup |
| Closed channel from receiver | Only sender closes |
| Type switch wrong order | Concrete before interface |

## Run any file
### Run any file for Type and Variable
```bash
go run 01-types-and-variables/01_hello_world.go
```

```bash
go run 01-types-and-variables/02_variables.go
```

```bash
go run 01-types-and-variables/03_primitives.go
```

```bash
go run 01-types-and-variables/04_slices_maps.go
```

### Run any file for Control Flow

```bash
go run 02-control-flow/01_loops.go
```

```bash
go run 02-control-flow/02_functions.go
```

```bash
go run 02-control-flow/03_defer_panic.go
```