FROM golang:1.13 AS builder

# Copy the code from the host and compile it
WORKDIR $GOPATH/src/github.com/MaritimeOptima/tide-service
COPY Gopkg.toml Gopkg.lock ./
COPY Makefile ./
COPY . ./
RUN make bin
COPY ./bin/app /app

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app ./

ENTRYPOINT ["./app"]

