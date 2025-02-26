FROM cgr.dev/chainguard/go AS builder
COPY . /app
RUN cd /app && go build

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/todos /usr/bin/

EXPOSE 8080

ENTRYPOINT ["/usr/bin/todos"]
