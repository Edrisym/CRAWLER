# Stage 1: Build the Go application
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy Go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/CrawlerBot .

# Stage 2: Create a minimal image with the built binary
FROM scratch

FROM alpine:latest
COPY --from=build /app/CrawlerBot .
EXPOSE 8080

CMD ["/CrawlerBot"]
