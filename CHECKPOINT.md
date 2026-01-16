# Stage 4: Fixed

The bug is fixed!

## What we did
1. Added host check in `ValidateURL`
2. Added test cases for `empty host` and `empty string`

## Verify
```bash
go test ./shortener -v
```

Now `https://` correctly returns an error.

## MVP Complete!
You've built a URL shortener with tests that catch edge cases.

## Next
```bash
git checkout stage-5-bad-code
```
