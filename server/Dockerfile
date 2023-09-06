#---------------------------------------
# Build a executable binary
#---------------------------------------
FROM golang:1.21-alpine as builder

# Set working directory
WORKDIR /app/server

# Install app dependencies
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

#Build tools
RUN mkdir -p /app/bin
COPY ./tools ./tools
RUN go generate -tags devtools ./tools

# Add source code
COPY ./ ./

# Build API Server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /app/main ./cmd/platform

#---------------------------------------
# Build a small image
#---------------------------------------
FROM scratch
# Copy go bin
WORKDIR /app

COPY --from=builder /app/bin ./bin
COPY --from=builder /app/main ./main
COPY --from=builder /app/server/migrations ./migrations
COPY --from=builder /app/server/email/templates ./email/templates

EXPOSE 80
CMD ["./main", "serve"]

