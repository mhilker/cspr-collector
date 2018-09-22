# CSP Reporter

## Synopsis

```
$ ./bin/cspreporter --help
Usage of ./bin/cspreporter:
  -host string
        address to listen for http requests on (default "127.0.0.1:8080")
  -n int
        the number of workers to start (default 4)
  -output-es
        enable elasticsearch output
  -output-es-host string
        elasticsearch host to send the csp violations to (default "http://localhost:9200/")
  -output-http
        enable http output
  -output-http-host string
        http host to send the csp violations to (default "http://localhost:80/")
  -output-stdout
        enable stdout output (default true)
```

## Installation

```bash
$ go install github.com/mhilker/cspreporter && ./bin/cspreporter
```

## Dependencies

```bash
$ dep ensure
```
## Requirements

### Elasticsearch Output

The elasticsearch output requires an elasticsearch index called `csp-violations` with a doc-type `_doc`.
A mapping template is included in the `template.json` file.

## Code Style

```bash
$ go fmt github.com/mhilker/cspreporter
```

## Build via docker

### Build

```bash
$ docker build . -t mhilker/cspreporter:latest
```

### Push

```bash
$ docker push mhilker/cspreporter:latest
```
