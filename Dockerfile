FROM golang:1.11-alpine3.8 as builder
WORKDIR /go/src/app
RUN apk --no-cache add git
RUN wget -qO- https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:3.8
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .
EXPOSE 80
ENTRYPOINT ["./app"]
