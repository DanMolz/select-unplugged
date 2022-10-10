# Run tests

```
make test
```

# Mocks

Using https://github.com/vektra/mockery .

Generate a mock with `( cd sp && mockery --name=$interface )`

# Compiling for a pi

```
env GOOS=linux GOARCH=arm GOARM=5 go build
```

# CLI code management

See https://github.com/spf13/cobra-cli/blob/main/README.md