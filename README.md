# gopok-backend
A backend for a simple blogging site written in go. It uses MongoDB as its database.

Please refernece the [API documentation](https://github.com/gopok/gopok-backend/blob/master/docs/api.md) for all supported features.

# Installation :cd:

| :warning: | Warning |
| ----------| --------|
||Remember to clone this repo in `$GOROOT/src/github.com/gopok/gopok-backend`, so that the code is able to find its own packages.
|| `$GOROOT` is by default set to `~/go`

1. Install [dep](https://golang.github.io/dep/docs/installation.html)
2. Run `dep ensure` (to download all libraries)
3. Run `go run cmd/run.go`
4. Done! :)

