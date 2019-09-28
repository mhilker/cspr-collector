FROM golang:1.13.1-alpine3.10 as builder
WORKDIR /go/src/app
RUN apk --no-cache add git
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go get -d -v ./...
RUN go install -v ./...
#RUN go build -a -installsuffix cgo -o app .

FROM alpine:3.10
RUN apk add --no-cache ca-certificates
WORKDIR /app/
COPY --from=builder /go/src/app .
EXPOSE 80
ENTRYPOINT ["./app"]
