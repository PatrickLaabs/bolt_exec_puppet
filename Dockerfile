# Dockerfile
FROM alpine
COPY puppet_bolt_exec /usr/bin/puppet_bolt_exec
ENTRYPOINT ["/usr/bin/puppet_bolt_exec"]