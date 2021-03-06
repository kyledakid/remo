# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -o kyle .


# generate clean, final image for end users
FROM alpine:3.11.3
COPY --from=builder /build/kyle .

# executable
ENTRYPOINT [ "./kyle" ]
