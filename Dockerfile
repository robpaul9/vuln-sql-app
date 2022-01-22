FROM golang:1.17-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
COPY static ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o vulnsqlapp cmd/vulnsqlapp/main.go


FROM alpine:3.12.8
RUN apk add --no-cache ca-certificates
COPY --from=builder /build/vulnsqlapp /vulnsqlapp

CMD ["sh", "-c", "/vulnsqlapp"]