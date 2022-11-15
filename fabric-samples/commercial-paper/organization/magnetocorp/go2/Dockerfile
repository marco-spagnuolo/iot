FROM golang:1.14-alpine as compiler
RUN apk --no-cache add git ca-certificates && \
  go get -u github.com/jokeyrhyme/go-coverage-threshold/cmd/go-coverage-threshold && \
  wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.30.0
COPY build.sh /usr/bin/build

ENV buildpath=go2
COPY . /go/src/${buildpath}
WORKDIR /go/src/${buildpath}

RUN build

FROM scratch
COPY --from=compiler /tmp/app /app/
COPY --from=compiler /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080

WORKDIR /app
ENTRYPOINT ["/app/app"]
CMD [ "/main.go" ]



