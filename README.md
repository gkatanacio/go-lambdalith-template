# go-lambdalith-template

Template for a Golang Lambdalith API. Deploys an [AWS Lambda application](https://docs.aws.amazon.com/lambda/latest/dg/deploying-lambda-apps.html) by utilizing [Serverless Framework](https://www.serverless.com) (v3). Env file (i.e., `.env`) is used for configuration.

Since this tries to follow the [3 Musketeers](https://3musketeers.pages.dev/) pattern, the following are heavily used:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

## Usage

#### configure

```bash
$ make .env
```

- see generated `.env` file for configuration

**Note:** For deployment via CI/CD, `CICD_MODE` environment variable should be set to `true` in the build server. All env configuration should then be set in the build server's environment variables. They will automatically be used if `.env` is generated from `.env.cicd`. See [Docker's handling of env variables from host environment](https://docs.docker.com/engine/reference/commandline/run/#:~:text=You%20can%20also%20use%20variables%20exported%20to%20your%20local%20environment) for more details.

#### tidy dependencies

```bash
$ make deps
```

#### run tests

```bash
$ make test
```

#### build serverless functions

```bash
$ make build
```

- this generates `bin` directory to be used in deployment

#### deploy serverless application

```bash
$ make deploy
```

### Helpers during development:

#### format all .go files in project (using go fmt)

```bash
$ make fmt
```

#### generate test mocks (to be used with [stretchr/testify](https://github.com/stretchr/testify)) for all interfaces in project

```bash
$ make mocks
```

- can be configured in `.mockery.yml`
