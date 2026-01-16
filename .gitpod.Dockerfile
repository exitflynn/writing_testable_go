FROM gitpod/workspace-go:latest

RUN go install go.uber.org/mock/mockgen@latest
