# Lab 6: mockgen

Generated mock with uber/mock: `mocks/mock_cache.go`

Generate:
```bash
mockgen -source=shortener/cache.go -destination=mocks/mock_cache.go -package=mocks
```

Features:
- `EXPECT()` - set expectations
- `Return()` - specify return values
- `gomock.Any()` - match any argument
- `gomock.InOrder()` - verify call order
- `Times(n)` - verify call count

Run tests:
```bash
go test ./shortener -v -run "TestURLServiceV2_"
```

Next: `git checkout lab-7-integration`
