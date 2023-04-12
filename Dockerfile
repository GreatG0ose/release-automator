# Build release-automator
FROM golang:alpine as build

WORKDIR release-automator

# Copies your code file from your action repository to the filesystem path `/` of the container
COPY . .

RUN go install ./cmd/release-automator/...


# Add bin to run release-automator
FROM alpine:latest

COPY --from=build /go/bin/release-automator /bin/release-automator
COPY entrypoint.sh /bin/entrypoint.sh

WORKDIR app

ENTRYPOINT ["/bin/entrypoint.sh"]
