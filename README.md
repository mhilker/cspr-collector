# CSP Reporter

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
