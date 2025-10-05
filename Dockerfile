FROM golang:1.25.1-alpine3.22 as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Add a non-root user to run our code as
RUN adduser --disabled-password appuser

# Copy the source code into the container
# and make sure appuser owns all of it
COPY --chown=appuser:appuser . /opt/representer

# Build and run the representer with appuser
USER appuser

# Default is 'go telemetry local' which saves telemetry locally.
# Since data will never be uploaded, turn it off to avoid unnecessary file writes.
RUN go telemetry off

# This populates the build cache with the standard library
# so future compilations are faster
RUN go build std

WORKDIR /opt/representer

# Download dependencies
RUN go mod download

# build
RUN go build --tags=build -o /opt/representer/bin/representer ./cmd/representer

ENTRYPOINT ["sh", "/opt/representer/bin/run.sh"]
