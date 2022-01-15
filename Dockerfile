FROM golang:1.17-alpine3.14 as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser && mkdir /go-representer

# source code
WORKDIR /go-representer
COPY ./go.mod /go-representer/go.mod

# download dependencies
RUN go mod download

# Create run.sh
RUN printf '%s\n' '#!/bin/sh' '/opt/representer/bin/representer "$@"' > /go/bin/run.sh
RUN chmod +x /go/bin/run.sh

# get the rest of the source code
COPY . /go-representer

# build
RUN go generate .
RUN GOOS=linux GOARCH=amd64 go build --tags=build -o /go/bin/representer ./cmd/representer

# Build a minimal and secured container
# The ast parser needs Go installed for import statements.
# Therefore, unfortunately we cannot build from scratch as we would normally do with Go.
FROM golang:1.17-alpine3.14
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin /opt/representer/bin
USER appuser
WORKDIR /opt/representer
ENTRYPOINT ["/opt/representer/bin/run.sh"]
