# Cross compiled on the build platform, so a multi arch build needs no
# emulation.
FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.27.1-alpine3.24@sha256:8a5910f31396cd4d89662f56c68b3ae31d374308270a1c3bd96672ee5ed43414 AS build

ARG TARGETOS
ARG TARGETARCH

WORKDIR /src

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

# vet and test run natively for the build platform, the binary is built for
# the target.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go vet ./... && \
    CGO_ENABLED=0 go test ./... && \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /out/qr-code-generator .

# No CA bundle, no files: the binary makes no outbound calls and reads
# nothing from disk in server mode.
FROM scratch

LABEL org.opencontainers.image.source="https://github.com/simonostendorf/qr-code-generator" \
      org.opencontainers.image.title="qr-code-generator" \
      org.opencontainers.image.description="HTTP API that renders QR codes as PNG, optionally with a logo in the middle"

COPY --from=build /out/qr-code-generator /qr-code-generator

USER 65532:65532
EXPOSE 8000

ENTRYPOINT ["/qr-code-generator"]
CMD ["server"]
