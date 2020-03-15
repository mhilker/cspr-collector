# CSPR Collector

Content Security Policy Report Collector.

## Synopsis

```bash
$ ./cspr-collector --help
Usage of ./build/cspr-collector:
  -host string
        address to listen for http requests on (default "127.0.0.1:8080")
  -n int
        the number of workers to start (default 4)
  -output-es
        enable elasticsearch output
  -output-es-ca-file string
        ca file for elasticsearch
  -output-es-cert-file string
        cert file for elasticsearch
  -output-es-host string
        elasticsearch host to send the csp violations to (default "http://localhost:9200/")
  -output-es-index string
        elasticsearch index to save the csp violations in (default "cspr-violations")
  -output-es-key-file string
        key file for elasticsearch
  -output-http
        enable http output
  -output-http-host string
        http host to send the csp violations to (default "http://localhost:80/")
  -output-stdout
        enable stdout output
```

## Build and run

### On your machine

```bash
go build -o build/cspr-collector ./cmd/cspr-collector/main.go
./build/cspr-collector -host 0.0.0.0:8080 -output-stdout
```

### Via docker

```bash
docker build -t mhilker/cspr-collector:latest -f cmd/cspr-collector/Dockerfile .
docker run -p 8080:8080 mhilker/cspr-collector:latest -host 0.0.0.0:8080 -output-stdout
```

### Via docker-compose

```bash
docker-compose -f cmd/cspr-collector/docker-compose.yml build
docker-compose -f cmd/cspr-collector/docker-compose.yml up
```

## Example request

```bash
curl -X POST \
  http://localhost:8080 \
  -H 'Content-Type: application/csp-report' \
  -d '{
    "csp-report": {
        "document-uri": "https://example.com/path/to/file",
        "referrer": "",
        "violated-directive": "script-src-elem",
        "effective-directive": "script-src-elem",
        "original-policy": "default-src '\''self'\''; img-src '\''self'\'' https://*.ytimg.com; script-src-elem '\''self'\'' https://storage.googleapis.com https://www.youtube.com; connect-src '\''self'\'' https://www.googleapis.com; frame-src '\''self'\'' https://www.youtube.com; base-uri '\''self'\''; frame-ancestors '\''none'\''; form-action '\''self'\''; block-all-mixed-content; report-uri https://reporting.example.com/;",
        "disposition": "report",
        "blocked-uri": "https://www.youtube.com/iframe_api",
        "line-number": 1,
        "column-number": 7982,
        "source-file": "://example.com/static/js/7.74a7cce6.chunk.js",
        "status-code": 0,
        "script-sample": ""
    }
}'
```

## Requirements

### Elasticsearch Output

The elasticsearch output requires an elasticsearch index called `csp-violations` with a doc-type `_doc`.
A mapping template is included in the `template.json` file.

```bash
curl -X POST \
    --header "Content-Type: application/json" \
    --data-binary @template.json \
    http://localhost:9200/_template/cspr-violations
```

## Code Style

```bash
go fmt ./...
```

## License

The MIT License (MIT). Please see the [license file](LICENSE.md) for more information.
