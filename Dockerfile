ARG GO_VERSION=1.14.4

FROM golang:1.19.5-alpine AS builder

LABEL maintainer="Samira Afrin Alam <samiraafrin.alam@gmail.com>"

RUN apk update && apk upgrade  

WORKDIR /src
COPY ./ ./

RUN CGO_ENABLED=0 go build -mod=vendor -o /app .

FROM scratch AS final

COPY --from=builder /app /app


EXPOSE 8000
ENTRYPOINT ["/app"]

