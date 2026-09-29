# QR Code Generator

Creates QR codes for a URL as PNG, in any color, with a transparent background
if wanted and with an optional logo in the middle.

It has two parts:

- **Backend** (Go, repository root): an HTTP API on port 8000 with
  `GET /api/health` and `POST /api/generate`, which takes a JSON body and
  answers with the PNG. The same binary also has a `generate` subcommand that
  writes a QR code to a file.
- **Frontend** (Vue and Vite, `frontend/`): the web interface. It calls `/api`
  on its own origin; in the image nginx proxies `/api/` to the backend.

## Local development

Backend:

```sh
go run . server            # listens on :8000, --port changes it
go test ./...
```

Frontend, in a second terminal:

```sh
cd frontend
npm ci
npm run dev
```

The vite dev server proxies `/api` to `http://localhost:8000`; set
`BACKEND_URL` to use another backend.

Both images together:

```sh
docker compose up --build  # http://localhost:8080
```

## Images

| Image | Dockerfile | Port | Runs as |
|---|---|---|---|
| `qr-code-generator` | `Dockerfile` | 8000 | uid 65532 on `scratch`, read only root filesystem |
| `qr-code-generator-frontend` | `Dockerfile.frontend` | 8080 | uid 101 on `nginx-unprivileged` |

The frontend image serves the built app and proxies `/api/` to `BACKEND_URL`
(default `http://backend:8000`, without a trailing slash). Its health check is
`GET /healthz`.

Both images are built and deployed from
[github.com/simonostendorf/infra](https://github.com/simonostendorf/infra),
where this repository is the submodule `docker-images/qr-code-generator`.
`make docker-images IMAGE=qr-code-generator` there builds both images for
linux/amd64 and linux/arm64. `VERSION` holds the release version and is the
tag of both images; raise it with every change, since an image whose version
already exists in the registry is not built again.
