FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="gin-base"

RUN mkdir -p /app
WORKDIR /app

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy data vào WORKDIR
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o ./cmd/myapp/main ./cmd/myapp/main.go

########  ----------------------------------------------------------------
FROM alpine:latest

RUN apk --no-cache add tzdata zip ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder  /app .


CMD ["./cmd/myapp/main"]