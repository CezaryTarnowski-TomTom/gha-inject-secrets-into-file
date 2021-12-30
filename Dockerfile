FROM golang:1.17 as builder

WORKDIR /gha-inject-secrets-into-file
COPY . /gha-inject-secrets-into-file

RUN go get -d -v

# Statically compile our gha-inject-secrets-into-file for use in a distroless container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o gha-inject-secrets-into-file .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static

COPY --from=builder /gha-inject-secrets-into-file/gha-inject-secrets-into-file /gha-inject-secrets-into-file

ENTRYPOINT ["/gha-inject-secrets-into-file"]