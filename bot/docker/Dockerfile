# ---------------
# -- 1st Stage

FROM golang:1-alpine AS builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .

# RUN go mod tidy

COPY . .
COPY .env_isam .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
  -ldflags '-s -w -extldflags "-static"' -trimpath \
  -o ./bin/isambot ./cmd/isam/main.go

# ---------------
# -- 2nd Stage

FROM alpine:3.13

WORKDIR /app
ENV PATH=/app/:$PATH

LABEL maintainer="Waldir Borba Junior <wborbajr@gmail.com>" \
  version="v0.0.1-2022" \
  description="iSamBot | waldirborbajr/isambot:latest"

ENV LANG=en_US.UTF-8 \
  LANGUAGE=en_US.UTF-8

RUN apk add --update --no-cache \
  tzdata \
  htop \
  apk-cron \
  && cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime \
  && echo "America/Sao_Paulo" > /etc/timezone

RUN adduser -S -D -H -h /app isambot

USER isambot

WORKDIR /app

# COPY --from=build /go/src/app/run.sh run.sh
COPY --from=builder /go/src/app/bin/isambot .
COPY --from=builder /go/src/app/.env_isam .

EXPOSE 9090

ENTRYPOINT ["/app/isambot"]

CMD ["/bin/sh"]


