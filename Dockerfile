# Stage 1. Build the binary
FROM golang:1.11

# add a non-privileged user
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/mlukanova/go-workshop
ADD . /go/src/github.com/mlukanova/go-workshop
WORKDIR /go/src/github.com/mlukanova/go-workshop

# build the binary with go build
RUN go get ./... && \
	CGO_ENABLED=0 go build -o bin/go-workshop github.com/mlukanova/go-workshop/cmd/go-workshop

# Stage 2. Run the binary
FROM scratch

ENV GO_PORT 8585
ENV GO_DIAGNOSTICS_PORT 8080

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /etc/passwd /etc/passwd
USER myapp

COPY --from=0 /go/src/github.com/mlukanova/go-workshop/bin/go-workshop /go-workshop
EXPOSE $GO_PORT
EXPOSE $GO_DIAGNOSTICS_PORT

CMD ["/go-workshop"]