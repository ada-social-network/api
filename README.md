# Ada Social Network API

> Manage resources for the Ada Social Network.

## Usage

The application can be used with the following command:

```
./ada-api
```

if you need help use the following command:

```
./ada-api --help
```

## Development

- lint code: `golangci-lint run`
- format whole repository: `go fmt ./...`
- build the api: `go build -o ada-api .`
- run the api: `go run .`
- run in debug mode: `go run . --mode=debug`
- run test: `go test ./...`

### Workflow

- Before commit, ensure the following command are ok:
  - lint: `golangci-lint run`
  - test: `go test ./...`
  - build: `go build -o ada-api .`

### Hooks

In order to ensure that your commits are valid, you can install 
our hooks using the following command: 

```shell
ln -s ../../.githook/pre-commit .git/hooks/pre-commit
```

### Lint

Lints are performed by Golang CI Lint with the configuration in `.golangci.yml`

> golangci-lint is a fast Go linters runner. It runs linters in parallel,
> uses caching, supports yaml config, has integrations with all major IDE
> and has dozens of linters included.

More info, here: [https://golangci-lint.run](https://golangci-lint.run/)

## References

- https://golang.org/pkg/testing/