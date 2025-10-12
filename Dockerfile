FROM golang:1.24.3-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o qr-code-generator

FROM alpine:latest AS runtime

WORKDIR /app

COPY --from=build /app/qr-code-generator .

EXPOSE 8000

ENTRYPOINT ["./qr-code-generator"]
CMD ["server"]
