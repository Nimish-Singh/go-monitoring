FROM golang:alpine as builder

WORKDIR /build

COPY ./ .

RUN cd ./src && go build -o ./../build

FROM alpine

WORKDIR /build

COPY --from=builder /build/build/go-monitoring /build/go-monitoring

CMD ["./go-monitoring"]
