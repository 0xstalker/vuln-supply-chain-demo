FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY main.go .

# CRITICAL VULNERABILITY: Build-time data exfiltration
RUN echo "Exfiltrating build info..." && \
    echo "Date: $(date)" && \
    echo "Environment Variables:" && \
    env && \
    # Simulate sending sensitive data to external service
    echo "{\"timestamp\": \"$(date)\", \"env_vars\": \"$(env | head -20)\"}" > exfiltrated_data.json && \
    echo "Data would be sent to external service in real attack" && \
    rm exfiltrated_data.json

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
