# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build Commands
- Build: `go build`
- Run: `go run main.go`
- Format code: `go fmt ./...`
- Verify: `go vet ./...`
- Test: `go test ./...`
- Test single package: `go test ./[package-path]`

## Code Style Guidelines
- Follow standard Go formatting conventions
- Constants: UPPERCASE_SNAKE_CASE
- Functions/variables: camelCase
- Error handling: Check errors explicitly and return early
- Import ordering: standard library first, then third-party, then local packages
- Package organization: Group by functionality (user, order, config, etc.)
- Use Go pointers appropriately for return values
- JSON field handling: Use struct tags for serialization/deserialization
- Error responses should follow the jsonrpc error pattern
- Maintain consistent response structures for all tools