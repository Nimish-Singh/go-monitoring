FROM golang:alpine

WORKDIR /build

COPY ./ .

RUN go build -o ./build

CMD ["./build/go-monitoring"]
