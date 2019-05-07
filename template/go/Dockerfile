FROM golang:1.11.9-alpine3.9 as builder

ENV GOPROXY=https://proxy.golang.org
ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk --no-cache add curl \
    && echo "Pulling watchdog binary from Github." \
    && curl -sSLf https://github.com/openfaas-incubator/of-watchdog/releases/download/0.5.0/of-watchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && apk del curl --no-cache

RUN mkdir -p /handler
WORKDIR /handler
COPY . .

RUN mv ./function/go.mod .
RUN mv ./function/go.sum . || :
RUN mv ./function/vendor . || :
RUN go mod vendor

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*" -not -path "./function/vendor/*"))" || { echo "Run \"gofmt -s -w\" on your Golang code"; exit 1; }

RUN go build --ldflags "-s -w" -a -installsuffix cgo -o handler . && \
    go test $(go list ./... | grep -v /vendor/) -cover

FROM alpine:3.9
# Add non root user and certs
RUN apk --no-cache add ca-certificates \
    && addgroup -S app && adduser -S -g app app \
    && mkdir -p /home/app \
    && chown app /home/app

WORKDIR /home/app

COPY --from=builder /handler/handler    .
COPY --from=builder /usr/bin/fwatchdog         .
COPY --from=builder /handler/function/  .

RUN chown -R app /home/app

USER app

ENV fprocess="./handler"
ENV mode="serializing"

CMD ["./fwatchdog"]