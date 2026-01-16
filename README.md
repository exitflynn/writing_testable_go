# Writing Testable Go

A hands-on workshop: from Go basics to professional testing patterns.

## Prerequisites

- Git
- Go 1.21+ (`go version` to check)
- VS Code or any editor

## Quick Setup

```bash
git clone <repo-url>
cd writing-testable-go
go mod download
```

## Cloud Environment (No Install Required)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/exitflynn/writing-testable-go)

## Workshop Stages

| Stage | Branch | Topic |
|-------|--------|-------|
| 1 | `stage-1-template` | Go basics: structs, functions, errors |
| 2 | `stage-2-template` | Writing tests |
| 3 | `stage-3-interfaces` | Interface introduction |
| 4 | `stage-4-http` | HTTP server with Gin |
| 5 | `stage-5-bad-code` | What not to do |
| 6 | `stage-6-mockgen` | Mocking with uber/mock |
| 7 | `stage-7-integration` | Integration tests |

**Falling behind?** Checkout the `-complete` version of any stage:
```bash
git checkout stage-1-complete
```

## Verify Setup

```bash
go version
# Should print: go version go1.21+ ...
```
