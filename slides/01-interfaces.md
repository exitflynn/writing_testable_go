# Interfaces in Go

An interface defines behavior, not data.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

Any type with a `Write` method satisfies `Writer`.

No `implements` keyword - it's implicit.

---

## Why Interfaces Matter for Testing

```go
// Production: real database
var store URLStore = &PostgresStore{}

// Testing: fake in-memory
var store URLStore = &MemoryStore{}
```

Same code, different implementations.
