FROM golang:1.11.0-alpine3.8
WORKDIR /go/src/app
RUN apk --no-cache add git
RUN wget -qO- https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
COPY . .
RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/src/app .
ENTRYPOINT ["./app"]
