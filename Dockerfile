# Dockerfile
FROM alpine
COPY bolt_exec /usr/bin/bolt_exec
ENTRYPOINT ["/usr/bin/bolt_exec"]