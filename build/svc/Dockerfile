FROM golang:1.18 as builder

# Copy the rest of the source code
WORKDIR /usr/src/app
COPY Makefile go.mod go.sum ./
COPY cmd ./cmd
COPY src ./src
COPY docs ./docs

# Compiling...
RUN make compile

# Runtime
FROM alpine:3.15

RUN apk add --no-cache tzdata

ARG ARG_PORT=8000

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/bin/compile /usr/src/app
COPY migrations /usr/src/app/migrations
COPY docs /usr/src/app/docs
COPY configs /usr/src/app/configs

EXPOSE ${ARG_PORT}

ENTRYPOINT ["./svc"]
