FROM alpine

WORKDIR /build

COPY ./build/go-monitoring .

CMD ["./go-monitoring"]
