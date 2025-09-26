# -------- STAGE 1: Build --------
FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /tutorial-backend main.go

# -------- STAGE 2: Test --------
FROM builder AS tester
RUN go test -v ./...

# -------- STAGE 3: Release --------
FROM gcr.io/distroless/base-debian11 AS release

WORKDIR /

# Copy binary
COPY --from=builder /tutorial-backend /tutorial-backend

# Non-root user
USER nonroot:nonroot

# Expose port
EXPOSE 8080

ENTRYPOINT ["/tutorial-backend"]

#run: docker build -t tutorial-backend:latest .
