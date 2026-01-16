# Stage 5: Bad Code Example

Look at `bad_example/tightly_coupled.go`.

## What's wrong?
1. Redis client created inside methods
2. Hardcoded configuration (localhost:6379)
3. No way to swap implementations
4. Every call creates a new connection

## Why is this bad for testing?
- Need a real Redis running
- Can't mock the redis client
- Can't control return values
- Tests are slow and flaky

## Next
```bash
git checkout stage-6-decoupled
```
