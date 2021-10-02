###############
# Build Image
###############
FROM golang:1.17 as builder

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

WORKDIR /src/secret-share

# Copy the Go Modules manifests
COPY go.* ./
RUN go mod download

COPY main.go .
COPY app ./app
COPY cache ./cache
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /bin/app
RUN openssl req -x509 -newkey rsa:2048 -keyout /tmp/key.pem -out /tmp/cert.pem -days 365 -subj "/CN=secret-share" -nodes

###############
# Final Image
###############
FROM scratch

WORKDIR /secret-share/

COPY --from=builder /bin/app .
COPY --from=builder /tmp/*.pem ./ssl/
COPY ./minified/assets ./assets
COPY ./minified/templates ./templates

ENV GIN_MODE=release
ENV ADDR=0.0.0.0
ENTRYPOINT ["/secret-share/app"]
