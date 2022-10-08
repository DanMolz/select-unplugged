# Run tests

```
go test ./...
```

# Compiling for a pi

```
env GOOS=linux GOARCH=arm GOARM=5 go build
```

# CLI code management

See https://github.com/spf13/cobra-cli/blob/main/README.md