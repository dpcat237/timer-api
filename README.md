# Timer API

`timer-api` API provides endpoints for timer web. It's a simple example of API wrote with Golang

## Required environment variables

| Variable                     | Description                                | Default    | Optional |
|------------------------------|--------------------------------------------|:----------:|:--------:|
| `TIMER_HTTP_PORT`            | TCP port on which to start HTTP server     |    8080    |   true   |
| `TIMER_MODE`                 | Environment mode                           |     dev    |   true   |

## Run linting and testing

```
.scripts/check.sh
```

## Build and run locally

`timer-api` uses [dep](https://golang.github.io/dep/) to manage dependencies.

```
dep ensure -vendor-only -v
make
# export env variables described above ...
./timer-api
```

## Required environment variables to build Docker image

| Variable                  | Description                                            |
| --------------------------|:------------------------------------------------------:|
| `TIMER_GITLAB_TOKEN`      | The token for download of repositories from gitlab.com |
