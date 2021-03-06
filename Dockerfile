## Builder
FROM golang:1.13.4-alpine3.10 AS build

RUN apk --no-cache add make git curl

ENV GOPATH /.go/
ENV GOPROXY https://proxy.golang.org/,direct

# Go modules
WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY Makefile /app
RUN make modules
RUN make prereq

# Build app
COPY . /app
RUN make build

## Destination image
FROM alpine:3.10

RUN apk --no-cache add ca-certificates supervisor tzdata jq python2 py2-pip
RUN pip install yq

COPY --from=build /app/go-rest-api /opt/go-rest-api/go-rest-api
COPY --from=build /.go/bin/migrate /opt/go-rest-api/migrate

COPY .docker/run.sh /opt/go-rest-api/run.sh
COPY .docker/supervisord.conf /etc/supervisord.conf

VOLUME ["/var/log/supervisor", "/opt/go-rest-api/configs", "/opt/go-rest-api/migrations"]
WORKDIR /opt/go-rest-api/
EXPOSE 8080

CMD ["/opt/go-rest-api/run.sh"]
