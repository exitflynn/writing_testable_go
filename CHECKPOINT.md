# Stage 6: Decoupled

Compare `bad_example/tightly_coupled.go` with `shortener/cache.go`.

## The fix
1. Define `CacheClient` interface
2. Accept interface in constructor
3. `RedisCache` implements the interface

## Key insight
`URLServiceV2` doesn't know or care if it's using Redis, memory, or anything else.

## Files
- `shortener/cache.go` - interface + URLServiceV2
- `shortener/redis_cache.go` - Redis implementation

## Next
```bash
git checkout stage-6-mock-diy
```
