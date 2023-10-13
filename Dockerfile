FROM golang:1.20.1-alpine3.17 as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser && mkdir /go-representer

# source code
WORKDIR /go-representer
COPY ./go.mod /go-representer/go.mod

# download dependencies
RUN go mod download

# get the rest of the source code
COPY . /go-representer

# build
RUN go generate .
RUN GOOS=linux GOARCH=amd64 go build --tags=build -o /go/bin/representer ./cmd/representer

# Build a minimal and secured container
# The ast parser needs Go installed for import statements.
# Therefore, unfortunately we cannot build from scratch as we would normally do with Go.
FROM golang:1.20.1-alpine3.17
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin /opt/representer/bin
COPY bin/run.sh /opt/representer/bin/

USER appuser
WORKDIR /opt/representer
ENTRYPOINT ["/opt/representer/bin/run.sh"]
