FROM golang:1.12 as builder

ARG GOPROXY
ENV GORPOXY ${GOPROXY}

ADD . /builder

WORKDIR /builder

RUN go build main.go

FROM golang:1.12

COPY --from=builder /builder/main /app/api

WORKDIR /app

CMD ["./api"]

EXPOSE 8080