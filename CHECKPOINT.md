# Stage 6: DIY Mock

We wrote our own mock: `shortener/mock_cache_diy.go`

## How it works
- Implements `CacheClient` interface
- Stores data in a map
- Has `SetErr`/`GetErr` fields to simulate errors
- Tracks call counts

## Run tests
```bash
go test ./shortener -v -run TestURLServiceV2
```

## Downsides of DIY mocks
- Tedious to write
- Need to update when interface changes
- No built-in expectation checking

## Next
```bash
git checkout stage-6-mockgen
```
