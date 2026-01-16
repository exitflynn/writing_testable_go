# Stage 7: Integration Tests

Test with real (or near-real) Redis.

## testcontainers (requires Docker)
Spins up real Redis in a container:
```bash
go test ./integration -v -tags=integration -run TestWithRealRedis
```

## miniredis (no Docker)
In-memory Redis implementation:
```bash
go test ./integration -v -run TestWithMiniredis
```

## When to use what
- **Unit tests**: mocks (fast, isolated)
- **Integration tests**: miniredis or testcontainers
- **E2E/staging**: real infrastructure

## Record-Replay (brief mention)
For HTTP APIs, check out `go-vcr` or `httpmock` to record real responses and replay in tests.

## Workshop Complete!
You've learned:
1. Go basics with real project
2. Table-driven tests
3. Finding bugs through testing
4. Interfaces for testability
5. Dependency injection
6. DIY mocks and mockgen
7. Integration testing
