# Stage 2: Writing Tests

## Goal
Write tests for the shortener package.

## Your Task
Edit `shortener/shortener_test.go` and add test cases for:
- `TestValidateURL` - valid URLs, invalid schemes, missing host
- `TestGenerateShortCode` - deterministic output, different inputs  
- `TestCreateMapping` - valid/invalid URL handling

## Verify
```bash
go test ./shortener -v
```

## Hints
- Use table-driven tests with `t.Run`
- Think about edge cases: empty strings, weird protocols

## Stuck?
```bash
git checkout stage-2-complete
```
