#---------------------------------------
# Build a executable binary
#---------------------------------------
FROM golang:1.21-alpine as builder

# Environment
ARG GOPROXY_DEFAULT
ENV GOPROXY=$GOPROXY_DEFAULT

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
FROM alpine:latest
# Copy go bin
COPY --from=builder /app/bin /app/bin
COPY --from=builder /app/main /app/main
COPY --from=builder /app/server/migrations /app/migrations
COPY --from=builder /app/server/email/templates /app/email/templates

RUN chmod +x /app/main
EXPOSE 8000
CMD ["/app/main", "serve"]

