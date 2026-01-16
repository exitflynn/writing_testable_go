# Stage 4: HTTP Server

We now have a real web service.

## Run it
```bash
go run .
```

Open http://localhost:8080 in your browser.

## Try it
1. Enter `https://google.com` - works!
2. Enter `https://` (no host) - what happens?

## Key files
- `server/server.go` - Gin routes and handlers
- `server/static/index.html` - simple frontend
- `main.go` - ties everything together

## Next
```bash
git checkout stage-4-fixed
```
