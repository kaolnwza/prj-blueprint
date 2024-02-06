FROM golang:1.21.6 AS builder


WORKDIR /go/app

COPY . .

RUN go mod tidy

RUN cd cmd && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build main.go

FROM alpine:3.13
WORKDIR /app

COPY --from=builder /go/app/cmd/main .

ENV TZ=Asia/Bangkok

CMD ["./main"]