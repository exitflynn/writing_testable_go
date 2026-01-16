# Dependency Injection

## Bad: Create dependencies inside

```go
func HandleRequest() {
    db := sql.Open("postgres", "...")  // Hard to test!
    cache := redis.NewClient(...)      // Needs real Redis!
    // ...
}
```

## Good: Accept dependencies

```go
func HandleRequest(db Database, cache Cache) {
    // Use whatever is passed in
}
```

In tests: pass fakes/mocks
In production: pass real clients

---

## The Pattern

1. Define interface for dependency
2. Accept interface in constructor/function
3. In tests: pass mock
4. In prod: pass real implementation
