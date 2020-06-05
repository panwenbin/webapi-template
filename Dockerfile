FROM golang:latest as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build main.go

FROM panwenbin/alpinetz:latest

COPY --from=builder /builder/main /app/api

WORKDIR /app

CMD ["./api"]
