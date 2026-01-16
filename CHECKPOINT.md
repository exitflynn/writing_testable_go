# Stage 3: Interfaces

This stage introduces the concept of interfaces.

Look at `shortener/store.go`:
- `URLStore` is an interface
- `MemoryStore` implements it
- Later we could add `RedisStore`, `PostgresStore`, etc.

The key insight: code that uses `URLStore` doesn't care *how* data is stored.

## Next
```bash
git checkout stage-4-http
```
