# Stage 1: Go Basics

## Goal
Implement the URL shortener core logic.

## Your Task
Edit `shortener/shortener.go` and implement the three functions:
- `ValidateURL` - check if URL is valid
- `GenerateShortCode` - create a deterministic short code  
- `CreateMapping` - combine both and return a mapping

## Verify
```bash
go run .
# Should print: Shortened: <code> -> https://example.com/very/long/path
```

## Hints
- Use `net/url` package's `Parse` function
- Use `hash/fnv` for hashing
- Use `strconv.FormatUint` with base 36 for short codes

## Stuck?
```bash
git checkout stage-1-complete
```
