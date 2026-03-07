FROM golang:1.26.1-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/app ./main.go

FROM gcr.io/distroless/static-debian12:nonroot AS runtime

WORKDIR /app

COPY --from=builder /out/app /app/app
COPY --chown=nonroot:nonroot templates /app/templates
COPY --chown=nonroot:nonroot assets /app/assets

ENV APP_PORT=8080

EXPOSE 8080

ENTRYPOINT ["/app/app"]
