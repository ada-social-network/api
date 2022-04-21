# Ada Social Network API

> Manage resources for the Ada Social Network.

## Usage

### Using docker-compose

The application can be run using docker-compose:

```shell
docker-compose up
```

More details in [the Docker Compose section](#docker-compose).

### Using the ada-api binary

The application can be used with the following command:

```
./ada-api
```

if you need help use the following command:

```
./ada-api --help

Usage of ada-api:

  -auth
        Use api authentication (default true)
  -graceful-timeout duration
        the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m (default 15s)
  -http-host string
        Default interface (default "0.0.0.0")
  -http-port int
        Default port (default 8080)
  -mode string
        Running mode, can be 'debug', 'release' or 'test' (default "release")
  -sqlite-dsn string
        sqlite database file (dsn) that will store data (default "gorm.db")
  -version
        Show application current version
```

## CORS

CORS is disabled by default. it means that all request should have the same domain
as the API.

In order to use another domain, it is possible to activate CORS by passing to command arguments
an allowed domain.

For example:

```shell
# Allow request from adahub.com (for example)
./ada-api --allowed-domain="https://adahub.com"
```

For development, you can also allow all domain (useful for accepting any request from any domain):

```shell
# Allow all
./ada-api --allowed-domain="*"

# Or specific front-end
./ada-api --allowed-domain="http://localhost:3000"
```

You can also use a CORS plugin (see Chrome Store for a CORS Plugin) in order to allow Cross domain request
during development. For example : https://chrome.google.com/webstore/detail/allow-cors-access-control/lhobafahddgcelffkeicbaginigeejlf

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
- For releasing the application:
  - version: `./scripts/release.sh v1.0.0`
  More info about version here: [semantic version](https://www.jvandemo.com/a-simple-guide-to-semantic-versioning/)

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

### Docker

- Docker image: `ghcr.io/ada-social-network/api`
- Docker tags: 
  - `latest`: latest version of the image
  
#### Get locally `latest` docker image

You can get the last image of the api using `docker pull`.

For example:

```shell
docker pull ghcr.io/ada-social-network/api:latest
```

if you don't have the image locally or if the local image is outdated you will
have a stdout like that:

```text
latest: Pulling from ada-social-network/api
Digest: sha256:3eb90ac5515e0acf8d3829ec1c860c5c231e204fc3d5747d085800c5cfd6e1c3
Status: Downloaded newer image for ghcr.io/ada-social-network/api:latest
ghcr.io/ada-social-network/api:latest
```

if you already have the image locally you should have a stdout like that:

```text
latest: Pulling from ada-social-network/api
Digest: sha256:3eb90ac5515e0acf8d3829ec1c860c5c231e204fc3d5747d085800c5cfd6e1c3
Status: Image is up to date for ghcr.io/ada-social-network/api:latest
ghcr.io/ada-social-network/api:latest
```

#### Run the image locally

You can run the image and bind:
 
- your local port (e.g. `8080`) to the container port (e.g. `8080`) for accessing the container through your machine

```shell
# run latest
docker run  --rm -p 8080:8080 ghcr.io/ada-social-network/api
```

if you want to reuse your local file gorm.db in the container you can run:
```shell
docker run -p 8080:8080 -v $PWD/gorm.db:/usr/local/ada/data/gorm.db ghcr.io/ada-social-network/api
```

you can add extra flags like that:

```shell
# add flag for getting help
docker run ghcr.io/ada-social-network/api --help

# add flag for debug mode
docker run ghcr.io/ada-social-network/api --mode=debug
```
### Docker compose

In order to simplify the use of the api you can use docker-compose.

For example: 

```shell
# start the application in docker-compose attached mode, you can interrupt the process with CTRL+C
docker-compose up

# start the application in docker-compose detached mode(run in the background)
docker-compose up -d

# stop the application
docker-compose stop

# remove stopped application
docker-compose rm
```

## References

- https://golang.org/pkg/testing/