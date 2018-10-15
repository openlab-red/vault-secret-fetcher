FROM golang:1.11 AS builder

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/vault-agent-token-handler

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure --vendor-only

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app .

FROM registry.access.redhat.com/rhel7-atomic:latest

COPY --from=builder /app ./

ENTRYPOINT ["./app"]
