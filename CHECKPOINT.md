# Lab 6: DIY Mock

Hand-written mock: `shortener/mock_cache_diy.go`

Implements `CacheClient` interface with:
- In-memory data store
- Error simulation (`SetErr`, `GetErr`)
- Call tracking (`SetCalls`, `GetCalls`)

Run tests:
```bash
go test ./shortener -v -run TestURLServiceV2
```

Next: `git checkout lab-6-mockgen`
