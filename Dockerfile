# Build tool
FROM golang:alpine as build
WORKDIR app
COPY . .
RUN go build -ldflags="-w -s" -o /go/bin/release ./cmd/release/...

# Use minimalistic image to run code
FROM scratch
COPY --from=build /go/bin/release /bin/release
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR app
ENTRYPOINT ["/bin/release"]
