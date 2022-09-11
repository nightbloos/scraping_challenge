FROM golang:1.19-alpine as build

WORKDIR /src/services/cometco-scraper
COPY . /src
RUN go mod download
RUN	CGO_ENABLED=0 go build -o /bin/srv ./cmd/

FROM alpine:3 as production

RUN apk --no-cache upgrade && apk add --no-cache chromium

COPY --from=build /bin/srv /app/srv
ENTRYPOINT /app/srv
