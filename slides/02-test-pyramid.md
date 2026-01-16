# The Test Pyramid

```
        /\
       /  \      E2E (few, slow, expensive)
      /----\
     /      \    Integration (some)
    /--------\
   /          \  Unit (many, fast, cheap)
  --------------
```

## Unit Tests
- Test one function/method
- No external dependencies
- Fast: thousands per second

## Integration Tests
- Test components together
- Real database, real cache
- Slower but realistic

## E2E Tests
- Full system test
- Browser, API, everything
- Slow, flaky, but catches real bugs
