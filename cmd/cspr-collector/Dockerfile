FROM golang:1.13.1-alpine3.10 as builder
WORKDIR /go/src
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -o cspr-collector cmd/cspr-collector/main.go

FROM alpine:3.10
RUN apk add --no-cache ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/cspr-collector .
EXPOSE 80
ENTRYPOINT ["./cspr-collector"]
