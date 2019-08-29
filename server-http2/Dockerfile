# golang alpine 1.12.7
FROM golang@sha256:87e527712342efdb8ec5ddf2d57e87de7bd4d2fedf9f6f3547ee5768bb3c43ff as builder
# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/mmorejon/server-http2/
COPY . .

# Fetch dependencies.
RUN go get -d -v

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/server-http2 .

ENTRYPOINT ["/go/bin/server-http2"]